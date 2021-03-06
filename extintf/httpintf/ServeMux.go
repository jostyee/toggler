package httpintf

import (
	"github.com/toggler-io/toggler/extintf/httpintf/swagger"
	"net/http"

	"github.com/toggler-io/toggler/extintf/httpintf/httpapi"
	"github.com/toggler-io/toggler/extintf/httpintf/webgui"
	"github.com/toggler-io/toggler/usecases"
)

func NewServeMux(uc *usecases.UseCases) (*ServeMux, error) {
	mux := http.NewServeMux()

	mux.Handle(`/api/v1/`, letsCORSit(http.StripPrefix(`/api/v1`, httpapi.NewServeMux(uc))))

	ui, err := webgui.NewServeMux(uc)
	if err != nil {
		return nil, err
	}

	mux.Handle(`/`, ui)
	mux.Handle(`/swagger.json`, letsCORSit(http.HandlerFunc(swagger.HandleSwaggerConfigJSON)))
	mux.Handle(`/swagger-ui/`, http.StripPrefix(`/swagger-ui`, swagger.HandleSwaggerUI()))

	return &ServeMux{
		ServeMux: mux,
		UseCases: uc,
	}, nil
}

type ServeMux struct {
	*http.ServeMux
	*usecases.UseCases
}

func letsCORSit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(`Access-Control-Request-Method`, `*`)
		w.Header().Set(`Access-Control-Allow-Headers`, `*`)
		w.Header().Set(`Access-Control-Allow-Origin`, `*`)
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}
