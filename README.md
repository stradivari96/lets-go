# lets-go

## Notes

- Code living in the "internal" directory cannot be imported by code outside of the module.
- HTTP requests are served in their own goroutine, https://www.alexedwards.net/blog/understanding-mutexes
- Closure pattern example: https://gist.github.com/alexedwards/5cd712192b4831058b21

## Used commands

```
go mod init snippetbox.xiang.es
go run ./cmd/web
```

```
go get github.com/go-sql-driver/mysql
go mod tidy -v
```
