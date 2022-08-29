package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handlers) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "room-id")
	roomID, err := strconv.Atoi(id)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}
	var room models.Room
	room.ID = roomID
	err = h.DB.DeleteRoom(room)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}
	res := data.Response{
		Error:   false,
		Message: "room deleted",
	}
	err = data.WriteJSON(w, http.StatusOK, res, "response")
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

}
