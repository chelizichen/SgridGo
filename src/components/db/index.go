package component_db

import (
	"com_sgrid/src/framework/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func LoadDB(conf *config.SgridConf) {
	db_master := conf.GetString("db")
	db, err := gorm.Open(mysql.Open(db_master), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "trade_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	DB = (db)
}
