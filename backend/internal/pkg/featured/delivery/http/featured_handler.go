package http

import (
	"encoding/json"
	"net/http"
	"project/internal/app/middleware"
	"project/internal/pkg/models"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type featuredHandler struct {
	featuredUsecase models.FeaturedUsecaseI
}

func NewFeaturedHandler(m *mux.Router,
	featuredUsecase models.FeaturedUsecaseI) {
	handler := &featuredHandler{
		featuredUsecase: featuredUsecase,
	}

	m.HandleFunc("/api/grandprix/featured", handler.GetAll).Methods("GET")
	m.Handle("/api/grandprix/{id}/featured", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "user", "admin")).Methods("POST")
	m.Handle("/api/grandprix/{id}/featured", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "user", "admin")).Methods("DELETE")
}

// @Summary Get all gp
// @Tags gp
// @Description Get all gp
// @ID get-all-gp
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GrandPrix
// @Failure 500
// @Router /api/grandprix [get]
func (handler *featuredHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaim := r.Context().Value("userClaim")
	user_id := userClaim.(jwt.MapClaims)["user_id"].(int64)

	encoder := json.NewEncoder(w)
	gp, err := handler.featuredUsecase.GetAll(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Create gp
// @Tags gp
// @Description Create gp
// @ID create-gp
// @Accept  json
// @Produce  json
// @Param input body models.GrandPrix true "GP info"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/grandprix [post]
func (handler *featuredHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaim := r.Context().Value("userClaim")
	user_id := userClaim.(jwt.MapClaims)["user_id"].(int64)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feature := models.Featured{
		GpID:   int64(id),
		UserID: user_id,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gp_id, err := handler.featuredUsecase.Create(&feature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(gp_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Delete gp
// @Tags gp
// @Description delete gp
// @ID delete-gp
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/grandprix/{id} [delete]
func (handler *featuredHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userClaim := r.Context().Value("userClaim")
	user_id := userClaim.(jwt.MapClaims)["user_id"].(int64)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feature := models.Featured{
		GpID:   int64(id),
		UserID: user_id,
	}
	err = handler.featuredUsecase.Delete(&feature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
