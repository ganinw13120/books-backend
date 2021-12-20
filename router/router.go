package router

import (
	"github.com/gofiber/fiber/v2"

	"books-backend/handler"
)

type route struct {
	app *fiber.App

	reviewHandler handler.IReviewHandler
}

var router *route = nil

func New(app *fiber.App, reviewHandler handler.IReviewHandler) *route {
	if router == nil {
		router = &route{
			app:           app,
			reviewHandler: reviewHandler,
		}
		router.setUp()
	}
	return router
}
func (r route) setUp() {
	group := r.app.Group("reviews")
	{
		group.Get("/get/all", r.reviewHandler.GetAllReview)
		group.Get("/get/:id", r.reviewHandler.GetReviewByID)
	}
}
