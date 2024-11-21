[![Go Report Card](https://goreportcard.com/badge/github.com/dxc0522/sql2struct)](https://goreportcard.com/report/github.com/dxc0522/sql2struct)
[![GoDoc](https://godoc.org/github.com/dxc0522/sql2struct?status.svg)](https://godoc.org/github.com/dxc0522/sql2struct)
![License](https://img.shields.io/badge/license-GPL-blue.svg)
# sql2struct
mysql/postgresql database to golang struct for gorm model

# install
```shell
go install github.com/dxc0522/sql2struct@latest
```

# usage
```shell
# 直接执行
sql2struct --dsn="mysql://root:123456@tcp(localhost:3306)/test?charset=utf8mb4"

# 读取文件执行
# 无dsn则自动读取./etc/config.yaml 文件下的
# DBConfig:
#  DSN: mysql://root:123456@tcp(localhost:3306)/test?charset=utf8mb4
sql2struct -t "user,to_do"
```

#  support
- [x] mysql
- [x] postgreSQL

