package configuration

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	DB *gorm.DB
}

func NewDatabase(params DatabaseParams) (DatabaseResult, error) {

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

	db, err := gorm.Open(postgres.Open(params.Config.DB.Uri), &gorm.Config{})

	if err != nil {
		return DatabaseResult{}, err
	}

	return DatabaseResult{
		DB: db,
	}, nil
}
