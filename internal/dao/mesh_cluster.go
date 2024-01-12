package dao

import (
	"github.com/istiofy/istiofy/internal/model"
)

type MeshClusterDao interface {
	Create(cluster model.MeshCluster) (*model.MeshCluster, error)
	Update(cluster model.MeshCluster) (*model.MeshCluster, error)
	Delete(cluster model.MeshCluster) (*model.MeshCluster, error)
	List(page, size int32, keyword string) ([]*model.MeshCluster, error)
	Get(cluster model.MeshCluster) (*model.MeshCluster, error)
}
