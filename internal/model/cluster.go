package model

type Cluster struct {
	Model

	Name       string `gorm:"column:name;size:255"`
	KubeConfig string `gorm:"column:kube_config;size:10000"`
}
