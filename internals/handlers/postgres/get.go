package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handlers) GetRooms(w http.ResponseWriter, r *http.Request) {

	//call the database to fetch all rooms
	rooms, err := h.DB.FetchAllRooms()
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}
	//write the data out as json
	err = data.WriteJSON(w, http.StatusOK, rooms, "rooms")
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

}
func (h *Handlers) GetRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "room-id")
	roomID, err := strconv.Atoi(id)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

	room, err := h.DB.FetchRoom(roomID)
	err = data.WriteJSON(w, http.StatusOK, room, "room")
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

}
