package configuration

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config Configuration, log *zap.SugaredLogger) (db *gorm.DB, err error) {
	dsn := config.DB.Uri

	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			config.DB.User,
			config.DB.Pass,
			config.DB.Host,
			config.DB.Port,
			config.DB.Name,
		)
	}

	log.Debugf("Database connection string: %s", dsn)

	db, err = gorm.Open(postgres.Open(config.DB.Uri), &gorm.Config{})

	return
}
