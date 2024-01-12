package dao

import (
	"context"

	"github.com/istiofy/istiofy/internal/model"
)

type DemoDbDao interface {
	DemoDb(ctx context.Context, demoDb string) (*model.DemoDb, error)
}
