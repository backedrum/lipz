package routes

import (
	"net/http"

	"github.com/backedrum/lipz/api"
	"github.com/gorilla/mux"
	_ "net/http/pprof"
	"net/http/pprof"
)

// NewRoutes builds the routes for the api
func NewRoutes() *mux.Router {

	m := mux.NewRouter()

	// client static files
	m.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	m.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := m.PathPrefix("/api").Subrouter()

	// devices
	d := a.PathPrefix("/devices").Subrouter()
	d.HandleFunc("/list", api.Devices).Methods("GET")

	// captured packets
	p := a.PathPrefix("/capture/{interfaceName}").Subrouter()
	p.HandleFunc("", api.CapturePackets).Methods("POST")

	m.NewRoute().PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)

	return m
}
