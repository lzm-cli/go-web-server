package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/<%= organization %>/<%= repo %>/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.Name,
		config.C.Database.Password)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}

func createIgnoreIfExist(v interface{}) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(v).Error
}

func createUpdateAllIfExist(v interface{}) error {
	return db.Clauses(clause.OnConflict{UpdateAll: true}).Create(v).Error
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

func runInTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
	if tx.Error != nil {
		return tx.Error
	}
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit().Error
}
