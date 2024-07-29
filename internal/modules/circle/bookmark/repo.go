package bookmark

import (
	"catalog-be/internal/domain"
	"catalog-be/internal/entity"

	"gorm.io/gorm"
)

type CircleBookmarkRepo struct {
	db *gorm.DB
}

// FindByCircleIDAndUserID implements CircleBookmarkRepo.
func (c *CircleBookmarkRepo) FindByCircleIDAndUserID(circleID int, userID int) (*entity.UserBookmark, *domain.Error) {
	bookmark := new(entity.UserBookmark)
	err := c.db.Table("user_bookmark").Where("circle_id = ? AND user_id = ?", circleID, userID).First(bookmark).Error

	if err != nil {
		return nil, domain.NewError(500, err, nil)
	}

	return bookmark, nil
}

// DeleteBookmark implements CircleBookmarkRepo.
func (c *CircleBookmarkRepo) DeleteBookmark(circleID int, userID int) *domain.Error {
	err := c.db.Table("user_bookmark").Where("circle_id = ? AND user_id = ?", circleID, userID).Delete(&entity.UserBookmark{}).Error

	if err != nil {
		return domain.NewError(500, err, nil)
	}

	return nil
}

// CreateBookmark implements CircleBookmarkRepo.
func (c *CircleBookmarkRepo) CreateBookmark(circleID int, userID int) *domain.Error {
	err := c.db.Table("user_bookmark").Create(&entity.UserBookmark{
		UserID:   userID,
		CircleID: circleID,
	}).Error

	if err != nil {
		return domain.NewError(500, err, nil)
	}

	return nil
}

func NewCircleBookmarkRepo(db *gorm.DB) *CircleBookmarkRepo {
	return &CircleBookmarkRepo{db: db}
}
