package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handlers) PostRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := data.ReadJSON(w, r, &room)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(room)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

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

}

func (h *Handlers) PostCheckAllRoomsAvailability(w http.ResponseWriter, r *http.Request) {

	var payload models.AvailabilityPayload

	err := data.ReadJSON(w, r, &payload)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, payload.StartDate)
	endDate, _ := time.Parse(layout, payload.EndDate)

	rooms, err := h.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	foundedRooms := make(map[string]interface{})
	foundedRooms["rooms"] = rooms

	roomID := make(map[string]interface{})
	roomID["room"] = payload.RoomID

	if len(rooms) == 0 {
		res := models.AvailabilityPayload{
			Ok:        false,
			Message:   "no rooms for you",
			StartDate: payload.StartDate,
			EndDate:   payload.EndDate,
		}
		err = data.WriteJSON(w, http.StatusOK, res, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}
	} else {
		res := models.AvailabilityPayload{
			Ok:        true,
			Message:   "here is the rooms for you",
			StartDate: payload.StartDate,
			EndDate:   payload.EndDate,
			Rooms:     foundedRooms,
		}
		err = data.WriteJSON(w, http.StatusOK, res, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}
	}

}

func (h *Handlers) PostCheckSingleRoomAvailability(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	roomID, err := strconv.Atoi(exploded[2])
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
	}

	var payload models.AvailabilityPayload

	err = data.ReadJSON(w, r, &payload)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, payload.StartDate)
	endDate, _ := time.Parse(layout, payload.EndDate)

	available, err := h.DB.SearchAvailabilityByDatesByRoomID(startDate, endDate, roomID)
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := models.AvailabilityPayload{
		Ok:        available,
		Message:   "room is available",
		StartDate: payload.StartDate,
		EndDate:   payload.EndDate,
		RoomID:    strconv.Itoa(roomID),
	}
	err = data.WriteJSON(w, http.StatusOK, res, "response")
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

}
