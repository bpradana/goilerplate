package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bpradana/goilerplate/pkg/domain"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase domain.UserUsecase
}

func NewHandler(e *echo.Group, usecase domain.UserUsecase) {
	handler := &userHandler{
		usecase: usecase,
	}

	// Routes
	e.GET("/users", handler.GetAll)
	e.GET("/users/:id", handler.GetById)
	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.Update)
	e.DELETE("/users/:id", handler.Delete)
}

func (h *userHandler) GetAll(c echo.Context) error {
	users, err := h.usecase.GetAll()
	if err != nil {
		log.Println("[userHandler] [GetAll] error getting all users, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[userHandler] [GetById] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := h.usecase.GetById(id)
	if err != nil {
		log.Println("[userHandler] [GetById] error getting user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) Create(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		log.Println("[userHandler] [Create] error binding user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := h.usecase.Create(user)
	if err != nil {
		log.Println("[userHandler] [Create] error creating user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[userHandler] [GetById] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		log.Println("[userHandler] [Update] error binding user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err = h.usecase.Update(id, user)
	if err != nil {
		log.Println("[userHandler] [Update] error updating user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[userHandler] [GetById] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		log.Println("[userHandler] [Delete] error deleting user, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "User deleted")
}
