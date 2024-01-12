package mysql

import (
	"github.com/istiofy/istiofy/internal/model"
	"gorm.io/gorm"
)

type MeshClusterDao struct {
	DB *gorm.DB
}

func NewMeshCluster(d *gorm.DB) *MeshClusterDao {
	return &MeshClusterDao{d}
}

func (d MeshClusterDao) Create(cluster model.MeshCluster) (*model.MeshCluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshClusterDao) Update(cluster model.MeshCluster) (*model.MeshCluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshClusterDao) Delete(cluster model.MeshCluster) (*model.MeshCluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshClusterDao) List(page, size int32, keyword string) ([]*model.MeshCluster, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshClusterDao) Get(cluster model.MeshCluster) (*model.MeshCluster, error) {
	//TODO implement me
	panic("implement me")
}
