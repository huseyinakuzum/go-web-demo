package main

type CreateReviewRequest struct {
	UserId    int64  `json:"userId,omitempty"`
	ContentId string `json:"contentId,omitempty"`
	Rate      int    `json:"rate,omitempty"`
	Comment   string `json:"comment,omitempty"`
	UserName  string `json:"userName,omitempty"`
}

type Review struct {
	Id               string `json:"id,omitempty"`
	UserId           int64  `json:"userId,omitempty"`
	ContentId        string `json:"contentId,omitempty"`
	ReviewStatus     string `json:"reviewStatus,omitempty"`
	Rate             int    `json:"rate,omitempty"`
	Comment          string `json:"comment,omitempty"`
	UserName         string `json:"userName,omitempty"`
	CreatedDate      int64  `json:"createdDate,omitempty"`
	LastModifiedDate int64  `json:"lastModifiedDate,omitempty"`
}

type ReviewDTO struct {
	Id        string `json:"id,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
	ContentId string `json:"contentId,omitempty"`
	Rate      int    `json:"rate,omitempty"`
	Comment   string `json:"comment,omitempty"`
	UserName  string `json:"userName,omitempty"`
}
