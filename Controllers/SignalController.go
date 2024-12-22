package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"omeggleClone-signallingServer/Services"
)

func RequestOfferController(w http.ResponseWriter, r *http.Request) {
	connectionID, err := Services.RequestOfferService(r)
	if err != nil {
		if err.Detail != nil {
			log.Printf("Error sending message: %v", err.Detail)
		}
		w.WriteHeader(err.Code)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err.Message)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "Offer requested successfully to %v"}`, connectionID)))
}

func CreateConnectionController(w http.ResponseWriter, r *http.Request) {
	err := Services.CreateConnectionService(r)
	if err != nil {
		if err.Detail != nil {
			log.Printf("Error creating connection: %v", err.Detail)
		}
		w.WriteHeader(err.Code)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err.Message)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "Connection created successfully"}`))
}

func DeactivateConnectionController(w http.ResponseWriter, r *http.Request) {
	err := Services.DeactivateConnectionService(r)
	if err != nil {
		if err.Detail != nil {
			log.Printf("Error deactivating connection: %v", err.Detail)
		}
		w.WriteHeader(err.Code)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err.Message)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "Connection deactivated successfully"}`))
}
