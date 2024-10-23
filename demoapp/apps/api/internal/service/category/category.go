package category

import (
	"context"

	"api/interfaces"
	"api/internal/domain"
	"api/internal/dto"
	"api/internal/util/db"
	"api/pkg/syserr"
)

type Service struct {
	categoryRepository interfaces.CategoryRepository
}

// New creates a default api service
func New(
	categoryRepository interfaces.CategoryRepository,
) *Service {
	return &Service{
		categoryRepository: categoryRepository,
	}
}

func (s *Service) ListCategories(ctx context.Context, request *domain.ListCategoriesRequest) (*domain.ListCategoriesResponse, error) {
	request.Pagination = domain.SanitizePaginationRequest(request.Pagination)

	count, err := s.categoryRepository.CountCategories(ctx, nil, &dto.ListCategoriesParameters{
		Filter: &dto.ListCategoriesFilter{},
	})
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not count categories")
	}

	result := &domain.ListCategoriesResponse{
		Pagination: domain.NewPaginationResponseFromRequest(request.Pagination, count),
	}

	if count > 0 {
		res, err := s.categoryRepository.ListCategories(ctx, nil, &dto.ListCategoriesParameters{
			Filter:     &dto.ListCategoriesFilter{},
			Pagination: db.NewPaginationFromDomainRequest(request.Pagination),
		})
		if err != nil {
			return nil, syserr.Wrap(err, syserr.InternalCode, "could not list categories")
		}

		for _, category := range res {
			domainCategory, err := category.ToDomain()
			if err != nil {
				return nil, syserr.Wrap(err, syserr.InternalCode, "could not convert category to domain")
			}
			result.Categories = append(result.Categories, domainCategory)
		}
	}

	return result, nil
}
