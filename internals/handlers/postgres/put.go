package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handlers) PutRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	var roomFetched models.Room

	id := chi.URLParam(r, "room-id")
	roomID, _ := strconv.Atoi(id)

	err := data.ReadJSON(w, r, &room)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

	roomFetched, _ = h.DB.FetchRoom(roomID)

	if roomFetched.ID == 0 {

		err = h.DB.InsertRoom(room)
		if err != nil {
			data.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
		res := data.Response{
			Error:   false,
			Message: "room inserted",
		}
		err = data.WriteJSON(w, http.StatusOK, res, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}

	} else {
		if room.ID == 0 {
			room.ID = roomFetched.ID
		}

		if room.RoomName == "" {
			room.RoomName = roomFetched.RoomName
		}
		if room.Price == 0 {
			room.Price = roomFetched.Price
		}
		if room.Rate == 0 {
			room.Rate = roomFetched.Rate
		}
		//err = h.DB.UpdateRoom(room)
		err = h.DB.UpdateRoom(room)
		if err != nil {
			data.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
		res := data.Response{
			Error:   false,
			Message: "room updated",
		}
		err = data.WriteJSON(w, http.StatusOK, res, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}

	}

}
