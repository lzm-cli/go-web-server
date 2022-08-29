package models

import (
	"context"

	"github.com/<%= organization %>/<%= repo %>/durables"
	"github.com/<%= organization %>/<%= repo %>/session"
)

var Ctx context.Context

func init() {
	db := durables.NewDB()
	Ctx = session.WithDatabase(context.Background(), db)
	db.AutoMigrate(
		&User{},
	)
}
