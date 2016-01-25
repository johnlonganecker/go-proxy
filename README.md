# go-proxy
A Simple Reverse Proxy Server in golang that can intercept calls

## start up simple http server
```bash
cd to-folder-with-files
python -m SimpleHTTPServer
```

## start up our reverse proxy
```bash
go run main.go
```

Request to `/sayblah` are intercepted and return with the word BLAH

All other requests are passed through to our python simple http server
