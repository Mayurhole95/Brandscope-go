package db

import "errors"

var (
	ErrUserNotExist        = errors.New("Users does not exist in db")
	ErrEmptyData           = errors.New("Books does not exist in db")
	ErrTransactionNotExist = errors.New("Transactions does not exist in db")
	ErrBookNotAvailable    = errors.New("Book is not available")
	ErrBookAlreadyIssued   = errors.New("Book has already been issued")
)
