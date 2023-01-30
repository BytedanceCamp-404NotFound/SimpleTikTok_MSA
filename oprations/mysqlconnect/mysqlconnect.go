package mysqlconnect

import (
	"SimpleTikTok/oprations/viperconfigread"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormDB *gorm.DB

func init() {
	var err error
	GormDB, err = SqlConnect()
	if err != nil {
		logx.Errorf("get sql connect fail, err:%v", err)
	}

	sql, err := GormDB.DB()
	if err != nil {
		logx.Errorf("sel sql connpool fail, err:%v", err)
	}
	data, _ := json.Marshal(sql.Stats())
	logx.Infof("sql pool:" + string(data))
	// sql.Close()
}

// 函数功能：连接数据库
// 返回值 *gorm.DB为链接上的数据库
func SqlConnect() (*gorm.DB, error) {
	mysqlConfig, err := viperconfigread.ConfigReadToMySQL()
	if err != nil {
		logx.Errorf("SqlConnect error:%v", err)
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConfig.UserName, mysqlConfig.PassWord, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DBname, mysqlConfig.TimeOut)
	db, err := gormInit(dsn)
	if err != nil {
		logx.Errorf("gorm init fail, error:%v", err.Error())
		return nil, err
	}

	sqlpool, err := db.DB()
	if err != nil {
		logx.Errorf("gorm init fail, error:%v", err.Error())
		return nil, err
	}

	// 设置连接池参数，参数的具体意义可以查看配置文件
	sqlpool.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	sqlpool.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	// sqlpool.Close() //这里看文档关闭sql.db()不会影响gorm.db()

	// err = gormTableInit(db)
	// if err != nil {
	// 	logx.Errorf("init tables fail, error:%v", err.Error())
	// 	return nil, err
	// }

	return db, nil
}

// gorm初始化
func gormInit(dsn string) (*gorm.DB, error) {
	// 日志的配置
	logLevel := logger.Warn
	if true {
		logx.Info("gorm with debug mode")
		logLevel = logger.Info
	}
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	// 配置gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true, // 跳过默认开启事务模式
		PrepareStmt:            false,
		AllowGlobalUpdate:      true, // 在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名，启用该选项.
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// // 通过gorm 创建数据库表
// func gormTableInit(db *gorm.DB) error {
// 	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").AutoMigrate(&VideoInfo{}, &UserInfo{}); err != nil {
// 		logx.Error("opendb fialed", err)
// 		return err
// 	}
// 	return nil
// }
