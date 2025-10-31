// 封装mysql数据库相关操作
package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/kevinjosephdavis/chatroom/internal/config"
	"github.com/kevinjosephdavis/chatroom/internal/repository/model"
)

var MySQL *gorm.DB

func InitMySQL(cfg *config.Config) error {
	var err error

	//连接 MySQL
	MySQL, err = gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("连接MySQL失败: %v", err)
	}

	sqlDB, err := MySQL.DB()
	if err != nil {
		return err
	}

	//设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//自动迁移表结构
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("自动迁移失败: %v", err)
	}

	log.Println("MySQL 连接成功")
	return nil
}

func autoMigrate() error {
	return MySQL.AutoMigrate(
		&model.User{},
		&model.Friend{},
		//&model.Group{},
		//&model.GroupMember{},
		//&model.Message{},
	)
}

func GetMySQL() *gorm.DB {
	return MySQL
}
