package Model

import (
	"ZongzhiCui/go_gin/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("mysql", config.Mysql_conn())
	//if err != nil {
	//	panic("failed to connect database")
	//}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//db.DB().SetMaxIdleConns(10) //设置连接池最大闲置数量

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	//db.DB().SetMaxOpenConns(100) //设置连接池最大打开数量

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//db.DB().SetConnMaxLifetime(time.Hour)

	db.LogMode(true) //开启sql debug 模式
}

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index;not null"` // 类似 索引，创建一个唯一的索引 不知道怎么用 - 20200522
	Password string
}

// 设置 `User` 的表名为 `use`
func (u User) TableName() string {
	return "user"
}

func User_info(user string) User {
	//defer db.Close()

	// 查
	var product User
	//fmt.Println(db.HasTable(product)) //表是否存在

	//db.First(&product, 1) // 找到id为1的产品
	//db.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品
	db.First(&product, "name = ?", user)

	return product
}

func User_create(u User) bool {
	//defer db.Close()

	db.Create(&u)

	b := db.NewRecord(u)
	return b
}
