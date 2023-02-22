package app

import (
	"Logi/api/model/scm"
	"Logi/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect() {
	var dbs [6]string
	dbs[0] = helper.Cfg("DB_USER")
	dbs[1] = helper.Cfg("DB_PASSWORD")
	dbs[2] = helper.Cfg("DB_HOST")
	dbs[3] = helper.Cfg("DB_PORT")
	dbs[4] = helper.Cfg("DB_NAME")
	dbs[5] = "?charset=utf8mb4&parseTime=True&loc=Local"

	var dsn = dbs[0] + ":" + dbs[1] + "@tcp(" + dbs[2] + ":" + dbs[3] + ")/" + dbs[4] + dbs[5]
	Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.Pie(err)
	log.Println("Connected to Database...")
}

func Migrate() {
	//delAllTable()
	err = Instance.AutoMigrate(&scm.Member{}, &scm.MemberProfile{}, &scm.BalanceTransaction{})
	helper.Pie(err)
	log.Println("Migrasi database selesai ...")
}

func delAllTable() {
	err := Instance.Migrator().DropTable("balance_members", "members", "member_profiles")
	helper.Pie(err)
}
