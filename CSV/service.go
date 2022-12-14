package csv

import "context"

type Service interface {
	validate(ctx context.Context) (response BrandHeaderDetails, err error)
}

func (cs *UserService) validate(ctx context.Context) (response ListResponse, err error) {
	users, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	response.Users = users
	return
}
