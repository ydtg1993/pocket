package pocket

import (
	"net/http"
)

type MethodMiddleHandler struct {
	Method []string
}

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

func (M MethodMiddleHandler) Handler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*handel request*/
		for _, method := range M.Method {
			if r.Method == method {
				f.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
