package openapi

import (
	"mwstorage/internal/auth"
	"mwstorage/internal/config"
	"net/http"

	openapiStorage "mwstorage/apiAutogenerated/storage"
)

type authTransport struct {
	rt http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	if token, ok := ctx.Value(auth.ContextKeyAuthorization).(string); ok {
		req.Header.Set(auth.HeaderKeyAuthorization, token)
	}
	return t.rt.RoundTrip(req)
}

func MakeStorageAPIClient(cfg *config.Config) *openapiStorage.APIClient {
	chatAPIConfig := &openapiStorage.Configuration{
		Host:   cfg.StorageAPIHost,
		Scheme: "http",
		Servers: openapiStorage.ServerConfigurations{
			{
				URL:         cfg.StorageBaseURL,
				Description: "mw-storage",
			},
		},
		HTTPClient: &http.Client{
			Transport: &authTransport{rt: http.DefaultTransport},
		},
	}
	return openapiStorage.NewAPIClient(chatAPIConfig)
}
