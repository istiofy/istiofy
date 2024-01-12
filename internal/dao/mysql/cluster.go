package mysql

import (
	"github.com/istiofy/istiofy/internal/model"

	"gorm.io/gorm"
)

type ClusterDao struct {
	DB *gorm.DB
}

func NewClusterDao(d *gorm.DB) *ClusterDao {
	return &ClusterDao{d}
}

func (d ClusterDao) Create(cluster model.Cluster) (*model.Cluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d ClusterDao) Update(cluster model.Cluster) (*model.Cluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d ClusterDao) Delete(cluster model.Cluster) (*model.Cluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d ClusterDao) List(page, size int32, keyword string) ([]*model.Cluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d ClusterDao) Get(cluster model.Cluster) (*model.Cluster, error) {
	//TODO implement me
	panic("implement me")
}
