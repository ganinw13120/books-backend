package router

import (
	"github.com/gofiber/fiber/v2"

	"books-backend/handler"
)

type route struct {
	app *fiber.App

	bookHandler   handler.IBookHandler
	reviewHandler handler.IReviewHandler
}

var router *route = nil

func New(app *fiber.App, bookHandler handler.IBookHandler, reviewHandler handler.IReviewHandler) *route {
	if router == nil {
		router = &route{
			app:           app,
			bookHandler:   bookHandler,
			reviewHandler: reviewHandler,
		}
		router.setUp()
	}
	return router
}
func (r route) setUp() {
	group := r.app.Group("books")
	{
		group.Get("/get", r.bookHandler.GetBooks)
	}
	group = r.app.Group("reviews")
	{
		group.Get("/get", r.reviewHandler.GetReview)
	}
}
