package context_holder

import (
	"context"
	"sync"
)

const ContextHolder = "ContextHolder"

func Init(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextHolder, &sync.Map{})
}

func GetAttributeInt(ctx context.Context, attribute string) int {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if i, ok := value.(int); ok {
			return i
		}
	}
	return 0
}

func GetAttributeInt64(ctx context.Context, attribute string) int64 {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if i, ok := value.(int64); ok {
			return i
		}
	}
	return 0
}

func GetAttributeBool(ctx context.Context, attribute string) bool {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if b, ok := value.(bool); ok {
			return b
		}
	}
	return false
}

func GetAttributeString(ctx context.Context, attribute string) string {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if s, ok := value.(string); ok {
			return s
		}
	}
	return ""
}

func getAttribute(ctx context.Context, attribute string) interface{} {
	if contextHolder, ok := ctx.Value(ContextHolder).(*sync.Map); ok {
		value, ok := contextHolder.Load(attribute)
		if ok {
			return value
		}
	}
	return nil
}

func SetAttribute(ctx context.Context, attribute string, value interface{}) {
	if contextHolder, ok := ctx.Value(ContextHolder).(*sync.Map); ok {
		contextHolder.Store(attribute, value)
	}
}
