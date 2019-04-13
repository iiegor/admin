package admin

import (
	"github.com/go-xorm/xorm"
)

var (
	db  *xorm.Engine
	err error
)

func NewDB(driverName string, dataSourceName string) *xorm.Engine {
	db, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	return db
}
