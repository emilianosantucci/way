package configuration

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexlast/bunzap"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DatabaseParams struct {
	fx.In
	Config Configuration
	Log    *zap.SugaredLogger
	Logger *zap.Logger
	Lc     fx.Lifecycle
}

type DatabaseResult struct {
	fx.Out
	DB *bun.DB
}

func NewDatabase(params DatabaseParams) DatabaseResult {

	dsn := params.Config.DB.Uri

	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			params.Config.DB.User,
			params.Config.DB.Pass,
			params.Config.DB.Host,
			params.Config.DB.Port,
			params.Config.DB.Name,
		)
	}

	params.Log.Debugf("Database connection string: %s", dsn)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, pgdialect.New())

	db.AddQueryHook(bunzap.NewQueryHook(bunzap.QueryHookOptions{
		Logger: params.Logger,
		//SlowDuration: 200 * time.Millisecond, // Omit to log all operations as debug
	}))

	params.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			if db != nil {
				return db.Close()
			}
			return nil
		},
	})

	return DatabaseResult{
		DB: db,
	}
}
