package response

import "books-backend/model"

type GetReviewListResponse struct {
	Reviews []model.Review `json:"data"`
}

type GetReviewResponse struct {
	Review model.Review `json:"data"`
}
