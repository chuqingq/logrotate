```
go build myapp.go
go build logrotate.go

./myapp 2>&1 | ./logrotate config.json

tail -F myapp.log
```
