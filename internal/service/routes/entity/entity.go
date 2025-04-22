package entity

type Route struct {
	ID            int    `json:"id"`
	PublicPath    string `json:"public_path"`
	Method        string `json:"method"`
	ServiceHost   string `json:"service_host"`
	ServicePath   string `json:"service_path"`
	ServiceScheme string `json:"service_scheme"`
}

type AddRoute struct {
	PublicPath    string `json:"public_path"`
	Method        string `json:"method"`
	ServiceHost   string `json:"service_host"`
	ServicePath   string `json:"service_path"`
	ServiceScheme string `json:"service_scheme"`
}
