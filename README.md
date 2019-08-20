# shorty
url-shortener service in go. The idea is that you trade a long url for a shorter one that's easier to remember. See usage for details.

# usage
run server
```bash
$ go run ./cmd/shorty [-host -port -version -help]
```
make request
```bash
$ curl <host>:<port>/ -d "http://hemnet.com" -s | pbcopy
```
and paste that into a browser window.

# api
```bash
# create short url by passing a long url in body
POST host:port
# visit long url
GET host:port/<id>
# list available short urls
GET host:port/
```

# test
```bash
$ go test
```

# todos
- [x] random key
- [x] create short url
- [x] follow short url
- [x] list stored urls
- [x] store interface
- [x] in-mem store
- [x] http server
- [ ] deploy to do

# license
MIT
