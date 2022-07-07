package models

import (
	"context"
	"database/sql"

	"github.com/<%= organization %>/<%= repo %>/session"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func createIgnoreIfExist(ctx context.Context, v interface{}) error {
	return session.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(v).Error
}

func createUpdateAllIfExist(ctx context.Context, v interface{}) error {
	return session.DB(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(v).Error
}

func runInTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return session.DB(ctx).Transaction(fn, &sql.TxOptions{Isolation: sql.LevelSerializable})
}
