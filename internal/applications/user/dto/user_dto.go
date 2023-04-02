package dto

type (
	UserRequest struct {
		RoleId   int32  `json:"role_id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required,password"`
	}

	UserResponse struct {
		Id              string `json:"id" validate:"required"`
		Name            string `json:"name" validate:"required"`
		Email           string `json:"email" validate:"required"`
		IsVerified      string `json:"is_verified" validate:"required"`
		EmailVerifiedAt string `json:"email_verified_at" validate:"required"`
		Password        string `json:"password" validate:"required"`
		RememberToken   string `json:"remember_token" validate:"required"`
		SocialMediaId   string `json:"social_media_id" validate:"required"`
		Avatar          string `json:"avatar" validate:"required"`
		RoleId          string `json:"role_id" validate:"required"`
		LoginType       string `json:"login_type" validate:"required"`
		SubSpecialist   string `json:"sub_specialist" validate:"required"`
		FirebaseToken   string `json:"firebase_token" validate:"required"`
		Info            string `json:"info" validate:"required"`
		Description     string `json:"description" validate:"required"`
		Specialist      string `json:"specialist" validate:"required"`
		Phone           string `json:"phone" validate:"required"`
		LastAccessAt    string `json:"last_access_at" validate:"required"`
		PregnancyMode   string `json:"pregnancy_mode" validate:"required"`

		LatestSkipUpdate string `json:"latest_skip_update" validate:"required"`
		LatestDeletedAt  string `json:"latest_deleted_at" validate:"required"`

		CreatedAt string `json:"created_at" validate:"required"`
		UpdatedAt string `json:"update_at" validate:"required"`
		DeletedAt string `json:"deleted_at" validate:"required"`
	}
)
