/**
 * @author XieKong
 * @date   2019/5/13 18:42
 */
package modules

import (
	"log"
	"novel-reader/db"
	"time"
)

// 源条目数据模型
type Source struct {
	Id      int64
	Name    string
	Address     string    `xorm:"varchar(255)"`
	Created time.Time `xorm:"created"`
}

func init() {
	err := db.Engine.Sync2(new(Source))

	if err != nil {
		log.Fatal(err)
	}
}

func ListSource() []*Source {
	result := make([]*Source, 0)

	err := db.Engine.Find(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func InsertSource(source *Source) error {
	_, err := db.Engine.Insert(source)

	return err
}

func UpdateSource(source *Source) error {
	_, err := db.Engine.Id(source.Id).Update(source)

	return err
}

func HasSource(source *Source) (bool, error) {
	has, err := db.Engine.Get(source)
	return has, err
}

func DeleteSource(source *Source) error {
	_, err := db.Engine.Delete(source)

	return err
}