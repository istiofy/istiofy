package mysql

import (
	"github.com/istiofy/istiofy/internal/model"
	"gorm.io/gorm"
)

type MeshDao struct {
	DB *gorm.DB
}

func NewMeshDao(d *gorm.DB) *MeshDao {
	return &MeshDao{d}
}

func (d MeshDao) Create(cluster model.Mesh) (*model.Mesh, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshDao) Update(cluster model.Mesh) (*model.Mesh, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshDao) Delete(cluster model.Mesh) (*model.Mesh, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshDao) List(page, size int32, keyword string) ([]*model.Mesh, error) {
	//TODO implement me
	panic("implement me")
}

func (d MeshDao) Get(cluster model.Mesh) (*model.Mesh, error) {
	//TODO implement me
	panic("implement me")
}
