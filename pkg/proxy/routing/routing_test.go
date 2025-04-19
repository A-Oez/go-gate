package routing

import (
	"strings"
	"testing"
)


func TestProxyRouting(t *testing.T) {
	jsonData := `{
		"mapping": [
			{
				"method": "GET",
				"client_path": "/api",
				"service_scheme": "http",
				"service_host": "example.com",
				"service_path": "/real-api"
			}
		]
	}`
	r := strings.NewReader(jsonData)

    proxyRoute, err := FindRouteMapping(ClientRequest{
		Method: "GET",
		Path: "/api",
	}, r)
	if err != nil{
		t.Error(err)
	}
	t.Log(proxyRoute.ServiceHost)
	t.Log(proxyRoute.ServicePath)
}

func TestTrimSuffix(t *testing.T) {
	path := "/api/"
	str := TrimSuffix(path)
	t.Log(str)

	if strings.HasSuffix(str, "/") {
		t.Errorf("path %s has suffix /", path)
	}

}