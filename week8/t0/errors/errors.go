package errors

var (
	ErrBookNotFound  = "Book not found"
	ErrUserNotFound  = "User not found"
	ErrPasswordSame  = "New password same as old one"
	ErrAlreadyExists = "Already exists"
	ErrUnauthorized  = "Unauthorized"
	ErrBookLent      = "Book is already lent, try another book or just wait"
	ErrBlankInfo     = "Blank info not allowed"
)
