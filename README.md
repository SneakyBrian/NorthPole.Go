# NorthPole.Go

## A HTTP Post Data Signing and Verification Service ##

North Pole provides 2 HTTP end points that except any arbitrary POST data.

`/elf` is the end point that signs the data with a unix epoch timestamp in milliseconds and a HMAC

`/santa` is the end point that verifies the values returned by `/elf`

`/santa` responds with the age in milliseconds of the data if it is valid, or -1 if it is invalid

## Examples

`/elf` request (using curl):

```
curl "http://localhost:8181/elf/?key=some"%"20unique"%"20key" -H "Origin: http://localhost:8181" -H "Accept-Encoding: gzip, deflate" -H "Accept-Language: en-US,en;q=0.8" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" -H "Accept: text/plain, */*; q=0.01" -H "Referer: http://localhost:8181/static/test.html" -H "Connection: keep-alive" --data "This is some test data" --compressed
```

`/elf` response:

```
{"Hash":"cdX3dqChEMW4mfB-3b1cRiJF9xdRvYwOKx0_9iSAPAg=","Timestamp":1464293266949}
```

`/santa` request (using curl):

```
curl "http://localhost:8181/santa/?key=some"%"20unique"%"20key&timestamp=1464293266949&hash=cdX3dqChEMW4mfB-3b1cRiJF9xdRvYwOKx0_9iSAPAg"%"3D" -H "Origin: http://localhost:8181" -H "Accept-Encoding: gzip, deflate" -H "Accept-Language: en-US,en;q=0.8" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" -H "Accept: text/plain, */*; q=0.01" -H "Referer: http://localhost:8181/static/test.html" -H "Connection: keep-alive" --data "This is some test data" --compressed
```

`/santa` response:

`3627`

## Getting

Get the source:

`go get github.com/sneakybrian/northpole.go`

## Building

From within the source directory:

`go generate`

Embeds the contents of the ./static directory into the executable binary

`go build`

Builds the executable binary

And optionally:

`go install`

To install to your $GOPATH$\bin directory

## Development

The program has a test page available at:

http://localhost:8181/static/test.html

The program also has a configurable port number for the HTTP Server:

`NorthPole.Go -port=1234`

Would run the web server at:

http://localhost:1234/

Run the program with the `-useLocal` flag in order to serve the static resources from the file system rather than the embedded resources:

`NorthPole.Go -useLocal`

This allows you to edit the contents of the ./static folder on-the-fly and be able to refresh the browser to get the updated changes

## Packages

Uses the excellent `esc` package for generating the embedded static resources:

https://github.com/mjibson/esc