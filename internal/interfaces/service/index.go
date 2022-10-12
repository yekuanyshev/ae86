package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/transport/views"
)

type IIndexService interface {
	Get(ctx context.Context, storeID int64) (view views.Index, err error)
}
