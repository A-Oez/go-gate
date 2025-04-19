package routing

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

type ClientRequest struct {
	Method string
	Path   string
}

type ProxyRoute struct {
	Method        string `json:"method"`
	ClientPath    string `json:"client_path"`
	ServiceScheme string `json:"service_scheme"`
	ServiceHost   string `json:"service_host"`
	ServicePath   string `json:"service_path"`
}

type proxyRoutes struct {
	Routes []ProxyRoute `json:"mapping"`
}

func FindRouteMapping(req ClientRequest, r io.Reader) (*ProxyRoute, error) {	
	var proxyRoutes proxyRoutes
	
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&proxyRoutes); err != nil {
		return nil, err
	}

	return findRouteByRequest(&proxyRoutes, req)
}

func findRouteByRequest(proxyRoutes *proxyRoutes, req ClientRequest) (*ProxyRoute, error) {
	for _, route := range proxyRoutes.Routes {
		if route.Method == req.Method && route.ClientPath == req.Path {
			return &route, nil
		}
	}
	return nil, errors.New("no route mapping found for method " + req.Method + " and path " + req.Path)
}

func TrimSuffix(path string) string {
	if strings.HasSuffix(path, "/") {
		return path[:len(path)-1]
	}
	return path
}