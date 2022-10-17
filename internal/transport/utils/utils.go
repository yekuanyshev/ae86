package utils

import (
	"context"
	contextHolder "github.com/supernova0730/ae86/pkg/context_holder"
)

type Meta struct {
	RequestID    string `json:"request_id"`
	StoreID      int64  `json:"store_id"`
	ExternalID   string `json:"external_id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

func GetMeta(ctx context.Context) Meta {
	return Meta{
		StoreID:      contextHolder.GetAttributeInt64(ctx, AttributeStoreID),
		RequestID:    contextHolder.GetAttributeString(ctx, AttributeRequestID),
		ExternalID:   contextHolder.GetAttributeString(ctx, AttributeCustomerExternalID),
		Username:     contextHolder.GetAttributeString(ctx, AttributeCustomerUsername),
		FirstName:    contextHolder.GetAttributeString(ctx, AttributeCustomerFirstName),
		LastName:     contextHolder.GetAttributeString(ctx, AttributeCustomerLastName),
		LanguageCode: contextHolder.GetAttributeString(ctx, AttributeCustomerLanguageCode),
	}
}

func GetCustomerID(ctx context.Context) int64 {
	return contextHolder.GetAttributeInt64(ctx, AttributeCustomerID)
}
