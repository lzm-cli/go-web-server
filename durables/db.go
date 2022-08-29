package durables

import (
	"errors"
	"fmt"

	"github.com/<%= organization %>/<%= repo %>/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.Name,
		config.C.Database.Password)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	)
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&User{})
	return db
}

func CheckEmptyError(err error) error {
	if CheckIsNotFound(err) {
		return nil
	}
	return err
}

func CheckIsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
