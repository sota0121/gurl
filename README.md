# cURL emulator in Go

This is a simple cURL emulator written in Go. It is not intended to be a full replacement for cURL, but rather a simple tool to demonstrate how to use the Go standard library to make HTTP requests.

## Usage

```
$ gurl -h
Usage: gurl [options...] <url>
 -d, --data <data>          HTTP POST data
 -f, --fail                 Fail silently (no output at all) on HTTP errors
 -h, --help <category>      Get help for commands
 -i, --include              Include protocol response headers in the output
 -o, --output <file>        Write to file instead of stdout
 -O, --remote-name          Write output to a file named as the remote file
 -s, --silent               Silent mode
 -T, --upload-file <file>   Transfer local FILE to destination
 -u, --user <user:password> Server user and password
 -A, --user-agent <name>    Send User-Agent <name> to server
 -v, --verbose              Make the operation more talkative
 -V, --version              Show version number and quit
```

## Examples

```
$ gurl -v https://httpbin.org/get
GET /get HTTP/1.1
Host: httpbin.org
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
```


