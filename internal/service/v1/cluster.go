package v1

import (
	"context"

	instancev1 "github.com/istiofy/istiofy/api/instance/v1"
	v1 "github.com/istiofy/istiofy/api/istiofy/v1"
	"github.com/istiofy/istiofy/internal/dao"
)

type ClusterService struct {
	v1.UnimplementedClusterServer
	clusterDao dao.ClusterDao
}

func (c ClusterService) Create(ctx context.Context, request *instancev1.CreateClusterRequest) (*instancev1.CreateClusterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterService) Update(ctx context.Context, request *instancev1.UpdateClusterRequest) (*instancev1.UpdateClusterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterService) Delete(ctx context.Context, request *instancev1.DeleteClusterRequest) (*instancev1.DeleteClusterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterService) List(ctx context.Context, request *instancev1.ListClusterRequest) (*instancev1.ListClusterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewClusterService(dao dao.Interface) *ClusterService {
	return &ClusterService{clusterDao: dao.ClusterDao()}
}
