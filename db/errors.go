package db

import "errors"

var (
	ErrUserNotExist        = errors.New("Users does not exist in db")
<<<<<<< HEAD
	ErrBookNotExist        = errors.New("Books does not exist in db")
=======
	ErrEmptyData           = errors.New("Books does not exist in db")
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	ErrTransactionNotExist = errors.New("Transactions does not exist in db")
	ErrBookNotAvailable    = errors.New("Book is not available")
	ErrBookAlreadyIssued   = errors.New("Book has already been issued")
)
