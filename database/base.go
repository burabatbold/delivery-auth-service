package database

import (
	"fmt"

	"github.com/burabatbold/delivery-auth-service/database/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Connect() {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=%s",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.dbname"),
		viper.GetString("db.password"),
		viper.GetString("db.timezone"),
	)), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&entity.AdminEntity{}, &entity.CustomerEntity{}, &entity.DriverEntity{}, &entity.MerchantEntity{}); err != nil {
		panic("failed to migrate database")
	}

	DB = db
}
