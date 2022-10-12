package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/model"
)

type ICustomerRepository interface {
	ByExternalID(ctx context.Context, externalID string) (result model.Customer, err error)
	Create(ctx context.Context, customer model.Customer) (id int64, err error)
	UpdateByExternalID(ctx context.Context, externalID string, customer model.Customer) (err error)
}
