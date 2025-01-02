package image

import "backend/interfaces"

type Service struct {
	sessionManager  interfaces.SessionManager
	imageRepository interfaces.ImageRepository
}

func NewImageService(sessionManager interfaces.SessionManager, imageRepository interfaces.ImageRepository) *Service {
	return &Service{
		sessionManager:  sessionManager,
		imageRepository: imageRepository,
	}
}
