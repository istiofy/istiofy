package model

type Mesh struct {
	Model

	Name string `gorm:"column:name;size:255"`
}
