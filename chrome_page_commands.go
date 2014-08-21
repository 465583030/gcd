// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Page commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Page() *ChromePage {
	if c.page == nil {
		c.page = newChromePage(c)
	}
	return c.page
}

type ChromePage struct {
	target *ChromeTarget
}

func newChromePage(target *ChromeTarget) *ChromePage {
	c := &ChromePage{target: target}
	return c
}

// Enables page domain notifications.
func (c *ChromePage) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.enable"})
}

// Disables page domain notifications.
func (c *ChromePage) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.disable"})
}

// Clears the overriden device metrics.
func (c *ChromePage) ClearDeviceMetricsOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearDeviceMetricsOverride"})
}

// Requests that scroll offsets and page scale factor are reset to initial values.
func (c *ChromePage) ResetScrollAndPageScaleFactor() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.resetScrollAndPageScaleFactor"})
}

// Clears the overriden Geolocation Position and Error.
func (c *ChromePage) ClearGeolocationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearGeolocationOverride"})
}

// Clears the overridden Device Orientation.
func (c *ChromePage) ClearDeviceOrientationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearDeviceOrientationOverride"})
}

// Stops sending each frame in the <code>screencastFrame</code>.
func (c *ChromePage) StopScreencast() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.stopScreencast"})
}

// removeScriptToEvaluateOnLoad -
// identifier -
func (c *ChromePage) RemoveScriptToEvaluateOnLoad(identifier *types.ChromePageScriptIdentifier) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["identifier"] = identifier
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.removeScriptToEvaluateOnLoad", Params: paramRequest})
}

// reload - Reloads given page optionally ignoring the cache.
// ignoreCache - If true, browser cache is ignored (as if the user pressed Shift+refresh).
// scriptToEvaluateOnLoad - If set, the script will be injected into all frames of the inspected page after reload.
// scriptPreprocessor - Script body that should evaluate to function that will preprocess all the scripts before their compilation.
func (c *ChromePage) Reload(ignoreCache bool, scriptToEvaluateOnLoad string, scriptPreprocessor string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["ignoreCache"] = ignoreCache
	paramRequest["scriptToEvaluateOnLoad"] = scriptToEvaluateOnLoad
	paramRequest["scriptPreprocessor"] = scriptPreprocessor
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.reload", Params: paramRequest})
}

// navigateToHistoryEntry - Navigates current page to the given history entry.
// entryId - Unique id of the entry to navigate to.
func (c *ChromePage) NavigateToHistoryEntry(entryId int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["entryId"] = entryId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.navigateToHistoryEntry", Params: paramRequest})
}

// deleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *ChromePage) DeleteCookie(cookieName string, url string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.deleteCookie", Params: paramRequest})
}

// setDocumentContent - Sets given markup as the document's HTML.
// frameId - Frame id to set HTML for.
// html - HTML content to set.
func (c *ChromePage) SetDocumentContent(frameId *types.ChromePageFrameId, html string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["html"] = html
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDocumentContent", Params: paramRequest})
}

// setDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
// scale - Scale to apply to resulting view image. Ignored in |fitWindow| mode.
// offsetX - X offset to shift resulting view image by. Ignored in |fitWindow| mode.
// offsetY - Y offset to shift resulting view image by. Ignored in |fitWindow| mode.
func (c *ChromePage) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 8)
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["deviceScaleFactor"] = deviceScaleFactor
	paramRequest["mobile"] = mobile
	paramRequest["fitWindow"] = fitWindow
	paramRequest["scale"] = scale
	paramRequest["offsetX"] = offsetX
	paramRequest["offsetY"] = offsetY
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDeviceMetricsOverride", Params: paramRequest})
}

// setShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *ChromePage) SetShowPaintRects(result bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["result"] = result
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowPaintRects", Params: paramRequest})
}

// setShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *ChromePage) SetShowDebugBorders(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowDebugBorders", Params: paramRequest})
}

// setShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *ChromePage) SetShowFPSCounter(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowFPSCounter", Params: paramRequest})
}

// setContinuousPaintingEnabled - Requests that backend enables continuous painting
// enabled - True for enabling cointinuous painting
func (c *ChromePage) SetContinuousPaintingEnabled(enabled bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setContinuousPaintingEnabled", Params: paramRequest})
}

// setShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *ChromePage) SetShowScrollBottleneckRects(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowScrollBottleneckRects", Params: paramRequest})
}

// setScriptExecutionDisabled - Switches script execution in the page.
// value - Whether script execution should be disabled in the page.
func (c *ChromePage) SetScriptExecutionDisabled(value bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setScriptExecutionDisabled", Params: paramRequest})
}

// setGeolocationOverride - Overrides the Geolocation Position or Error.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *ChromePage) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setGeolocationOverride", Params: paramRequest})
}

// setDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *ChromePage) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["alpha"] = alpha
	paramRequest["beta"] = beta
	paramRequest["gamma"] = gamma
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDeviceOrientationOverride", Params: paramRequest})
}

// setTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
func (c *ChromePage) SetTouchEmulationEnabled(enabled bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setTouchEmulationEnabled", Params: paramRequest})
}

// setEmulatedMedia - Emulates the given media for CSS media queries.
// media - Media type to emulate. Empty string disables the override.
func (c *ChromePage) SetEmulatedMedia(media string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["media"] = media
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setEmulatedMedia", Params: paramRequest})
}

// startScreencast - Starts sending each frame using the <code>screencastFrame</code> event.
// format - Image compression format.
// quality - Compression quality from range [0..100].
// maxWidth - Maximum screenshot width.
// maxHeight - Maximum screenshot height.
func (c *ChromePage) StartScreencast(format string, quality int, maxWidth int, maxHeight int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["format"] = format
	paramRequest["quality"] = quality
	paramRequest["maxWidth"] = maxWidth
	paramRequest["maxHeight"] = maxHeight
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.startScreencast", Params: paramRequest})
}

// handleJavaScriptDialog - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
// accept - Whether to accept or dismiss the dialog.
// promptText - The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (c *ChromePage) HandleJavaScriptDialog(accept bool, promptText string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["accept"] = accept
	paramRequest["promptText"] = promptText
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.handleJavaScriptDialog", Params: paramRequest})
}

// setShowViewportSizeOnResize - Paints viewport size upon main frame resize.
// show - Whether to paint size or not.
// showGrid - Whether to paint grid as well.
func (c *ChromePage) SetShowViewportSizeOnResize(show bool, showGrid bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["show"] = show
	paramRequest["showGrid"] = showGrid
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowViewportSizeOnResize", Params: paramRequest})
}

// getNavigationHistory - Returns navigation history for the current page.
// Returns -
// Index of the current navigation history entry.
// Array of navigation history entries.
func (c *ChromePage) GetNavigationHistory() (float64, []*types.ChromePageNavigationEntry, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getNavigationHistory"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CurrentIndex float64
			Entries      []*types.ChromePageNavigationEntry
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return 0, nil, &ChromeRequestErr{Resp: cerr}
		}
		return 0, nil, err
	}

	return chromeData.Result.CurrentIndex, chromeData.Result.Entries, nil
}

// getCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -
// Array of cookie objects.
func (c *ChromePage) GetCookies() ([]*types.ChromePageCookie, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Cookies []*types.ChromePageCookie
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

	return chromeData.Result.Cookies, nil
}

// getResourceTree - Returns present frame / resource tree structure.
// Returns -
// Present frame / resource tree structure.
func (c *ChromePage) GetResourceTree() (*types.ChromePageFrameResourceTree, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getResourceTree"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameTree *types.ChromePageFrameResourceTree
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

	return chromeData.Result.FrameTree, nil
}

// getScriptExecutionStatus - Determines if scripts can be executed in the page.
// Returns -
// Script execution status: "allowed" if scripts can be executed, "disabled" if script execution has been disabled through page settings, "forbidden" if script execution for the given page is not possible for other reasons.
func (c *ChromePage) GetScriptExecutionStatus() (string, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getScriptExecutionStatus"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.Result, nil
}

// captureScreenshot - Capture page screenshot.
// Returns -
// Base64-encoded image data (PNG).
func (c *ChromePage) CaptureScreenshot() (string, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.captureScreenshot"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.Data, nil
}

// canScreencast - Tells whether screencast is supported.
// Returns -
// True if screencast is supported.
func (c *ChromePage) CanScreencast() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.canScreencast"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}

// addScriptToEvaluateOnLoad -
// Returns -
// Identifier of the added script.
func (c *ChromePage) AddScriptToEvaluateOnLoad(scriptSource string) (*types.ChromePageScriptIdentifier, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scriptSource"] = scriptSource
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.addScriptToEvaluateOnLoad", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Identifier *types.ChromePageScriptIdentifier
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

	return chromeData.Result.Identifier, nil
}

// navigate - Navigates current page to the given URL.
// Returns -
// Frame id that will be navigated.
func (c *ChromePage) Navigate(url string) (*types.ChromePageFrameId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.navigate", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameId *types.ChromePageFrameId
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

	return chromeData.Result.FrameId, nil
}

// getResourceContent - Returns content of the given resource.
// Returns -
// Resource content.
// True, if content was served as base64.
func (c *ChromePage) GetResourceContent(frameId *types.ChromePageFrameId, url string) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getResourceContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Content       string
			Base64Encoded bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", false, &ChromeRequestErr{Resp: cerr}
		}
		return "", false, err
	}

	return chromeData.Result.Content, chromeData.Result.Base64Encoded, nil
}

// searchInResource - Searches for given string in resource content.
// Returns -
// List of search matches.
func (c *ChromePage) SearchInResource(frameId *types.ChromePageFrameId, url string, query string, caseSensitive bool, isRegex bool) ([]*types.ChromePageSearchMatch, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	paramRequest["query"] = query
	paramRequest["caseSensitive"] = caseSensitive
	paramRequest["isRegex"] = isRegex
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.searchInResource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result []*types.ChromePageSearchMatch
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

	return chromeData.Result.Result, nil
}

// queryUsageAndQuota - Queries more detailed quota and usage data than Storage API provides.
// Returns -
// Quota for requested security origin.
// Current usage for requested security origin.
func (c *ChromePage) QueryUsageAndQuota(securityOrigin string) (*types.ChromePageQuota, *types.ChromePageUsage, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["securityOrigin"] = securityOrigin
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.queryUsageAndQuota", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Quota *types.ChromePageQuota
			Usage *types.ChromePageUsage
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, err
	}

	return chromeData.Result.Quota, chromeData.Result.Usage, nil
}
