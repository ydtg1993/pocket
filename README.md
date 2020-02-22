import (
	"./pocket"
	"net/http"
	"time"
)

func main()  {
	p := &pocket.Pocket{
		Connect:32,
		Port:":3000",
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
	}

	Api := pocket.Router{
		Routes: map[string]http.HandlerFunc{
			"/login":          Login,
		},
		Middles: []pocket.MiddleHandler{
			pocket.MethodMiddleHandler{Method: []string{pocket.METHOD_GET}},
		},
	}

	p.Open([]pocket.Router{Api})
}

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("开始注册!!"))
}
