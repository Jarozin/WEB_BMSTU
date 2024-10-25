package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/app/middleware"
	"project/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

type driverHandler struct {
	driverUsecase     models.DriverUsecaseI
	raceResultUsecase models.RaceResultUsecaseI
}

func NewDriverHandler(m *mux.Router, driverUsecase models.DriverUsecaseI, raceResultUsecase models.RaceResultUsecaseI) {
	handler := &driverHandler{
		driverUsecase:     driverUsecase,
		raceResultUsecase: raceResultUsecase,
	}

	//m.Handle("/api/drivers", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAll), "admin", "user")).Methods("GET")
	m.HandleFunc("/api/drivers", handler.GetAll).Methods("GET")
	m.Handle("/api/drivers", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDriverById), "admin", "user")).Methods("GET")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
}

// @Summary Get all drivers
// @Tags drivers
// @Description Get all drivers
// @ID get-all-drivers
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Driver
// @Failure 500
// @Router /api/drivers [get]
func (handler *driverHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	drivers, err := handler.driverUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Get driver by id
// @Tags drivers
// @Description Get driver by id
// @ID get-driver-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Driver
// @Failure 400
// @Failure 500
// @Router /api/drivers/{id} [get]
func (handler *driverHandler) GetDriverById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	driver, err := handler.driverUsecase.GetDriverById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Create driver
// @Tags drivers
// @Description Create driver
// @ID create-driver
// @Accept  json
// @Produce  json
// @Param input body models.Driver true "driver info"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/drivers [post]
func (handler *driverHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err := decoder.Decode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.driverUsecase.Create(driver)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Update driver
// @Tags drivers
// @Description Update driver
// @ID update-driver
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param input body models.Driver true "driver info"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/drivers/{id} [put]
func (handler *driverHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err = decoder.Decode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.driverUsecase.Update(id, driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Delete driver
// @Tags drivers
// @Description delete driver
// @ID delete-driver
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/drivers/{id} [delete]
func (handler *driverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.driverUsecase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
