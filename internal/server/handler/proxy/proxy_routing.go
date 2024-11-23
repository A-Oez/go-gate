package proxy

import (
	"encoding/json"
	"errors"
	"os"
)

type ClientRequest struct {
	Method string
	Path   string
}

type proxyRoute struct {
	Method            string `json:"method"`
	ClientPath 		  string `json:"client_path"`
	ServiceScheme     string `json:"service_scheme"`
	ServiceHost       string `json:"service_host"`
	ServicePath       string `json:"service_path"`
}

type proxyRoutes struct {
	Routes []proxyRoute `json:"mapping"`
}

func FindRouteMapping(req ClientRequest, path string) (*proxyRoute, error) {
	if path == ""{
		path = "./internal/server/config/proxy_mapping.json"
	}
	var proxyRoutes proxyRoutes
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&proxyRoutes); err != nil {
		return nil, err
	}

	return findRouteByRequest(&proxyRoutes, req)
}

func findRouteByRequest(proxyRoutes *proxyRoutes, req ClientRequest) (*proxyRoute, error) {
	for _, route := range proxyRoutes.Routes {
		if route.Method == req.Method && route.ClientPath == req.Path {
			return &route, nil
		}
	}
	return nil, errors.New("no route mapping found for method " + req.Method + " and path " + req.Path)
}
