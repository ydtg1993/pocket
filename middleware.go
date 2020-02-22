package pocket

import "net/http"

type MiddleHandler interface {
	Handler(http.HandlerFunc) http.HandlerFunc
}
