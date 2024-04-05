package responses

import (
	"go-rest-postgres/domain/posts/models"
	"go-rest-postgres/helpers"

	"github.com/google/uuid"
)

type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedBy string    `json:"created_by"`
}

func (pr *PostResponse) MapResponse(post models.Post) PostResponse {
	postResponse := PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Image:     post.Image,
		CreatedBy: post.UserData.Name,
	}
	return postResponse
}
func (pr *PostResponse) MapResponses(posts []models.Post) []PostResponse {
	var postsResponse []PostResponse
	for _, post := range posts {

		postResponse := PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Image:     post.Image,
			CreatedBy: post.UserData.Name,
		}
		postsResponse = append(postsResponse, postResponse)
	}
	return postsResponse
}
func (pr *PostResponse) MapPaginationResponses(posts []models.Post, pagination helpers.Pagination) (postsResponse []PostResponse, meta helpers.Meta) {
	for _, post := range posts {

		postResponse := PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Image:     post.Image,
			CreatedBy: post.UserData.Name,
		}
		postsResponse = append(postsResponse, postResponse)
	}

	meta.Pagination = pagination

	return postsResponse, meta
}
