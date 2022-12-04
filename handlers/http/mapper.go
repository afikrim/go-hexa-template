package http_handler

import (
	"net/http"

	"github.com/afikrim/go-hexa-template/core/entity"
	httputil "github.com/afikrim/go-hexa-template/pkg/http"
	"github.com/labstack/echo"
)

func parseCreateTodoResp(e echo.Context, todo *entity.Todo) error {
	return e.JSON(
		http.StatusCreated,
		httputil.Response{}.NewResponse(httputil.StatusSuccess).WithDataAsMap([]string{"todo"}, []interface{}{todo}),
	)
}

func parseFindAllResp(e echo.Context, todos entity.Todos) error {
	return e.JSON(
		http.StatusOK,
		httputil.Response{}.NewResponse(httputil.StatusSuccess).WithDataAsMap([]string{"todos"}, []interface{}{todos}),
	)
}

func parseUpdateResp(e echo.Context, todo *entity.Todo) error {
	return e.JSON(
		http.StatusOK,
		httputil.Response{}.NewResponse(httputil.StatusSuccess).WithDataAsMap([]string{"todo"}, []interface{}{todo}),
	)
}

func parseRemoveResp(e echo.Context) error {
	return e.JSON(
		http.StatusOK,
		httputil.Response{}.NewResponse(httputil.StatusSuccess),
	)
}
