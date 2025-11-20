package concurrency

import (
	"time"
)

// Тип для анонимного метода
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // init Chan

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)} // write to Chan
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // read from Chan
		results[r.string] = r.bool
	}

	return results
}

func MockChecker(url string) bool {
	return url != "rdp://false.com"
}

func SlowMockChecker(url string) bool {
	time.Sleep(1 * time.Second)
	return url != "rdp://false.com"
}
