package AppRouter

import (
	"github.com/gorilla/mux"
	"omeggleClone-signallingServer/Controllers"
	"omeggleClone-signallingServer/middleware"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.AwsMiddleware)
	r.HandleFunc("/offer", Controllers.RequestOfferController).Methods("POST")
	r.HandleFunc("/connect", Controllers.CreateConnectionController).Methods("GET")
	r.HandleFunc("/disconnect", Controllers.DeactivateConnectionController).Methods("GET")
	return r
}
