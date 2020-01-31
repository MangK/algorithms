package http

type Server struct {
	Name string `json:"name"`
	Port string `json:"port"`
	Addr string `json:"addr"`
	Pwd  string `json:"pwd"`
	Dir  string `json:"dir"`
}
