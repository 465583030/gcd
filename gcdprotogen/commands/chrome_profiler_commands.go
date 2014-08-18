// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Profiler commands.
// API Version: 1.1

package gcd


import (
	
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Profiler() *ChromeProfiler {
	if c.profiler == nil {
		c.profiler = newChromeProfiler(c)
	}
	return c.profiler
}


type ChromeProfiler struct {
	target *ChromeTarget
}

func newChromeProfiler(target *ChromeTarget) *ChromeProfiler {
	c := &ChromeProfiler{target: target}
	return c
}

// start non parameterized commands 
// 
func (c *ChromeProfiler) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.enable"})
}
 
// 
func (c *ChromeProfiler) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.disable"})
}
 
// 
func (c *ChromeProfiler) Start() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.start"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setSamplingInterval - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
// interval - New sampling interval in microseconds.
func (c *ChromeProfiler) SetSamplingInterval(interval int, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["interval"] = interval
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.setSamplingInterval", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// stop - 
// Returns - 
// Recorded profile.
func (c *ChromeProfiler) Stop() (*types.ChromeProfilerCPUProfile, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.stop"})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Profile *types.ChromeProfilerCPUProfile 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Profile, nil
}


// end commands with no parameters but special return types


// start commands with parameters and special return types


// end commands with parameters and special return types

