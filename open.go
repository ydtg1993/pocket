package pocket

import (
	"net/http"
	"time"
)

type Pocket struct {
	Connect      int
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (pocket *Pocket) Open(routers []Router) {
	mux := http.NewServeMux()

	Driver(mux, routers)

	s := &http.Server{
		Addr:           pocket.Port,
		ReadTimeout:    pocket.ReadTimeout,
		WriteTimeout:   pocket.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}
	s.ListenAndServe()
}
