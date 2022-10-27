package middlewares

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/transport/utils"
	contextHolder "github.com/supernova0730/ae86/pkg/context_holder"
	"github.com/supernova0730/ae86/pkg/tools"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
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

func SetLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		started := time.Now()
		err := c.Next()
		status := c.Response().StatusCode()
		log := logger.LogWithCtx(c.UserContext())
		message := "[FIBER]"
		fields := []zap.Field{
			zap.String("method", c.Method()),
			zap.String("path", c.OriginalURL()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(started)),
			zap.Error(err),
		}

		switch {
		case status >= 400 && status < 500:
			log.Warn(message, fields...)
		case status >= 500:
			log.Error(message, fields...)
		default:
			log.Info(message, fields...)
		}
		return err
	}
}
