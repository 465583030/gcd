// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the FileSystem commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) FileSystem() *ChromeFileSystem {
	if c.filesystem == nil {
		c.filesystem = newChromeFileSystem(c)
	}
	return c.filesystem
}

type ChromeFileSystem struct {
	target *ChromeTarget
}

func newChromeFileSystem(target *ChromeTarget) *ChromeFileSystem {
	c := &ChromeFileSystem{target: target}
	return c
}

// Enables events from backend.
func (c *ChromeFileSystem) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.enable"})
}

// Disables events from backend.
func (c *ChromeFileSystem) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.disable"})
}

// requestFileSystemRoot - Returns root directory of the FileSystem, if exists.
// Returns -
// 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value.
// Contains root of the requested FileSystem if the command completed successfully.
func (c *ChromeFileSystem) RequestFileSystemRoot(origin string, theType string) (float64, *types.ChromeFileSystemEntry, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["origin"] = origin
	paramRequest["type"] = theType
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.requestFileSystemRoot", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode float64
			Root      *types.ChromeFileSystemEntry
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Root, nil
}

// requestDirectoryContent - Returns content of the directory.
// Returns -
// 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value.
// Contains all entries on directory if the command completed successfully.
func (c *ChromeFileSystem) RequestDirectoryContent(url string) (float64, []*types.ChromeFileSystemEntry, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.requestDirectoryContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode float64
			Entries   []*types.ChromeFileSystemEntry
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Entries, nil
}

// requestMetadata - Returns metadata of the entry.
// Returns -
// 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value.
// Contains metadata of the entry if the command completed successfully.
func (c *ChromeFileSystem) RequestMetadata(url string) (float64, *types.ChromeFileSystemMetadata, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.requestMetadata", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode float64
			Metadata  *types.ChromeFileSystemMetadata
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Metadata, nil
}

// requestFileContent - Returns content of the file. Result should be sliced into [start, end).
// Returns -
// 0, if no error. Otherwise, errorCode is set to FileError::ErrorCode value.
// Content of the file.
// Charset of the content if it is served as text.
func (c *ChromeFileSystem) RequestFileContent(url string, readAsText bool, start int, end int, charset string) (float64, string, string, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["url"] = url
	paramRequest["readAsText"] = readAsText
	paramRequest["start"] = start
	paramRequest["end"] = end
	paramRequest["charset"] = charset
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.requestFileContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode float64
			Content   string
			Charset   string
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, "", "", &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, "", "", err
	}

	return chromeData.Result.ErrorCode, chromeData.Result.Content, chromeData.Result.Charset, nil
}

// deleteEntry - Deletes specified entry. If the entry is a directory, the agent deletes children recursively.
// Returns -
// 0, if no error. Otherwise errorCode is set to FileError::ErrorCode value.
func (c *ChromeFileSystem) DeleteEntry(url string) (float64, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.deleteEntry", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ErrorCode float64
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.ErrorCode, nil
}
