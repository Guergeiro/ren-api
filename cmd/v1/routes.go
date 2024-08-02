package v1

import (
	"net/http"

	"github.com/guergeiro/fator-conversao-gas-portugal/internal/adapters/controller"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/service"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/application/usecase/pcs"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/infra/connection"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/infra/ren"
	"github.com/guergeiro/fator-conversao-gas-portugal/internal/infra/timescale"
)

func NewRoutes() *http.ServeMux {
	router := http.NewServeMux()
	pcsGetController := controller.NewPcsGetController(
		pcs.NewGetAverageUseCase(
			timescale.NewTimescaleReadingRepository(
				ren.NewRenReadingRepository(
					"https://www.ign.ren.pt/web/guest/monitorizacao-horaria-qualidade",
				),
				connection.PostgresConn,
			),
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
