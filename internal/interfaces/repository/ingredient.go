package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/model"
)

type IIngredientRepository interface {
	ListActiveIngredientsByProductID(ctx context.Context, productID int64) (result []model.Ingredient, err error)
}
