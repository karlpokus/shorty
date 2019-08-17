# shorty
url-shortener service in go

# usage
```bash
# create short url
POST host:port/<long-url>
# visit short url
GET host:port/<short-url>
# list all shortened urls
GET host:port/
```

# test
```bash
$ go test
```

# todos
- [x] random sequence
- [ ] shortener pkg
- [ ] store interface
- [ ] http server

# license
MIT
