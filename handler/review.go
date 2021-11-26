package handler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"books-backend/service"
)

type reviewHandler struct {
	reviewService service.IReviewService
	redis         *redis.Client
}

type IReviewHandler interface {
	GetReview(c *fiber.Ctx) error
}

func NewReviewHandler(service service.IReviewService, redis *redis.Client) reviewHandler {
	return reviewHandler{
		reviewService: service,
		redis:         redis,
	}
}

func (handler reviewHandler) GetReview(c *fiber.Ctx) error {
	response, err := handler.reviewService.GetReviewList()
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}
	return c.Status(http.StatusAccepted).JSON(response)
}
