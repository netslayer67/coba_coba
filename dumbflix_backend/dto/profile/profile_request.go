package profiledto

import "testdumpflix/models"

type CreateProfileRequest struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string                      `json:"phone" gorm:"type: varchar(255)" validate:"required"`
	Gender  string                      `json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Address string                      `json:"address" gorm:"type: text" validate:"required"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
}

type UpdateProfileRequest struct {
	Phone   string                      `json:"phone" gorm:"type: varchar(255)"`
	Gender  string                      `json:"gender" gorm:"type: varchar(255)"`
	Address string                      `json:"address" gorm:"type: text"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
}
