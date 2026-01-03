package main

import (
	_ "bee_demo/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	user, err := beego.AppConfig.String("mysqluser")
	pass, err := beego.AppConfig.String("mysqlpass")
	dataSource, err := beego.AppConfig.String("mysqlurls")
	mysqldb, err := beego.AppConfig.String("mysqldb")
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", user, pass, dataSource, mysqldb)
	fmt.Println(err, user, pass, dataSource, mysqldb, sqlUrl)

	err = orm.RegisterDataBase("default", "mysql", sqlUrl)
	if err != nil {
		panic(err)
	}

	err = orm.RunSyncdb("default", true, false)
	if err != nil {
		panic(err)
	}

}
func main() {
	beego.Run()
}
