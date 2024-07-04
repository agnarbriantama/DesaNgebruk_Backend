package response

import "github.com/agnarbriantama/DesaNgembruk-Backend/models"

type BloggerResponse struct {
	Blogger      models.Blogger `json:"blogger"`
	PresignedURL string         `json:"presignedURL,omitempty"`
	// Sertakan field lain yang diperlukan untuk respons
}
