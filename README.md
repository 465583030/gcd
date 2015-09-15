# Google Chrome Debugger (GCD)
This is primarly an auto-generated client library for communicating with a Google Chrome Browser over their [remote client debugger protocol](https://developer.chrome.com/devtools/docs/debugger-protocol). Note that their documentation is partially incorrect and does not contain a lot of the API calls that are actually available.

Because I'm lazy and there are hundereds of different custom types and API methods, this library has been automatically generated using their [protocol.json](https://code.google.com/p/chromium/codesearch#chromium/src/third_party/WebKit/Source/devtools/protocol.json&q=protocol.json&sq=package:chromium&type=cs). What this means is probably a lot of stuff doesn't work. 

The [gcdprotogen](https://github.com/wirepair/gcd/gcdprotogen) package was created to generate types and commands for gcd. Non-biolerplate code (and the majority of the logic) are in [gcd.go](https://github.com/wirepair/gcd/blob/master/gcd.go) and [chrome_target.go](https://github.com/wirepair/gcd/blob/master/chrome_target.go).


## The API
The API consists of three types, synchronous requests, asynchronous requests and events. Synchronous requests are handled by using non-buffered channels and methods can be called and will return once the value is available. Asynchronous requests have not been heavily tested. Events are handled by subscribing the response method type and calling the API's "Enable()" such as:
```Go
    target := debugger.NewTab()
	console := target.Console()
	target.Subscribe("Console.messageAdded", func(target *ChromeTarget, v []byte) {
		type EventData struct {
			Method string         `json:"method"`
			Params *ConsoleParams `json:"params"`
		}
		type ConsoleParams struct {
			Message *gcdapi.ConsoleConsoleMessage `json:"message"`
		}
		msg := &EventData{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			log.Fatalf("error unmarshalling event data: %v\n", err)
		}
		fmt.Printf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		fmt.Printf("Got event: %v\n", eventData)
		fmt.Printf("Timestamp: %f\n", *eventData.Timestamp)
	})
	console.Enable()
	// recv events
	console.Disable()
```

## Usage
For a full list of api methods, types, event types & godocs: [Documentation](https://godoc.org/github.com/wirepair/gcd/gcdapi)

Loading a page using the Page API.
```Go
package main

import (
	"github.com/wirepair/gcd"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	debugger := gcd.NewChromeDebugger()
	// start process, specify a tmp profile path so we get a fresh profiled browser
	// set port 9222 as the debug port
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess() // exit when done
	targets := debugger.GetTargets() // get the 'targets' or tabs/background processes
	target := targets[0] // take the first one

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, v []byte) {
		wg.Done() // page loaded, we can exit now
		// if you wanted to inspect the full response data, you could do that here, or unmarshal it 
		// if you know the proper type
	})
	page := target.Page() // get the Page API
	page.Enable() // Enable so we can recv events
	ret, err := page.Navigate("http://www.veracode.com") // navigate
	if err != nil {
		log.Fatalf("Error navigating: %s\n", err)
	}
	log.Printf("ret: %#v\n", ret)
	wg.Wait() // wait for page load
}

```
Getting a document from a target.
```Go
package main

import (
	"github.com/wirepair/gcd"
	"log"
)

func main() {
	debugger := gcd.NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	targets := debugger.GetTargets()
	target := targets[0]
	dom := target.DOM()
	r, err := dom.GetDocument()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	log.Printf("NodeID: %d Node Name: %s, URL: %s\n", *r.NodeId, r.NodeName, r.DocumentURL)
}
```

Create a new tab and start working with it then close it.
```Go
package main

import (
	"github.com/wirepair/gcd"
	"time"
)

func main() {
	debugger := gcd.NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	target := debugger.NewTab()
	page := target.Page()
	page.Navigate("https://github.com/wirepair/gcd")
	time.Sleep(5 * time.Second)
	debugger.CloseTab(target)
}
```
Create 5 tabs, navigate to URLs, wait for them to load, take screen shot
```Go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/wirepair/gcd"
	"net/url"
	"os"
	"sync"
	"time"
)

const (
	numTabs = 5
)

var debugger *gcd.Gcd
var wg sync.WaitGroup

func main() {

	urls := []string{"http://www.google.com", "http://www.veracode.com", "http://www.microsoft.com", "http://bbc.co.uk", "http://www.reddit.com/r/golang"}
	debugger = gcd.NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	targets := make([]*gcd.ChromeTarget, numTabs)
	for i := 0; i < numTabs; i++ {
		wg.Add(1)
		targets[i] = debugger.NewTab()
		page := targets[i].Page()
		page.Enable()
		targets[i].Subscribe("Page.loadEventFired", PageLoaded)
		page.Navigate(urls[i])
	}
	wg.Wait()
	for i := 0; i < numTabs; i++ {
		TakeScreenShot(targets[i])
	}
}

func PageLoaded(target *gcd.ChromeTarget, event []byte) {
	wg.Done()
}

func TakeScreenShot(target *gcd.ChromeTarget) {
	dom := target.DOM()
	page := target.Page()
	doc, err := dom.GetDocument()
	if err != nil {
		fmt.Errorf("error getting doc: %s\n", err)
		return
	}
	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Errorf("error parsing url: %s\n", urlErr)
		return
	}
	fmt.Printf("Taking screen shot of: %s\n", u.Host)
	img, errCap := page.CaptureScreenshot()
	if errCap != nil {
		fmt.Errorf("error getting doc: %s\n", errCap)
		return
	}
	imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	if errDecode != nil {
		fmt.Errorf("error decoding image: %s\n", errDecode)
		return
	}
	f, errFile := os.Create(u.Host + ".png")
	defer f.Close()
	if errFile != nil {
		fmt.Errorf("error creating image file: %s\n", errFile)
		return
	}
	f.Write(imgBytes)
}
```

## Licensing
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
