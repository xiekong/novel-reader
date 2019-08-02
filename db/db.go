/**
 * @author XieKong
 * @date   2019/7/29 14:46
 */
package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"novel-reader/utils"
)

var Engine *xorm.Engine

func init() {
	initSqlite3()
}

func initSqlite3() {
	var err error
	Engine, err = xorm.NewEngine(utils.GetDriverName(), utils.GetDataSource());

	if err != nil {
		utils.Logger.Error(err)
	}

	err = Engine.Ping()

	if err != nil {
		utils.Logger.Error(err)
	}
}
