package pocket

import "net/http"

type Router struct {
	Routes  map[string]http.HandlerFunc
	Middles []MiddleHandler
}

func Driver(mux *http.ServeMux, routers []Router) {
	for _, router := range routers {
		for path, function := range router.Routes {
			/*middleware*/
			for _, M := range router.Middles {
				function = M.Handler(function)
			}
			mux.Handle(path, function)
		}
	}
}
