package custom_error

type BaseError struct {
	Message string
}

func (b BaseError) Error() string {
	return b.Message
}

type AccountNotRegisteredError struct {
	BaseError
}

type InsufficentBalanceError struct {
	BaseError
}

type InsufficientStockError struct {
	BaseError
}

func NewAccountNotRegisteredError() AccountNotRegisteredError {
	return AccountNotRegisteredError{
		BaseError: BaseError{
			Message: "Account Not Registered",
		},
	}
}

func NewInsufficentBalanceError() InsufficentBalanceError {
	return InsufficentBalanceError{
		BaseError: BaseError{
			Message: "Insufficient Balance",
		},
	}
}

func NewInsufficientStockError() InsufficientStockError {
	return InsufficientStockError{
		BaseError: BaseError{
			Message: "Insufficient Stock",
		},
	}
}
