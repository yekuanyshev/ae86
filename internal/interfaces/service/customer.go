package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/model"
)

type ICustomerService interface {
	CheckExistence(ctx context.Context, customer model.Customer) (id int64, err error)
}
