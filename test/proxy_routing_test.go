package test

import (
	"go-gate/internal/server/handler/proxy"
	"testing"
)


func TestProxyRouting(t *testing.T) {
    proxyRoute, err := proxy.FindRouteMapping(proxy.ClientRequest{
		Method: "GET",
		Path: "/api",
	},"../internal/server/config/proxy_mapping.json")
	if err != nil{
		t.Error(err)
	}
	t.Log(proxyRoute.ServiceHost)
	t.Log(proxyRoute.ServicePath)
}
