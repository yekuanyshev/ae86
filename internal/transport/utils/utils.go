package utils

import (
	"context"
	contextHolder "github.com/supernova0730/ae86/pkg/context_holder"
)

type Meta struct {
	RequestID    string `json:"request_id"`
	StoreID      int64  `json:"store_id"`
	CustomerID   string `json:"customer_id"`
	LanguageCode string `json:"language_code"`
}

func GetMeta(ctx context.Context) Meta {
	return Meta{
		StoreID:      contextHolder.GetAttributeInt64(ctx, AttributeStoreID),
		RequestID:    contextHolder.GetAttributeString(ctx, AttributeRequestID),
		CustomerID:   contextHolder.GetAttributeString(ctx, AttributeCustomerID),
		LanguageCode: contextHolder.GetAttributeString(ctx, AttributeCustomerLanguageCode),
	}
}

func GetCustomerID(ctx context.Context) string {
	return contextHolder.GetAttributeString(ctx, AttributeCustomerID)
}
