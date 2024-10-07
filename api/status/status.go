//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml -include-tags=Status ../../tools/openapi.yaml
package status

import "net/http"

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
