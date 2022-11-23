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

The results are:
```
$ ls -lh
-rwxrwxrwx 1 chuqq chuqq 2.7K Nov 23 17:31 myapp-2022-11-23T17-31-26.637.log.gz*
-rwxrwxrwx 1 chuqq chuqq 2.7K Nov 23 17:32 myapp-2022-11-23T17-32-03.204.log.gz*
-rwxrwxrwx 1 chuqq chuqq 2.7K Nov 23 17:32 myapp-2022-11-23T17-32-39.780.log.gz*
-rwxrwxrwx 1 chuqq chuqq 196K Nov 23 17:32 myapp.log*
```
