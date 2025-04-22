package server

import (
	"database/sql"
	routes "go-gate/internal/service/routes/handler"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Server struct {
	Db  	*sql.DB
	Mux 	*http.ServeMux
	Routes  *routes.RoutesHandler
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		Db:  db,
		Mux: http.NewServeMux(),
		Routes: routes.NewRoutesHandler(db),
	}
}

func (s *Server) Start(port string) {
    s.registerRouter()

    str := `
                                       __          
   ____   ____             _________ _/  |_  ____  
  / ___\ /  _ \   ______  / ___\__  \\   __\/ __ \ 
 / /_/  >  <_> ) /_____/ / /_/  > __ \|  | \  ___/ 
 \___  / \____/          \___  (____  /__|  \___  >
/_____/                 /_____/     \/          \/ 
	`

    green := color.New(color.FgGreen).SprintFunc()
    cyan := color.New(color.FgCyan).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

    log.Printf("%s\n%s %s\n%s %s", 
		blue(str), 
		green("Server is running on: 🚀 "), 
		cyan("http://127.0.0.1" + port),
		green("Find the source code here: 🔗 "),
		cyan("https://github.com/A-Oez/go-gate"))

    http.ListenAndServe(port, s.Mux)
}