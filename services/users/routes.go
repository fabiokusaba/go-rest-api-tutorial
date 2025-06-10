package users

import (
	"fmt"
	"net/http"

	"github.com/fabiokusaba/go-rest-api-tutorial/auth"
	"github.com/fabiokusaba/go-rest-api-tutorial/types"
	"github.com/fabiokusaba/go-rest-api-tutorial/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegister).Methods("POST")
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	// parse JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload"))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %v", errors))
		return
	}

	// check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	// check password
	err = auth.ComparePassword(payload.Password, user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "login successfully",
	})
}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload

	// parse JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload"))
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payload validation error: %v", errors))
		return
	}

	// check if user already exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	// hashed password
	hashedPassword, err := auth.CreateHashedPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error hashing plain password"))
		return
	}

	// register user
	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating a new user: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "user registered successfully",
	})
}
