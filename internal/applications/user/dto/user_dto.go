package dto

import "time"

type (
	UserRequest struct {
		RoleId   int32  `json:"role_id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	UserResponse struct {
		Id       int64  `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Avatar   string `json:"avatar" validate:"required"`
		Password string `json:"password" validate:"required"`
		RoleId   int32  `json:"role_id" validate:"required"`

		IsVerified      bool      `json:"is_verified" validate:"required"`
		EmailVerifiedAt time.Time `json:"email_verified_at" validate:"required"`
		RememberToken   string    `json:"remember_token" validate:"omitempty"`
		SocialMediaId   string    `json:"social_media_id" validate:"omitempty"`
		LoginType       string    `json:"login_type" validate:"omitempty"`
		SubSpecialist   string    `json:"sub_specialist" validate:"omitempty"`
		FirebaseToken   string    `json:"firebase_token" validate:"omitempty"`
		Info            string    `json:"info" validate:"omitempty"`
		Description     string    `json:"description" validate:"omitempty"`
		Specialist      string    `json:"specialist" validate:"omitempty"`
		Phone           string    `json:"phone" validate:"omitempty"`
		LastAccessAt    time.Time `json:"last_access_at" validate:"required"`
		PregnancyMode   bool      `json:"pregnancy_mode" validate:"required"`

		LatestSkipUpdate time.Time `json:"latest_skip_update" validate:"omitempty"`
		LatestDeletedAt  time.Time `json:"latest_deleted_at" validate:"omitempty"`

		CreatedAt time.Time `json:"created_at" validate:"required"`
		UpdatedAt time.Time `json:"update_at" validate:"required"`
		DeletedAt time.Time `json:"deleted_at" validate:"required"`
	}
)
