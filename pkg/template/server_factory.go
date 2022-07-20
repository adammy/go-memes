package template

import (
	"github.com/adammy/memepen-services/pkg/httpapi"
)

// NewServer constructs a Server based on the ServerType argument.
func NewServer(t httpapi.ServerType, cfg *Config) (httpapi.Server, error) {
	switch t {
	case httpapi.ChiServerType:
		return NewChiServer(cfg)
	case httpapi.GinServerType:
		return NewGinServer(cfg)
	case httpapi.StdLibServerType:
		return nil, httpapi.ErrServerTypeNotImplemented

	}
	return nil, httpapi.ErrInvalidServerType
}
