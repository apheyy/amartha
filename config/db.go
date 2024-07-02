package config

import (
	"amarthaTest/model"
	"log"
	"time"

	"xorm.io/xorm"
)

func InitDB(conf model.AppConfig) *xorm.Engine {
	db, err := xorm.NewEngine(conf.Database.Driver, conf.Database.Credential)
	if err != nil {
		log.Printf("fail to open db connection with errors: %v,\n", err)
		return db
	} else {
		err = db.Ping()
		if err != nil {
			log.Printf("error ping db connection with errors: %v,\n", err)
			return db
		}
	}
	db.ShowSQL(true)
	db.SetMaxOpenConns(conf.Database.MaxOpenConn)
	db.SetMaxIdleConns(conf.Database.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(conf.Database.MaxLifeTime))
	return db
}
