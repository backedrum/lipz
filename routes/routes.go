package routes

import (
	"net/http"

	"github.com/backedrum/lipz/api"
	"github.com/gorilla/mux"
)

// NewRoutes builds the routes for the api
func NewRoutes() *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// devices
	d := a.PathPrefix("/devices").Subrouter()
	d.HandleFunc("/list", api.Devices).Methods("GET")

	// captured packets
	p := a.PathPrefix("/capture/{interfaceName}").Subrouter()
	p.HandleFunc("", api.Packets).Methods("GET")

	return mux
}
