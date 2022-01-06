package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repo   Repository
	logger log.Logger
}

func FirstService(rep Repository, logger log.Logger) AccountService {
	return &service{
		repo:   rep,
		logger: logger,
	}

}

func (ss service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(ss.logger, "function", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	err := ss.repo.CreateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("userCreated", id)

	return "ok", nil
}

func (ss service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(ss.logger, "function", "CreateUser")

	email, err := ss.repo.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("foundUser", id)

	return email, nil
}
