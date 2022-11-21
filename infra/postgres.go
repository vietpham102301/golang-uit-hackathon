package infra

import (
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"go.uber.org/zap"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func InitPostgres(cfg *config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(pg.Open(cfg.DataSource))

	if err != nil {
		zap.S().Debugf("open portgres fail: %+v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Debugf("get DB instance failed %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTimeMiliseconds) * time.Millisecond)

	return db, nil

}
