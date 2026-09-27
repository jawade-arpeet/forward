package middleware

type Middleware struct {
	Request *RequestMiddleware
}

func New() *Middleware {
	return &Middleware{
		Request: newRequestMiddleware(),
	}
}
