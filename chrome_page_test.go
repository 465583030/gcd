package gcd

import (
	//"encoding/json"
	"fmt"
	//"log"
	"testing"
	"time"
)

func TestPageFPS(t *testing.T) {
	debugger := NewChromeDebugger()
	debugger.StartProcess("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "C:\\tmp\\", "9222")
	defer debugger.ExitProcess()
	time.Sleep(1 * time.Second)
	targets := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("page: %s\n", targets[0].Target.Url)
	page := targets[0].Page()

	fmt.Printf("Sending enable...")
	_, err := page.SetShowFPSCounter(true)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	cookies, newErr := page.GetCookies()
	fmt.Printf("%v %v\n", cookies, newErr)
}
