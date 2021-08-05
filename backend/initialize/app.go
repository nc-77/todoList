package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nc-77/gtils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"todo/global"
	"todo/router"
)

type App struct {
	Config             *TomlConfig
	httpServerInstance *http.Server
}

func NewApp(config *TomlConfig) (app App) {
	app.Config = config
	// 初始化httpServer
	app.httpServerInstance = &config.HS
	return
}

func (app *App) openDB() {
	cfg := app.Config.DB
	var err error
	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName), // DSN data source name
		SkipInitializeWithVersion: false,                                                                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
	})
	gtils.Check(err)
	sqlDB, err := global.DB.DB()
	gtils.Check(err)
	sqlDB.SetMaxOpenConns(cfg.ConnMax)
	sqlDB.SetMaxIdleConns(cfg.ConnMax / 2)
}
func (app *App) Run() {
	// 连接数据库
	if app.Config.DB.Enabled {
		app.openDB()
		fmt.Println("Successfully connected to database")
	}

	// 切换生产模式
	//gin.SetMode(gin.ReleaseMode)

	// 启动http服务
	gin.ForceConsoleColor()
	r := router.InitRouter()

	err := r.Run(app.httpServerInstance.Addr)
	if err != nil {
		log.Fatal("GIN RUN ERROR:", err)
	}
}
