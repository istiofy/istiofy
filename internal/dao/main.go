package dao

// Interface is the interface for istiofy.
type Interface interface {
	DemoDbDao() DemoDbDao
	ClusterDao() ClusterDao
	MeshDao() MeshDao
	MeshClusterDao() MeshClusterDao
}
