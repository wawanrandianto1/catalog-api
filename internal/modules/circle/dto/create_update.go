package circle_dto

import (
	"catalog-be/internal/entity"
)

type ImageURLs struct {
	URL          string `json:"url" validate:"omitempty,url,max=255"`
	PictureURL   string `json:"picture_url" validate:"omitempty,url,max=255"`
	FacebookURL  string `json:"facebook_url" validate:"omitempty,url,max=255"`
	InstagramURL string `json:"instagram_url" validate:"omitempty,url,max=255"`
	TwitterURL   string `json:"twitter_url" validate:"omitempty,url,max=255"`
}

type OnboardNewCirclePayload struct {
	Name   string `json:"name" validate:"required,min=3,max=255"`
	Rating string `json:"rating" validate:"required,oneof=GA PG M"`
	ImageURLs
	ReferralCode string `json:"referral_code" validate:"omitempty"`
}

type CreateFandomCircleRelationPayload struct {
	ID   int    `json:"ID" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type CreateWorkTypeCircleRelationPayload struct {
	ID   int    `json:"ID" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type UpdateCirclePayload struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description" validate:"omitempty"`
	Rating      *string `json:"rating" validate:"omitempty,oneof=GA PG M"`

	PictureURL      *string `json:"picture_url" validate:"omitempty,max=255"`
	CoverPictureURL *string `json:"cover_picture_url" validate:"omitempty,max=255"`

	URL          *string `json:"url" validate:"omitempty,url_or_empty,max=255"`
	FacebookURL  *string `json:"facebook_url" validate:"omitempty,url_or_empty,max=255"`
	InstagramURL *string `json:"instagram_url" validate:"omitempty,url_or_empty,max=255"`
	TwitterURL   *string `json:"twitter_url" validate:"omitempty,url_or_empty,max=255"`

	FandomIDs   *[]int `json:"fandom_ids" validate:"omitempty,dive"`
	WorkTypeIDs *[]int `json:"work_type_ids" validate:"omitempty,dive"`
}

type UpdateCircleAttendingEventDayAndBlockPayload struct {
	CircleBlock string      `json:"circle_block" validate:"omitempty"`
	Day         *entity.Day `json:"day" validate:"omitempty,day_or_empty"`
	EventID     int         `json:"event_id" validate:"omitempty"`
}

type BlockResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CircleOneDetailedResponse struct {
	entity.Circle
	Fandom   []entity.Fandom   `json:"fandom"`
	WorkType []entity.WorkType `json:"work_type"`

	Bookmarked bool           `json:"bookmarked"`
	BlockEvent *BlockResponse `json:"block"`
	Event      *entity.Event  `json:"event"`
}

type CirclePaginatedResponse struct {
	entity.Circle
	Description *string           `json:"-"`
	Fandom      []entity.Fandom   `json:"fandom"`
	WorkType    []entity.WorkType `json:"work_type"`

	Bookmarked bool           `json:"bookmarked"`
	BlockEvent *BlockResponse `json:"block"`
	Event      *entity.Event  `json:"event"`
}
