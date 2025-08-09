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
	ErrUserDeleted = &DomainError{
		Code:    1011,
		Message: "user deleted",
	}
	ErrInvalidUsername = &DomainError{
		Code:    1012,
		Message: "invalid username",
	}
	ErrInvalidUsernameFormat = &DomainError{
		Code:    1013,
		Message: "invalid username format",
	}
	ErrInvalidEmail = &DomainError{
		Code:    1014,
		Message: "invalid email",
	}
	ErrInvalidEmailFormat = &DomainError{
		Code:    1015,
		Message: "invalid email format",
	}
	ErrInvalidPhone = &DomainError{
		Code:    1016,
		Message: "invalid phone",
	}
	ErrInvalidPhoneFormat = &DomainError{
		Code:    1017,
		Message: "invalid phone format",
	}
	ErrInvalidPasswordFormat = &DomainError{
		Code:    1018,
		Message: "invalid password format",
	}
)
