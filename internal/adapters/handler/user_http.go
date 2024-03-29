package handler

import (
	"net/http"

	"example.com/internal/core/domain"
	"example.com/internal/core/services"
	"example.com/utils"
	"github.com/IBM/fp-go/function"
	E "github.com/IBM/fp-go/ioeither"
	"github.com/IBM/fp-go/option"
	"github.com/go-chi/chi"
)

type HTTPUserHandler struct {
	svc services.UserService
}

func NewHTTPUserHandler(message services.UserService) *HTTPUserHandler {
	return &HTTPUserHandler{
		svc: message,
	}
}

func (h *HTTPUserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	function.Pipe2(id, h.svc.GetUser,
		E.Fold[error, *domain.GetUser, int](WriteLeft(w), WriteRight[*domain.GetUser](http.StatusOK, w, option.None[string]())))()
}

func (h *HTTPUserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res := h.svc.GetUsers()
	f := E.Fold[error, []*domain.GetUser, int](WriteLeft(w), WriteRight[[]*domain.GetUser](http.StatusOK, w, option.None[string]()))

	f(res)()
}

func (h *HTTPUserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	function.Pipe3(r.Body,
		utils.ParseJSON[domain.User],
		E.Chain(h.svc.AddUser),
		E.Fold[error, domain.User, int](WriteLeft(w), WriteRight[domain.User](http.StatusCreated, w, option.None[string]())))()
}

func (h *HTTPUserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	function.Pipe3(r.Body,
		utils.ParseJSON[domain.User],
		E.Chain(h.svc.UpdateUser),
		E.Fold[error, domain.User, int](WriteLeft(w), WriteRight[domain.User](http.StatusOK, w, option.None[string]())))()
}

func (h *HTTPUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	function.Pipe2(id, h.svc.DeleteUser,
		E.Fold[error, bool, int](WriteLeft(w), WriteRight[bool](http.StatusOK, w, option.None[string]())))()
}
