package custom_error

type BaseError struct {
	Message string
}

func (b BaseError) Error() string {
	return b.Message
}

type InsufficientBalanceError struct {
	BaseError
}

func NewInsufficientBalanceError() InsufficientBalanceError {
	return InsufficientBalanceError{
		BaseError: BaseError{
			Message: "Insufficient Balance",
		},
	}
}

type AccountDoesNotExistsError struct {
	BaseError
}

func NewAccountDoesNotExistsError() AccountDoesNotExistsError {
	return AccountDoesNotExistsError{
		BaseError: BaseError{
			Message: "Account does not exist",
		},
	}
}

type InsufficientCashError struct {
	BaseError
}

func NewInsufficientCashError() InsufficientCashError {
	return InsufficientCashError{
		BaseError: BaseError{
			Message: "Insufficient funds in cash dispenser",
		},
	}
}
