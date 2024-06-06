package v1

import (
	"net/http"

	"github.com/guergeiro/fator-conversao-gas-portugal/cmd/v1/controllers"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc(
		"GET /pcs",
		func(w http.ResponseWriter, r *http.Request) {
			output, err := controllers.PcsGet(r)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(output))
		},
	)
	return router
}
