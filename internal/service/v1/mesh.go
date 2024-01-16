package v1

import (
	"context"

	instancev1 "github.com/istiofy/istiofy/api/instance/v1"
	v1 "github.com/istiofy/istiofy/api/istiofy/v1"
	"github.com/istiofy/istiofy/internal/dao"
)

type MeshService struct {
	v1.UnimplementedMeshServer
	meshDao        dao.MeshDao
	meshClusterDao dao.MeshClusterDao
}

func NewMeshService(dao dao.Interface) *MeshService {
	return &MeshService{meshDao: dao.MeshDao(), meshClusterDao: dao.MeshClusterDao()}
}

func (m MeshService) Create(ctx context.Context, request *instancev1.CreateMeshRequest) (*instancev1.CreateMeshResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MeshService) Update(ctx context.Context, request *instancev1.UpdateMeshRequest) (*instancev1.UpdateMeshResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MeshService) Delete(ctx context.Context, request *instancev1.DeleteMeshRequest) (*instancev1.DeleteMeshResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MeshService) List(ctx context.Context, request *instancev1.ListMeshRequest) (*instancev1.ListMeshResponse, error) {
	//TODO implement me
	panic("implement me")
}
