package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/transport/utils"
	"strconv"
)

func SetContextHolder() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.SetUserContext(utils.SetContextHolder(c.UserContext()))
		return c.Next()
	}
}

func SetStoreID() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		contextHolder := utils.GetContextHolder(c.UserContext())
		if contextHolder != nil {
			storeIDHeader := c.GetReqHeaders()["X-Store-Id"]
			storeID, _ := strconv.ParseInt(storeIDHeader, 10, 64)
			contextHolder.Store(utils.AttributeStoreID, storeID)
		}
		return c.Next()
	}
}
