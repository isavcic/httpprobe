package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/ogier/pflag"
)

func parseOptions() (options, error) {
	opts := options{}
	pflag.IntVarP(&opts.timeout, "timeout", "t", 2000, "Request timeout in milliseconds")
	pflag.IntVarP(&opts.retries, "retries", "r", 2, "Retries")
	pflag.IntVarP(&opts.backoff, "backoff", "b", 1000, "Exponential backoff interval in milliseconds")
	pflag.Parse()
	opts.url = pflag.Arg(0)
	if opts.url == "" {
		return options{}, errors.New("parseOptions: URL not given")
	}
	if !strings.Contains(opts.url, "http") {
		opts.url = "http://" + opts.url
	}
	return opts, nil
}

func makeGETRequest(opts options) bool {
	client := retryablehttp.NewClient()
	client.HTTPClient.Timeout = time.Duration(opts.timeout) * time.Millisecond
	client.RetryMax = opts.retries
	client.RetryWaitMin = time.Duration(opts.backoff) * time.Millisecond
	req, err := retryablehttp.NewRequest("GET", opts.url, nil)
	if err != nil {
		fmt.Println(err)
		// os.Exit(2)
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		// os.Exit(2)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		// os.Exit(2)
		return false
	}
	return true
}

func main() {
	opts, err := parseOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if makeGETRequest(opts) {
		os.Exit(0)
	} else {
		os.Exit(2)
	}
}

type options struct {
	timeout int
	retries int
	backoff int
	url     string
}
