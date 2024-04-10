package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标(os.Stdout为控制台)，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second * 15, // 慢 SQL 阈值
			LogLevel:                  logger.Info,      // 日志级别
			IgnoreRecordNotFoundError: true,             // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,             // 禁用彩色打印
		},
	)

	//var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	Config.MySqlInfo.Username,
	//	Config.MySqlInfo.Password,
	//	Config.MySqlInfo.Host,
	//	Config.MySqlInfo.Port,
	//	Config.MySqlInfo.Database,
	//)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"123456",
		"192.168.31.124",
		3306,
		"mxshop_user_srv",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("db connect error")
	}
	DB = db
}

// Paginate 分页查询公共方法
func Paginate(pageSize, pageNum int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
