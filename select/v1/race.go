package v1

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	aDuration := measureResponse(a)

	bDuration := measureResponse(b)

	if aDuration > bDuration {
		return b
	}
	return a
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
