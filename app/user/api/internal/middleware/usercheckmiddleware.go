package middleware

import "net/http"

type UserCheckMiddleware struct {
}

func NewUserCheckMiddleware() *UserCheckMiddleware {
	return &UserCheckMiddleware{}
}

func (m *UserCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthroughs to next handler if need
		next(w, r)
	}
}
