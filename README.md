```
go build myapp.go
go build logrotate.go

./myapp 2>&1 | ./logrotate config.json &
# or
nohup sh -c "./myapp 2>&1 | ./logrotate config.json" &

tail -F myapp.log
# not
# tail -f myapp.log
```
