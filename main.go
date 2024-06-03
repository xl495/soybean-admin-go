package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/router"
	"time"
)

//	@title			soybean-admin-go Admin App
//	@version		1.0
//	@description	This is an API for soybean-admin-go Application
//	@contact.name	hxl
//	@contact.email	52553624@qq.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath	/api
func main() {

	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc

	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 连接数据库
	err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	app := fiber.New()

	router.Initalize(app)

	// 从环境变量中获取端口号
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal(app.Listen(port))
}
