package http_handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/afikrim/go-hexa-template/core/entity"
	"github.com/afikrim/go-hexa-template/core/service"
	errorutil "github.com/afikrim/go-hexa-template/pkg/error"
	httputil "github.com/afikrim/go-hexa-template/pkg/http"
	"github.com/labstack/echo"
)

type TodoHttpHandler struct {
	service service.TodoService
}

func NewTodoHttpHandler(service service.TodoService) *TodoHttpHandler {
	return &TodoHttpHandler{
		service: service,
	}
}

func (h *TodoHttpHandler) Create(e echo.Context) error {
	ctx := context.Background()

	dto := entity.CreateTodoRequest{}
	if err := e.Bind(&dto); err != nil {
		return parseErrCreate(e, fmt.Errorf("%w %v", errorutil.GENERAL_BAD_REQUEST, err.Error()))
	}

	todo, err := h.service.Create(ctx, &dto)
	if err != nil {
		return parseErrCreate(e, err)
	}

	return parseCreateTodoResp(e, todo)
}

func parseErrCreate(e echo.Context, err error) error {
	if errors.Is(err, errorutil.GENERAL_BAD_REQUEST) {
		return e.JSON(
			http.StatusBadRequest,
			httputil.Response{}.NewResponse(httputil.StatusFailed).WithMessage(err.Error()),
		)
	}

	fmt.Printf("err when creating: %v", err)
	return e.JSON(
		http.StatusInternalServerError,
		httputil.Response{}.NewResponse(httputil.StatusFailed).WithInternalError(err.Error()),
	)
}

func (h *TodoHttpHandler) FindAll(e echo.Context) error {
	ctx := context.Background()

	todos, err := h.service.FindAll(ctx)
	if err != nil {
		return parseErrFindAll(e, err)
	}

	return parseFindAllResp(e, todos)
}

func parseErrFindAll(e echo.Context, err error) error {
	return e.JSON(
		http.StatusInternalServerError,
		httputil.Response{}.NewResponse(httputil.StatusFailed).WithInternalError(err.Error()),
	)
}

func (h *TodoHttpHandler) Update(e echo.Context) error {
	ctx := context.Background()

	dto := entity.UpdateTodoRequest{}
	if err := e.Bind(&dto); err != nil {
		return parseErrUpdate(e, fmt.Errorf("%w %v", errorutil.GENERAL_BAD_REQUEST, err.Error()))
	}

	id := e.Param("id")
	if id == "" {
		return parseErrUpdate(e, fmt.Errorf("%w %v", errorutil.GENERAL_BAD_REQUEST, "Id is required"))
	}

	todo, err := h.service.Update(ctx, id, &dto)
	if err != nil {
		return parseErrUpdate(e, err)
	}

	return parseUpdateResp(e, todo)
}

func parseErrUpdate(e echo.Context, err error) error {
	if errors.Is(err, errorutil.GENERAL_BAD_REQUEST) {
		return e.JSON(
			http.StatusBadRequest,
			httputil.Response{}.NewResponse(httputil.StatusFailed).WithMessage(err.Error()),
		)
	}
	if errors.Is(err, errorutil.GENERAL_NOT_FOUND) {
		return e.JSON(
			http.StatusNotFound,
			httputil.Response{}.NewResponse(httputil.StatusFailed).WithMessage(err.Error()),
		)
	}

	return e.JSON(
		http.StatusInternalServerError,
		httputil.Response{}.NewResponse(httputil.StatusFailed).WithInternalError(err.Error()),
	)
}

func (h *TodoHttpHandler) Remove(e echo.Context) error {
	ctx := context.Background()

	id := e.Param("id")
	if id == "" {
		return parseErrRemove(e, fmt.Errorf("%w %v", errorutil.GENERAL_BAD_REQUEST, "Id is required"))
	}

	err := h.service.Remove(ctx, id)
	if err != nil {
		return parseErrRemove(e, err)
	}

	return parseRemoveResp(e)
}

func parseErrRemove(e echo.Context, err error) error {
	if errors.Is(err, errorutil.GENERAL_BAD_REQUEST) {
		return e.JSON(
			http.StatusBadRequest,
			httputil.Response{}.NewResponse(httputil.StatusFailed).WithMessage(err.Error()),
		)
	}
	if errors.Is(err, errorutil.GENERAL_NOT_FOUND) {
		return e.JSON(
			http.StatusNotFound,
			httputil.Response{}.NewResponse(httputil.StatusFailed).WithMessage(err.Error()),
		)
	}

	return e.JSON(
		http.StatusInternalServerError,
		httputil.Response{}.NewResponse(httputil.StatusFailed).WithInternalError(err.Error()),
	)
}
