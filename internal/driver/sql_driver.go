package driver

import (
	mysqlparser "github.com/dxc0522/sql2struct/internal/mysql"
	"github.com/dxc0522/sql2struct/internal/postgresql"
	"github.com/dxc0522/sql2struct/utils"
)

type ModelGenerator interface {
	Run() error
}

func NewSqlDriverGenerator(driverName string) ModelGenerator {

	if driverName == "mysql" {
		return mysqlparser.NewMysqlGenerator()
	}
	if driverName == "postgres" {
		return postgresql.NewPgParser()
	}
	utils.PrintRedf("unsupported driver %s, supported `mysql` `postgres` ", driverName)
	return nil
}
