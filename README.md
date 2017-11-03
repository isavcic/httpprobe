# httpprobe

A simple HTTP probe, supporting configurable retry number with configurable exponential backoff and per-attempt timeout.

```
Usage of ./httpprobe:
  -b, --backoff int
    	Exponential backoff interval in milliseconds (default 1000)
  -r, --retries int
    	Retries (default 2)
  -t, --timeout int
    	Request timeout in milliseconds (default 2000)
```
