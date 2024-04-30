package exception

type ErrorBadRequest struct {
	Error string
}

func NewErrorBadRequest(err string) ErrorBadRequest {
	return ErrorBadRequest{
		Error: err,
	}
}
