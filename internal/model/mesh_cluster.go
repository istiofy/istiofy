package model

type MeshCluster struct {
	Model

	MeshID               int64   `gorm:"column:mesh_id"`
	ClusterID            int64   `gorm:"column:cluster_id"`
	WithMeshControlPlane bool    `gorm:"column:with_mesh_control_plane"`
	FromMesh             Mesh    `gorm:"foreignKey:MeshID"`
	FromCluster          Cluster `gorm:"foreignKey:ClusterID"`
}
