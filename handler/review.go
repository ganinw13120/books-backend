package handler

import (
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"books-backend/model/response"
	"books-backend/service"
)

type reviewHandler struct {
	reviewService service.IReviewService
	redis         *redis.Client
}

type IReviewHandler interface {
	GetAllReview(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
}

func NewReviewHandler(service service.IReviewService, redis *redis.Client) reviewHandler {
	return reviewHandler{
		reviewService: service,
		redis:         redis,
	}
}

func (handler reviewHandler) GetAllReview(c *fiber.Ctx) error {
	name := c.Query("name")
	var response *response.GetReviewListResponse
	var err error
	if name != "" {
		response, err = handler.reviewService.GetReviewByBookName(name)
	} else {
		response, err = handler.reviewService.GetAllReviewList()
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}
	return c.Status(http.StatusOK).JSON(response)
}
func (handler reviewHandler) GetReviewByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to parse request")
	}
	response, err := handler.reviewService.GetReviewByID(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Internal server error")
	}
	return c.Status(http.StatusOK).JSON(response)
}
