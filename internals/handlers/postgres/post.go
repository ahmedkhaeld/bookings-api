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

func (h *Handlers) PostReservation(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	roomID, err := strconv.Atoi(exploded[2])
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
	}

	// get the room info by querying the database
	fetchedRoom, err := h.DB.FetchRoom(roomID)
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	var payload models.ReservationPayload

	err = data.ReadJSON(w, r, &payload)
	if err != nil {
		data.ErrorJSON(w, err)
		return
	}

	//"2006-01-02T15:04:05Z07:00"
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, payload.StartDate)
	endDate, _ := time.Parse(layout, payload.EndDate)
	id, _ := strconv.Atoi(payload.ID)

	var res models.Reservation
	res.ID = id
	res.RoomID = roomID
	res.FirstName = payload.FistName
	res.LastName = payload.LastName
	res.Email = payload.Email
	res.Phone = payload.Phone
	res.StartDate = startDate
	res.EndDate = endDate
	res.Room = fetchedRoom

	resID, err := h.DB.InsertReservation(res)
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload.ID = strconv.Itoa(resID)
	room := make(map[string]interface{})
	room["room-data"] = fetchedRoom

	if payload.ID == "" {
		output := models.ReservationPayload{
			Ok:      false,
			Message: "reservation failed",
		}
		err = data.WriteJSON(w, http.StatusInternalServerError, output, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}
	} else {
		output := models.ReservationPayload{
			ID:        payload.ID,
			FistName:  payload.FistName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Phone:     payload.Phone,
			RoomID:    payload.RoomID,
			StartDate: payload.StartDate,
			EndDate:   payload.EndDate,
			Room:      room,
			Ok:        true,
			Message:   "reservation successful",
		}
		err = data.WriteJSON(w, http.StatusOK, output, "response")
		if err != nil {
			data.ErrorJSON(w, err)
			return
		}
	}
	rr := models.RoomRestriction{
		ID:            1,
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		ReservationID: resID,
		RestrictionID: 1,
	}

	// insert room restriction of the new created reservation
	err = h.DB.InsertRoomRestriction(rr)
	if err != nil {
		data.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

}
