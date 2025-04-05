package app

import (
	"encoding/json"
	"net/http"

	"bookstore/auth/domain"
)

type AuthController struct {
	authService *domain.AuthService
}

func NewAuthController(authService *domain.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	email, err := domain.NewEmail(req.Email)
	if err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	}

	if err := c.authService.RegisterUser(email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	email, err := domain.NewEmail(req.Email)
	if err != nil {
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	}

	user, err := c.authService.AuthenticateUser(email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
