package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"github.com/supernova0730/ae86/internal/transport/rest/response"
	"github.com/supernova0730/ae86/internal/transport/utils"
	contextHolder "github.com/supernova0730/ae86/pkg/context_holder"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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
		storeID, err := strconv.ParseInt(c.GetReqHeaders()["X-Store-Id"], 10, 64)
		if err != nil {
			return response.Error(c, http.StatusBadRequest, err)
		}
		externalID := c.GetReqHeaders()["X-Customer-External-Id"]
		username := c.GetReqHeaders()["X-Customer-Username"]
		firstName := c.GetReqHeaders()["X-Customer-First-Name"]
		lastName := c.GetReqHeaders()["X-Customer-Last-Name"]
		languageCode := c.GetReqHeaders()["X-Customer-Language-Code"]
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeRequestID, requestID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeStoreID, storeID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerExternalID, externalID)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerUsername, username)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerFirstName, firstName)
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerLastName, lastName)
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

func CheckCustomer(mc container.IMainContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		meta := utils.GetMeta(c.UserContext())
		id, err := mc.Services().Customer().CheckExistence(c.UserContext(), model.Customer{
			ExternalID: meta.ExternalID,
			Username:   meta.Username,
			FirstName:  meta.FirstName,
			LastName:   meta.LastName,
		})
		if err != nil {
			return response.Error(c, http.StatusInternalServerError, err)
		}
		contextHolder.SetAttribute(c.UserContext(), utils.AttributeCustomerID, id)
		return c.Next()
	}
}
