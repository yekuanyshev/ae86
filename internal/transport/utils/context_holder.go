package utils

import (
	"context"
	"sync"
)

func SetContextHolder(ctx context.Context) context.Context {
	return context.WithValue(ctx, AttributeContextHolder, &sync.Map{})
}

func GetContextHolder(ctx context.Context) *sync.Map {
	holder, ok := ctx.Value(AttributeContextHolder).(*sync.Map)
	if !ok {
		return nil
	}
	return holder
}

func GetStoreID(ctx context.Context) int64 {
	contextHolder := GetContextHolder(ctx)
	if contextHolder == nil {
		return 0
	}

	storeIDValue, ok := contextHolder.Load(AttributeStoreID)
	if !ok {
		return 0
	}

	storeID, ok := storeIDValue.(int64)
	if !ok {
		return 0
	}

	return storeID
}
