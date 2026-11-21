package middleware

type Middleware struct {
	Request *RequestMW
}

func New() *Middleware {
	return &Middleware{
		Request: newRequestMW(),
	}
}
