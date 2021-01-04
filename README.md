## Try this!

1. [install golang 1.16-beta1](https://groups.google.com/g/golang-announce/c/2-Rj3P5uRLs/m/mYxD2RJkAQAJ)
2. clone this repo
3. run `go1.16beta1 run main.go`
4. open localhost:8083
5. see the contents of `hello.txt` served

Sadly because the file is embeded at compile time(1) changes to the file won't be reflected in the server response; the service must be restarted on file changes.

(1) needs citation
