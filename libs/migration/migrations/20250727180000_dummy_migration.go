package migrations

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" [up migration]")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" [down migration]")
		return nil
	})
}
