package image

import (
	"backend/interfaces"
	"backend/internal/util/syserr"
)

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

func (s *Service) SubmitImage() error {
	handle, err := s.sessionManager.Begin()
	if err != nil {
		return syserr.Wrap(err, "could not start transaction")
	}
	defer handle.RollbackUnlessCommitted()

	// start transaction
	// create image
	// create queue
	// commit transaction
	// create message queue event
}
