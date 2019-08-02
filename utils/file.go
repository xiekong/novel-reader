/**
 * @author XieKong
 * @date   2019/7/29 15:31
 */
package utils

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var (
	rootPath = "novel-reader"
	configPath = "config"
	logPath = "logs"
	driverName = "sqlite3"
	dataSource = "data.db";
)

func init() {
	if !Exists(GetRootPath()) {
		os.Mkdir(GetRootPath(), os.ModePerm)
	}
	if !Exists(GetConfigPath()) {
		os.Mkdir(GetConfigPath(), os.ModePerm)
	}
	if !Exists(GetLogPath()) {
		os.Mkdir(GetLogPath(), os.ModePerm)
	}
	if !Exists(filepath.Join(GetRootPath(), "db")) {
		os.Mkdir(filepath.Join(GetRootPath(), "db"), os.ModePerm)
	}
}

func GetRootPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(user.HomeDir, rootPath);
}

func GetDriverName() string {
	return driverName;
}

func GetConfigPath() string {
	return filepath.Join(GetRootPath(), configPath)
}


func GetLogPath() string {
	return filepath.Join(GetRootPath(), logPath)
}

func GetDataSource() string {
	return filepath.Join(GetRootPath(), "db", dataSource);
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}