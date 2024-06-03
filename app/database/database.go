package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stsyem "log"
	"os"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils"
	"time"
)

var DB *gorm.DB

func ConnectDB() error {

	dbLogger := logger.New(
		stsyem.New(os.Stdout, "\r\n", stsyem.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	)

	var err error // define error here to prevent overshadowing the global DB

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		username, password, host, db,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	initializeData(DB)
	// 如果没有默认用户 则初始化数据
	return err
}

func initializeData(db *gorm.DB) {
	var userCount int64
	db.Model(&model.User{}).Count(&userCount)

	if userCount == 0 {
		pws, err := utils.PasswordHash("123456")

		newUser := &model.User{
			Username: "Soybean",
			Password: pws,
		}

		dbUser := db.Create(newUser)

		if dbUser.Error != nil {
			log.Error("创建默认用户失败 ", dbUser.Error)
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Info("初始化数据成功, 创建默认用户成功, 用户名: %s", newUser.Username)
	}
}
