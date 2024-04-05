package http

import (
	"context"
	"log"
	"net/http"

	"github.com/stefanprodan/podinfo/pkg/version"
	"go.opentelemetry.io/otel"
)

// Version godoc
// @Summary Version
// @Description returns podinfo version and git commit hash
// @Tags HTTP API
// @Produce json
// @Router /version [get]
// @Success 200 {object} api.MapResponse
func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.GetTracerProvider().Tracer("podinfo").Start(r.Context(), "Version")
	defer span.End()
	childFunc(ctx)
	result := map[string]string{
		"version": version.VERSION,
		"commit":  version.REVISION,
	}
	s.JSONResponse(w, r, result)
}

func childFunc(ctx context.Context) {
	_, span := otel.GetTracerProvider().Tracer("podinfo").Start(ctx, "childFunc")
	defer span.End()
	log.Printf("childFunc")
}
