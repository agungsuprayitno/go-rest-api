package responses

import (
	"go-rest-postgres/domain/users/models"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt string    `json:"created_at"`
}

func (ur *UserResponse) MapResponse(user models.User) UserResponse {
	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Photo:     user.Photo,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt.Format("2006-05-10 15:05:05"),
	}
	return userResponse
}
