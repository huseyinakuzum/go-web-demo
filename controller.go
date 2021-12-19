package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ReviewController struct {
	service IReviewService
}

func NewReviewController(reviewService IReviewService) *ReviewController {
	return &ReviewController{
		service: reviewService,
	}
}

// GetReviewById godoc
// @tags controller
// @Summary get review by id
// @Accept  json
// @Produce  json
// @Param reviewId path string true "ReviewId"
// @Success 200 {object} ReviewDTO
// @Router /reviews/id/{reviewId} [get]
func (controller *ReviewController) GetReviewById(c echo.Context) error {
	reviewId := c.Param("reviewId")
	res, err := controller.service.GetReviewById(reviewId)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

// FilterByRate godoc
// @tags controller
// @Summary get reviews by rate
// @Accept  json
// @Produce  json
// @Param rate path string true "Rate"
// @Success 200 {array} ReviewDTO
// @Router /reviews/rate/{rate} [get]
func (controller *ReviewController) FilterByRate(c echo.Context) error {
	rate, _ := strconv.Atoi(c.Param("rate"))
	res, err := controller.service.GetReviewByRate(rate)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

// CreateReview godoc
// @tags controller
// @Summary create review
// @Accept  json
// @Produce  json
// @Success 200
// @Router /reviews [post]
func (controller *ReviewController) CreateReview(c echo.Context) error {
	var requestBody CreateReviewRequest
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		return err
	} else {
	}
	res, err := controller.service.CreateReview(requestBody)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

func (controller *ReviewController) Register(e *echo.Echo) {
	e.GET("/reviews/id/:reviewId", controller.GetReviewById)
	e.GET("/reviews/rate/:rate", controller.FilterByRate)
	e.POST("/reviews", controller.CreateReview)

}
