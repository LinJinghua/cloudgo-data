package entitiesorm

import (
	"os"
	// _ : for init
    _ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	// "github.com/go-xorm/core"
	// "time"
	"log"
)

var orm *xorm.Engine
var logFile *os.File

func init() {
    var err error
	if orm, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"); err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	// orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	// orm.SetMapper(core.SameMapper{})

	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// orm.SetDefaultCacher(cacher)
	// if logFile, err = os.Create("mysql.log"); err != nil {
	// 	panic(err)
	// }
	// orm.SetLogger(xorm.NewSimpleLogger(logFile))
	// defer orm.Close()
	orm.ShowSQL(true)
}
