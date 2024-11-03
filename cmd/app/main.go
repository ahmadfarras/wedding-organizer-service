package main

import (
	"fartech/wedding-organizer-service/internal/controller"
	zapLogger "fartech/wedding-organizer-service/pkg/log"
	"fartech/wedding-organizer-service/pkg/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initConfig() *model.Config {
	viper.SetConfigName("config")       // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var cfg model.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &cfg
}

func initDB(cfg *model.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.MySQL.User,
		cfg.Database.MySQL.Password,
		cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port,
		cfg.Database.MySQL.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}

func main() {
	// Init config
	config := initConfig()

	// Init database
	db := initDB(config)

	// Init logger
	logger := zapLogger.InitLog(config)

	// Init Fiber
	app := fiber.New()

	// Init route
	controller.InitRoute(app, db, logger)

	// Start server
	port := config.Port
	app.Listen(":" + port)
}
