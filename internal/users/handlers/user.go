package handlers

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"kravets-family-api/internal/users"
	"kravets-family-api/pkg/rest"
	"net/http"
	"strconv"
)

var (
	ErrServiceFindByID = func(err error, ID uint64) error {
		return errors.Wrapf(err, "findByID failed in user service, ID = %d", ID)
	}
)

type UserServiceInterface interface {
	FindByID(id uint64) (users.User, error)
}

type UserHandler struct {
	service UserServiceInterface
}

func NewUserHandler(service UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		rest.BadRequestResponse(w, ErrServiceFindByID(err, uint64(id)))
		return
	}

	user, err := handler.service.FindByID(uint64(id))
	if err != nil {
		rest.ErrResponse(w, ErrServiceFindByID(err, uint64(id)))
		return
	}

	rest.StructResponse(w, user, err)
}
