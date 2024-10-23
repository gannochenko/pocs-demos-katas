package tag

import (
	"context"

	"api/interfaces"
	"api/internal/domain"
	"api/internal/dto"
	"api/internal/util/db"
	"api/pkg/syserr"
)

type Service struct {
	tagRepository interfaces.TagRepository
}

// New creates a default api service
func New(
	tagRepository interfaces.TagRepository,
) *Service {
	return &Service{
		tagRepository: tagRepository,
	}
}

func (s *Service) ListTags(ctx context.Context, request *domain.ListTagsRequest) (*domain.ListTagsResponse, error) {
	request.Pagination = domain.SanitizePaginationRequest(request.Pagination)

	count, err := s.tagRepository.CountTags(ctx, nil, &dto.ListTagsParameters{
		Filter: &dto.ListTagsFilter{},
	})
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not count tags")
	}

	result := &domain.ListTagsResponse{
		Pagination: domain.NewPaginationResponseFromRequest(request.Pagination, count),
	}

	if count > 0 {
		res, err := s.tagRepository.ListTags(ctx, nil, &dto.ListTagsParameters{
			Filter:     &dto.ListTagsFilter{},
			Pagination: db.NewPaginationFromDomainRequest(request.Pagination),
		})
		if err != nil {
			return nil, syserr.Wrap(err, syserr.InternalCode, "could not list tags")
		}

		for _, tag := range res {
			domainTag, err := tag.ToDomain()
			if err != nil {
				return nil, syserr.Wrap(err, syserr.InternalCode, "could not convert tag to domain")
			}
			result.Pets = append(result.Pets, domainTag)
		}
	}

	return result, nil
}
