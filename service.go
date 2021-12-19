package main

import (
	"github.com/google/uuid"
	"time"
)

type IReviewService interface {
	GetReviewById(string) (*ReviewDTO, error)
	GetReviewByRate(int) ([]ReviewDTO, error)
	CreateReview(CreateReviewRequest) (*ReviewDTO, error)
}

func NewReviewService(couchbaseRepository ICouchbaseRepository) IReviewService {
	return &ReviewService{
		couchbaseRepository: couchbaseRepository,
	}
}

type ReviewService struct {
	couchbaseRepository ICouchbaseRepository
}

func (c *ReviewService) GetReviewByRate(rate int) ([]ReviewDTO, error) {

	reviews, err := c.couchbaseRepository.FilterByRate(rate)
	if err != nil {
		return nil, err
	}

	var reviewDTOs []ReviewDTO
	for _, review := range reviews {
		reviewDTO := convertReview(review)
		reviewDTOs = append(reviewDTOs, *reviewDTO)
	}

	return reviewDTOs, nil
}

func (c *ReviewService) GetReviewById(reviewId string) (*ReviewDTO, error) {
	review, err := c.couchbaseRepository.GetById(reviewId)
	if err != nil {
		return nil, err
	}
	reviewDTO := convertReview(*review)
	return reviewDTO, nil
}

func (c *ReviewService) CreateReview(request CreateReviewRequest) (*ReviewDTO, error) {

	r := Review{
		Id:               uuid.New().String(),
		UserId:           request.UserId,
		ContentId:        request.ContentId,
		Rate:             request.Rate,
		Comment:          request.Comment,
		UserName:         request.UserName,
		CreatedDate:      time.Now().UnixMilli(),
		LastModifiedDate: time.Now().UnixMilli(),
	}

	InsertedReview, err := c.couchbaseRepository.Upsert(r)

	if err != nil {
		return nil, err
	}
	reviewDTO := convertReview(*InsertedReview)

	return reviewDTO, nil
}

func convertReview(review Review) *ReviewDTO {
	return &ReviewDTO{
		Id:        review.Id,
		UserId:    review.UserId,
		ContentId: review.ContentId,
		Rate:      review.Rate,
		Comment:   review.Comment,
		UserName:  review.UserName,
	}
}
