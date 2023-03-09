package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type APIServer struct {
	service Service
}

func NewAPIServer(service Service) *APIServer {
	return &APIServer{
		service: service,
	}
}

func (self *APIServer) start(listenAddr string) error {
	http.HandleFunc("/", self.handleGetCatFact);
	return http.ListenAndServe(listenAddr, nil);
}


func (self APIServer) handleGetCatFact(responseWriter http.ResponseWriter, request *http.Request){
	fact, err := self.service.GetCatFact(context.Background());

	if err != nil {
		WriteJSON(responseWriter, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		});
		return;
	}

	WriteJSON(responseWriter, http.StatusOK, fact)
}

func WriteJSON(responseWriter http.ResponseWriter, status int, value any) error {
	responseWriter.WriteHeader(status);
	responseWriter.Header().Add("Content-Type", "application/json");
	return json.NewEncoder(responseWriter).Encode(value)
}