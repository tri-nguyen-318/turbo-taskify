package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"

	"turbo-taskify/backend/internal/domain/user"
	"turbo-taskify/backend/internal/infrastructure/token"
)

var (
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrEmailExists        = errors.New("email_already_exists")
	ErrUsernameTaken      = errors.New("username_taken")
	ErrInvalidToken       = errors.New("invalid_token")
	ErrTokenExpired       = errors.New("token_expired")
)

type Service struct {
	userRepo       user.Repository
	tokenSvc       *token.Service
	googleClientID string
}

func NewService(repo user.Repository, ts *token.Service, googleClientID string) *Service {
	return &Service{
		userRepo:       repo,
		tokenSvc:       ts,
		googleClientID: googleClientID,
	}
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) (*AuthResponse, error) {
	if exists, err := s.userRepo.ExistsEmail(ctx, req.Email); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrEmailExists
	}

	if exists, err := s.userRepo.ExistsUsername(ctx, req.Username); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrUsernameTaken
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := user.New(req.Email, req.Username, string(hashed))
	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, err
	}

	return s.buildResponse(u, false)
}

func (s *Service) SignIn(ctx context.Context, req SignInRequest) (*AuthResponse, error) {
	u, err := s.userRepo.FindByEmail(ctx, req.Email)
	if errors.Is(err, user.ErrNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.buildResponse(u, false)
}

func (s *Service) GoogleSignIn(ctx context.Context, req GoogleSignInRequest) (*AuthResponse, bool, error) {
	payload, err := idtoken.Validate(ctx, req.IDToken, s.googleClientID)
	if err != nil {
		return nil, false, ErrInvalidToken
	}

	googleID := payload.Subject
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	if name == "" {
		name = email
	}

	// Existing user by Google ID
	if u, err := s.userRepo.FindByGoogleID(ctx, googleID); err == nil {
		resp, err := s.buildResponse(u, false)
		return resp, false, err
	}

	// Existing user by email — link Google
	if u, err := s.userRepo.FindByEmail(ctx, email); err == nil {
		resp, err := s.buildResponse(u, false)
		return resp, false, err
	}

	// New user
	username := sanitizeUsername(name)
	if exists, _ := s.userRepo.ExistsUsername(ctx, username); exists {
		username = username + "_" + googleID[:6]
	}

	u := user.NewFromGoogle(email, username, googleID)
	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, false, err
	}

	resp, err := s.buildResponse(u, true)
	return resp, true, err
}

func (s *Service) GetMe(ctx context.Context, userID string) (*UserResponse, error) {
	u, err := s.userRepo.FindByID(ctx, userID)
	if errors.Is(err, user.ErrNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}
	resp := toUserResponse(u)
	return &resp, nil
}

func (s *Service) RefreshToken(_ context.Context, refreshToken string) (string, error) {
	claims, err := s.tokenSvc.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", ErrTokenExpired
	}
	return s.tokenSvc.GenerateAccessToken(claims.UserID)
}

func (s *Service) buildResponse(u *user.User, isNew bool) (*AuthResponse, error) {
	id := u.ID.Hex()
	accessToken, err := s.tokenSvc.GenerateAccessToken(id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.tokenSvc.GenerateRefreshToken(id)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		Success:      true,
		User:         toUserResponse(u),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsNewUser:    isNew,
	}, nil
}

func toUserResponse(u *user.User) UserResponse {
	return UserResponse{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		Username: u.Username,
		Language: u.Language,
		Theme:    u.Theme,
	}
}

func sanitizeUsername(name string) string {
	out := make([]byte, 0, len(name))
	for i := range len(name) {
		c := name[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' {
			out = append(out, c)
		} else {
			out = append(out, '_')
		}
	}
	if len(out) < 2 {
		out = append(out, '_', '_')
	}
	return string(out)
}
