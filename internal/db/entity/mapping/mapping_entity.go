package mapping

type ProxyMapping struct {
	ID            int    `json:"id"`
	Method        string `json:"method"`
	PublicPath    string `json:"public_path"`
	ServiceScheme string `json:"service_scheme"`
	ServiceHost   string `json:"service_host"`
	ServicePath   string `json:"service_path"`
}
