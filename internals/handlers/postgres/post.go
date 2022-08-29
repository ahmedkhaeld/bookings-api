package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/go-playground/validator/v10"
	"net/http"
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
