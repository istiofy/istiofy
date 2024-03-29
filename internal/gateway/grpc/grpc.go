package grpc

import (
	"context"
	golog "log"
	"net"
	"os"
	"time"

	v1 "github.com/istiofy/istiofy/api/istiofy/v1"
	"github.com/istiofy/istiofy/internal/dao"
	"github.com/istiofy/istiofy/internal/dao/mysql"
	"github.com/istiofy/istiofy/internal/service"
	servicev1 "github.com/istiofy/istiofy/internal/service/v1"
	"github.com/istiofy/istiofy/pkg/db"
	"github.com/istiofy/istiofy/pkg/log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm/logger"
)

func Run(ctx context.Context, network, address string) error {
	//init grpc server and run
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	go func() {
		defer func() error {
			if err := l.Close(); err != nil {
				return err
			}
			return nil
		}()
		<-ctx.Done()
	}()
	s := grpc.NewServer()

	var daoInterface dao.Interface
	if daoInterface, err = initDao(); err != nil {
		return err
	}
	demoService := service.NewDemoService()
	demoDbService := service.NewDemoDbService(daoInterface)
	clusterService := servicev1.NewClusterService(daoInterface)
	meshService := servicev1.NewMeshService(daoInterface)

	v1.RegisterDemoServer(s, demoService)
	v1.RegisterDemoDbServer(s, demoDbService)
	v1.RegisterClusterServer(s, clusterService)
	v1.RegisterMeshServer(s, meshService)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	go func() error {
		log.L(ctx).Infof("grpc listen on:%s\n", address)
		if err := s.Serve(l); err != nil {
			return err
		}
		return nil
	}()

	return nil
}

func initDao() (dao.Interface, error) {
	newLogger := logger.New(
		golog.New(os.Stdout, "", golog.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(viper.GetInt("data.database.log-level")),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	options := db.Options{
		Host:                  viper.GetString("data.database.host"),
		Port:                  viper.GetString("data.database.port"),
		Username:              viper.GetString("data.database.user"),
		Password:              viper.GetString("data.database.password"),
		Database:              viper.GetString("data.database.database"),
		MaxIdleConnections:    viper.GetInt("data.database.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("data.database.max-open-connections"),
		MaxConnectionLifeTime: time.Duration(viper.GetInt("data.database.max-connection-lifetime")) * time.Second,
		Logger:                newLogger,
	}
	databaseDao, err := mysql.GetDao(&options)
	if err != nil {
		return nil, err
	}
	return databaseDao, nil
}
