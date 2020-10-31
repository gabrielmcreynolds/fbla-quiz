package usecases

import (
	"backend/errorCodes"
	"backend/helpers"
	"backend/users/driver"
	"backend/users/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceImpl struct {
	repo driver.Repository
}

func NewService(repository driver.Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) CreateUser(auth *entity.Authentication) (*entity.User, *errorCodes.Slug) {
	passwordHash, err := helpers.HashPassword(auth.Password)
	if err != nil {
		return nil, errorCodes.NewErrPasswordHashingFailure()
	}
	user, slug := s.repo.AddUser(&entity.User{
		Email:        auth.Email,
		Name:         auth.Name,
		PasswordHash: passwordHash,
		TestsTaken:   0,
		TotalScores:  0,
		TotalTime: entity.Duration{
			Duration: 0,
		},
	})
	if slug != nil {
		return nil, slug
	}
	return user, nil
}

func (s *serviceImpl) Validate(auth *entity.Authentication) (*entity.User, *errorCodes.Slug) {
	user, err := s.repo.FindUserByEmail(auth.Email)
	if err != nil {
		return nil, err
	}

	isValid := helpers.CompareHashToPassword(user.PasswordHash, auth.Password)
	if isValid {
		return user, nil
	} else {
		return nil, errorCodes.NewErrInvalidPassword()
	}
}

func (s *serviceImpl) AddRefreshToken(user *entity.User, ip string) (*entity.RefreshToken, *entity.AccessToken, *errorCodes.Slug) {
	refreshToken := &entity.RefreshToken{
		UserId:      user.ID,
		CreatedByIp: ip,
	}
	_, err := s.repo.AddRefreshToken(refreshToken)

	accessToken := &entity.AccessToken{
		UserID: user.ID,
		Email:  user.Email,
	}

	return refreshToken, accessToken, err
}

func (s *serviceImpl) RefreshToken(jwtString string) (*entity.AccessToken, *errorCodes.Slug) {
	refreshToken, err := entity.NewRefreshTokenFromJWT(jwtString)
	if err != nil {
		return nil, err
	}

	// get the token document and make sure it still exists
	token, err := s.repo.FindRefreshTokenByID(&refreshToken.ID)
	if err != nil {
		return nil, err
	}

	// get the user document and make sure it still exits
	user, err := s.repo.FindUserById(&token.UserId)
	if err != nil {
		return nil, err
	}

	accessToken := &entity.AccessToken{
		UserID: token.UserId,
		Email:  user.Email,
	}
	return accessToken, nil
}

func (s *serviceImpl) RemoveRefreshToken(jwtString string) *errorCodes.Slug {
	refreshToken, err := entity.NewRefreshTokenFromJWT(jwtString)
	if err != nil {
		return err
	}

	err = s.repo.DeleteRefreshToken(&refreshToken.ID)
	return err
}

func (s *serviceImpl) AddTest(userId *primitive.ObjectID, testResult *entity.TestResult) (*entity.User, *errorCodes.Slug) {
	user, err := s.repo.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	user.TotalTime.Duration += testResult.Time.Duration
	user.TotalScores += testResult.Score
	user.TestsTaken++
	err = s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
