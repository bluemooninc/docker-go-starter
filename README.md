# docker-go-starter

You can quick start for GO language development with docker.

You can edit go file which is automatically hot reload by fresh module.

Sample main.go is simple api just response json data. 

if you want to more variety for Dockerfile visit below.

https://github.com/docker-library/golang

## 1. Starting docker
Docke compose up and start API server. The sample endpoint is http://localhost:8080/sample
When you edit the main.go, it will chang by fresh module has hot reload automatically.
```cassandraql
docker-compose up
Creating go ... done
Attaching to go
go     | 8:9:59 runner      | InitFolders
go     | 8:9:59 runner      | mkdir ./tmp
go     | 8:9:59 runner      | mkdir ./tmp: file exists
go     | 8:9:59 watcher     | Watching .
go     | 8:9:59 main        | Waiting (loop 1)...
go     | 8:9:59 main        | receiving first event /
go     | 8:9:59 main        | sleeping for 600 milliseconds
go     | 8:10:00 main        | flushing events
go     | 8:10:00 main        | Started! (5 Goroutines)
go     | 8:10:00 main        | remove tmp/runner-build-errors.log: no such file or directory
go     | 8:10:00 build       | Building...
go     | 8:10:04 runner      | Running...
go     | 8:10:04 main        | --------------------
go     | 8:10:04 main        | Waiting (loop 2)...
go     | 8:10:04 app         |
go     |    ____    __
go     |   / __/___/ /  ___
go     |  / _// __/ _ \/ _ \
go     | /___/\__/_//_/\___/ v3.3.10-dev
go     | High performance, minimalist Go web framework
go     | https://echo.labstack.com
go     | ____________________________________O/_______
go     |                                     O\

``` 

##2. CHeck by Curl
Call api from your local terminal.
```cassandraql
$ curl http://localhost:8080/sample
{"code":200,"text":"OK"}
```
