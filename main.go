/**
 * @author XieKong
 * @date   2019/7/29 14:16
 */
package main

import (
	_ "novel-reader/db"
	"novel-reader/ui"
	"novel-reader/utils"
)

const debug = true;

func init() {
	utils.Init(debug)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			// 捕获的异常处理
			utils.Logger.Error(r)
		}
	}()

	ui.Show()
}
