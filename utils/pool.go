package utils

import (
	"fmt"
	"github.com/chenyu116/postgres-server/config"
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	PgPoolRead  sync.Pool
	PgPoolWrite sync.Pool
)

func InitPool() {
	_config := config.GetConfig()
	PgPoolRead = sync.Pool{
		New: func() interface{} {
			dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable", _config.DbServer.ReadHost, _config.DbServer.Username, _config.DbServer.Password, _config.DbServer.ReadPort, _config.DbServer.Database)
			db, err := xorm.NewEngine("postgres", dsn)
			if err != nil {
				fmt.Println(err)
			}
			return db
		}}
	PgPoolWrite = sync.Pool{New: func() interface{} {
		db, err := xorm.NewEngine("postgres", fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable", _config.DbServer.WriteHost, _config.DbServer.Username, _config.DbServer.Password, _config.DbServer.WritePort, _config.DbServer.Database))
		if err != nil {
			fmt.Println(err)
		}
		return db
	}}
}
