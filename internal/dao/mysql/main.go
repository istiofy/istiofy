package mysql

import (
	"fmt"
	"sync"

	"github.com/istiofy/istiofy/internal/dao"
	"github.com/istiofy/istiofy/internal/model"
	"github.com/istiofy/istiofy/pkg/db"
	"gorm.io/gorm"
)

type Dao struct {
	Db *gorm.DB
}

func (d *Dao) DemoDbDao() dao.DemoDbDao {
	return NewDemoDbDao(d.Db)
}

func GetDao(opts *db.Options) (dao.Interface, error) {
	var daoInterface dao.Interface
	var once sync.Once

	if opts == nil && daoInterface == nil {
		return nil, fmt.Errorf("failed to get mysql database dao")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Port:                  opts.Port,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			Logger:                opts.Logger,
		}
		dbIns, err = db.NewGORM(options)
		daoInterface = &Dao{dbIns}
	})

	initErr := dbIns.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.DemoDb{},
	)
	if initErr != nil {
		return nil, fmt.Errorf("failed to init mysql istiofy database: %w", err)
	}

	if daoInterface == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql database dao, : %+v, error: %w", daoInterface, err)
	}

	return daoInterface, nil
}
