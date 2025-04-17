package test

import (
	"go-gate/internal/server/handler/proxy"
	"strings"
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

func TestTrimSuffix(t *testing.T) {
	path := "/api/"
	str := proxy.TrimSuffix(path)
	t.Log(str)

	if strings.HasSuffix(str, "/") {
		t.Errorf("path %s has suffix /", path)
	}

}

