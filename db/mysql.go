package db

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/fujimisakari/turntable-build/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

func Init() *dbr.Session {
	session := getSession()
	return session
}

func getSession() *dbr.Session {
	uriSchema := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.USER, conf.PASSWORD, conf.HOST, conf.PORT, conf.DB)
	db, err := dbr.Open("mysql", uriSchema, nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
