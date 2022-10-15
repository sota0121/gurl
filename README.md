# cURL emulator in Go

This is a simple cURL emulator written in Go. It is not intended to be a full replacement for cURL, but rather a simple tool to demonstrate how to use the Go standard library to make HTTP requests.

## Usage

```
$ gurl -h
Usage of curl:
  -d string
        HTTP request body
  -h    Show this help message
  -i    Include HTTP response headers
  -m string
        HTTP method (default "GET")
  -o string
        Write HTTP response body to file
  -v    Verbose output
  -x string
        HTTP proxy
  -X string
        HTTP method (default "GET")
```

## Examples

```
$ gurl -v https://httpbin.org/get
GET /get HTTP/1.1
Host: httpbin.org
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
```


