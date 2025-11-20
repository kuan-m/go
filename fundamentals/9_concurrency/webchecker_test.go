package concurrency

import (
	"fmt"
	"testing"
)

var websites = []string{
	"http://google.com",
	"http://google1.com",
	"http://google2.com",
	"http://google3.com",
	"http://google4.com",
	"http://google5.com",
	"http://google6.com",
	"http://google.7com",
	"http://google8.com",
	"http://google9.com",
	"http://google12.com",
	"http://google31.com",
	"http://google31543534.com",
	"rdp://false.com",
}

func TestMockChecker(t *testing.T) {
	fmt.Println(CheckWebsites(MockChecker, websites))
}

func TestSlowMockChecker(t *testing.T) {
	fmt.Println(CheckWebsites(SlowMockChecker, websites))
}
