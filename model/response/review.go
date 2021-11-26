package response

import "books-backend/model"

type GetReviewResponse struct {
	Reviews []model.Review `json:"data"`
}
