package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/transport/utils"
	contextHolder "github.com/supernova0730/ae86/pkg/context_holder"
	"github.com/supernova0730/ae86/pkg/tools"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
	"net/http"
)

func SetContextHolder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.SetUserContext(contextHolder.Init(c.UserContext()))
		return c.Next()
	}
}

func SetMeta() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := uuid.Generate()
		storeID := tools.StrToInt64(c.GetReqHeaders()["X-Store-Id"])
		externalID := c.GetReqHeaders()["X-Customer-Id"]
		languageCode := c.GetReqHeaders()["X-Customer-Language-Code"]
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeRequestID, requestID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeStoreID, storeID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerID, externalID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerLanguageCode, languageCode)
		return c.Next()
	}
}

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				logger.LogWithCtx(c.UserContext()).Error(
					"[PANIC_RECOVERED]",
					zap.Any("panic", r),
				)
				_ = c.SendStatus(http.StatusInternalServerError)
			}
		}()
		return c.Next()
	}
}
