# about this project

demo for https://medium.com/@erickrazed/gin-bufio-writer-with-404-handler-7ca1b1ce5b83


you can try different version of go and/or gin, bufio. but I think there is no breaking change about this

## usage

simply

```
go run demo1/main.go
```

to test
```
curl localhost:3000/demo/notFound -i
```

result

```
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Fri, 12 Aug 2022 11:18:19 GMT
Content-Length: 0

```

