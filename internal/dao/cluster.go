package dao

import (
	"github.com/istiofy/istiofy/internal/model"
)

type ClusterDao interface {
	Create(cluster model.Cluster) (*model.Cluster, error)
	Update(cluster model.Cluster) (*model.Cluster, error)
	Delete(cluster model.Cluster) (*model.Cluster, error)
	List(page, size int32, keyword string) ([]*model.Cluster, error)
	Get(cluster model.Cluster) (*model.Cluster, error)
}
