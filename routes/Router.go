package Router

import (
	"github.com/gorilla/mux"
	"omeggleClone-signallingServer/AppRouter"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/room", AppRouter.Routes())
	return r
}
