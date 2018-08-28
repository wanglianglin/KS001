# App

A sample 12 Factor Application.

## Prepare steps for Mac OSX
```
$ homebrew install go
```
> latest version of go will be installed. on my side, go 1.11 will be installed.

Add from lines into ~/.bash_profile
```
export GOPATH=$HOME/go
export GOROOT=/usr/local/Cellar/go/1.11/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
## Get the code
```
$ mkdir -p $GOPATH/src/github.com/wanglianglin/
$ cd $GOPATH/src/github.com/wanglianglin/
$ git clone https://github.com/wanglianglin/KS001
$ cd KS001/app
```

## Build and Run

```
$ go build -o server ./monolith
```

```
$ ./server -http :8088
```

```
2016/04/15 06:34:12 Starting server...
2016/04/15 06:34:12 HTTP service listening on 0.0.0.0:80
```

## Test with cURL

```
$ curl http://127.0.0.1/login -u user
```
```
Enter host password for user 'user':
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjE0NjA5ODcxOTcsImlhdCI6MTQ2MDcyNzk5NywiaXNzIjoiYXV0aC5zZXJ2aWNlIiwic3ViIjoidXNlciJ9.x3oFhRhWk5CGYfGcrNctPGWCENEsXpUuKPDQU2ZOLCY
```

> type "password" at the prompt

```
curl -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjE0NjA5ODcxOTcsImlhdCI6MTQ2MDcyNzk5NywiaXNzIjoiYXV0aC5zZXJ2aWNlIiwic3ViIjoidXNlciJ9.x3oFhRhWk5CGYfGcrNctPGWCENEsXpUuKPDQU2ZOLCY' http://127.0.0.1/secure
```
```
<h1>Hello</h1>
```
