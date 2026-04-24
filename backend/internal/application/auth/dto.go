package auth

type SignUpRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"required,min=2,max=50"`
}

type SignInRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GoogleSignInRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Language string `json:"language"`
	Theme    string `json:"theme"`
}

type AuthResponse struct {
	Success      bool         `json:"success"`
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken,omitempty"`
	IsNewUser    bool         `json:"isNewUser,omitempty"`
}
