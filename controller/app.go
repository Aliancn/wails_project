package controller

import (
	"context"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	var (
		host, user, password, dbname, port string
	)

	// 读取ini配置文件
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		println("Fail to read file: ", err)
	}
	section, err := cfg.GetSection("pgsql")
	host = section.Key("host").String()
	user = section.Key("user").String()
	password = section.Key("password").String()
	dbname = section.Key("dbname").String()
	port = section.Key("port").String()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	// 初始化数据库postgresql
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	InitDB(db)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
