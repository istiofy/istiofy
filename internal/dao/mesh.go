package dao

import (
	"github.com/istiofy/istiofy/internal/model"
)

type MeshDao interface {
	Create(cluster model.Mesh) (*model.Mesh, error)
	Update(cluster model.Mesh) (*model.Mesh, error)
	Delete(cluster model.Mesh) (*model.Mesh, error)
	List(page, size int32, keyword string) ([]*model.Mesh, error)
	Get(cluster model.Mesh) (*model.Mesh, error)
}
