package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/model"
)

type IStoreRepository interface {
	ByID(ctx context.Context, id int64) (result model.Store, err error)
}
