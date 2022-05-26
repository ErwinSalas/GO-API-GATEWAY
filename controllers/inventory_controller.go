package controllers

import (
	"log"
	"net/http"
	"strconv"

	inventorypb "github.com/ErwinSalas/inventory-service/proto"
	"github.com/ErwinSalas/webui/responses"
	"github.com/gorilla/mux"
)

func (server *Server) ListItems(w http.ResponseWriter, r *http.Request) {

	response, err := server.InventoryClient.ListItems(r.Context(), &inventorypb.ListItemsRequest{})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	responses.JSON(w, http.StatusOK, response.Items)
}

func (server *Server) GetItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	response, err := server.InventoryClient.GetItem(r.Context(), &inventorypb.ItemGetRequest{Id: uid})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf(`User Details:
		NAME: %s
		ID: %d`, response.Item.Id, response.Item.Name)

	responses.JSON(w, http.StatusOK, response)
}
