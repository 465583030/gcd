/*
The MIT License (MIT)

Copyright (c) 2015 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package gcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// When we get an error reading the body from the debugger api endpoint
type GcdBodyReadErr struct {
	Message string
}

func (g *GcdBodyReadErr) Error() string {
	return "error reading response body: " + g.Message
}

// Failure to unmarshal the JSON response from debugger API
type GcdDecodingErr struct {
	Message string
}

func (g *GcdDecodingErr) Error() string {
	return "error decoding inspectable page: " + g.Message
}

// The Google Chrome Debugger
type Gcd struct {
	sync.RWMutex  // for locking pages (i.e. websocket clients)
	Targets       []*ChromeTarget
	timeout       time.Duration // how much time to wait for debugger port to open up
	chromeProcess *os.Process
	chromeCmd     *exec.Cmd
	port          string
	host          string
	addr          string
	readyCh       chan struct{}
	apiEndpoint   string
	flags         []string
	env           []string
}

// Give it a friendly name.
func NewChromeDebugger() *Gcd {
	c := &Gcd{}
	c.timeout = 15
	c.host = "localhost"
	c.readyCh = make(chan struct{})
	c.Targets = make([]*ChromeTarget, 0)
	c.flags = make([]string, 0)
	c.env = make([]string, 0)
	return c
}

// Set the timeout for how long we should wait for debug port to become available.
func (c *Gcd) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Allows caller to add additional startup flags to the chrome process
func (c *Gcd) AddFlags(flags []string) {
	c.flags = append(c.flags, flags...)
}

// Add environment variables for the chrome process, useful for Xvfb etc.
func (c *Gcd) AddEnvironmentVars(vars []string) {
	c.env = append(c.env, vars...)
}

// Starts the process
// exePath - the path to the executable
// userDir - the user directory to start from so we get a fresh profile
// port - The port to listen on.
func (c *Gcd) StartProcess(exePath, userDir, port string) {
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)
	// profile directory
	c.flags = append(c.flags, fmt.Sprintf("--user-data-dir=%s", userDir))
	// debug port to use
	c.flags = append(c.flags, fmt.Sprintf("--remote-debugging-port=%s", port))

	c.chromeCmd = exec.Command(exePath, c.flags...)
	// add custom environment variables.
	c.chromeCmd.Env = os.Environ()
	c.chromeCmd.Env = append(c.chromeCmd.Env, c.env...)
	go func() {
		err := c.chromeCmd.Start()
		if err != nil {
			log.Fatalf("error starting chrome process: %s", err)
		}
		c.chromeProcess = c.chromeCmd.Process
		err = c.chromeCmd.Wait()
	}()
	go c.probeDebugPort()
	<-c.readyCh
}

// Kills the process
func (c *Gcd) ExitProcess() error {
	// close all websockets
	for _, target := range c.Targets {
		target.shutdown()
	}
	return c.chromeProcess.Kill()
}

// Gets the primary tabs/processes to work with. Each will have their own references
// to the underlying API components (such as Page, Debugger, DOM etc).
func (c *Gcd) GetTargets() ([]*ChromeTarget, error) {
	resp, err := http.Get(c.apiEndpoint)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, &GcdBodyReadErr{Message: errRead.Error()}
	}

	targets := make([]*TargetInfo, 0)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return nil, &GcdDecodingErr{Message: err.Error()}
	}

	for _, v := range targets {
		target := newChromeTarget(c.addr, v)
		c.Targets = append(c.Targets, target)
	}
	return c.Targets, nil
}

// Create a new empty tab, returns the chrome target.
func (c *Gcd) NewTab() (*ChromeTarget, error) {
	resp, err := http.Get(c.apiEndpoint + "/new")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, &GcdBodyReadErr{Message: errRead.Error()}
	}

	tabTarget := &TargetInfo{}
	err = json.Unmarshal(body, &tabTarget)
	if err != nil {
		fmt.Printf("body: %s\n", string(body))
		return nil, &GcdDecodingErr{Message: err.Error()}
	}
	return newChromeTarget(c.addr, tabTarget), nil
}

// Closes the tab
func (c *Gcd) CloseTab(target *ChromeTarget) error {
	target.shutdown() // close WS connection first
	resp, err := http.Get(fmt.Sprintf("%s/close/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// Activates (focus) the tab.
func (c *Gcd) ActivateTab(target *ChromeTarget) error {
	resp, err := http.Get(fmt.Sprintf("%s/activate/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// probes the debugger report and signals when it's available.
func (c *Gcd) probeDebugPort() {
	ticker := time.NewTicker(time.Millisecond * 100)
	timeoutTicker := time.NewTicker(time.Second * c.timeout)

	defer func() {
		ticker.Stop()
		timeoutTicker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			_, err := http.Get(c.apiEndpoint)
			if err == nil {
				c.readyCh <- struct{}{}
				return
			}
		case <-timeoutTicker.C:
			log.Fatalf("Unable to contact debugger at %s after %d seconds, gave up", c.apiEndpoint, c.timeout)
		}
	}
}
