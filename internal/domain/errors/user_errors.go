package errors

type DomainError struct {
	Code    int
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

var (
	ErrUserNotFound = &DomainError{
		Code:    1001,
		Message: "user not found",
	}
	ErrUsernameTaken = &DomainError{
		Code:    1002,
		Message: "username already taken",
	}
	ErrEmailTaken = &DomainError{
		Code:    1003,
		Message: "email already taken",
	}
	ErrPhoneTaken = &DomainError{
		Code:    1004,
		Message: "phone number already taken",
	}
	ErrInvalidUserInput = &DomainError{
		Code:    1005,
		Message: "invalid user input",
	}
	ErrInvalidPassword = &DomainError{
		Code:    1006,
		Message: "invalid password",
	}
	ErrUserDisabled = &DomainError{
		Code:    1007,
		Message: "user disabled",
	}
	ErrUserNotActive = &DomainError{
		Code:    1008,
		Message: "user not active",
	}
	ErrUserNotExist = &DomainError{
		Code:    1009,
		Message: "user not exist",
	}
	ErrUserAlreadyExist = &DomainError{
		Code:    1010,
		Message: "user already exist",
	}
)
