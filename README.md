# quickserve
CLI for deploying a simple static HTTP test server 

#Installation

You need go to build quickserve
http://golang.org/

```Shell
git clone github.com/madhavgharmalkar/quickserve
cd github.com/madhavgharmalkar/quickserve
go install
```
#Usage

cd into root directory and run
```Shell
quickserve --port 3000 //launches http serve on localhost:3000
```

Example output 
```Shell
quickserve --port 3030
2015/09/11 11:57:22 HTTP server listening on 127.0.0.1:3000
2015/09/11 11:57:28 [GET] "/"
2015/09/11 11:57:28 [GET] "/favicon.ico"
```
