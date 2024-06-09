package v1

import (
	"net/http"

	"github.com/guergeiro/fator-conversao-gas-portugal/cmd/v1/controllers"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/service"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/usecase/pcs"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/infra/ren"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	pcsGetController := controllers.NewPcsGetController(
		pcs.NewGetAverageUseCase(
			ren.NewRenReadingRepository(),
			service.NewReadingPruner(),
		),
	)
	router.HandleFunc(
		"GET /pcs",
		func(w http.ResponseWriter, r *http.Request) {
			output, err := pcsGetController.Handle(r)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(output))
		},
	)
	return router
}
