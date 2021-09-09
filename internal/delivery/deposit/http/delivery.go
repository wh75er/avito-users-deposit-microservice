package http

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DepositHandler struct {
	DepositUsecase models.DepositUsecase
	Logger *logrus.Logger
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func NewDepositHandler(depositUsecase models.DepositUsecase, e *echo.Echo, logger *logrus.Logger) {
	handler := DepositHandler {
		depositUsecase,
		logger,
	}

	e.GET("/api/v1/deposits/:userUuid", handler.GetDepositInfo)
}

func (h *DepositHandler) GetDepositInfo(c echo.Context) error {
	var targetUuid string
	if err := echo.PathParamsBinder(c).String("userUuid", &targetUuid).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{ err.Error() })
	}

	d, err := h.DepositUsecase.GetDepositByOwner(targetUuid)
	if err != nil {
		return c.JSON(errors.GetHttpError(err), ErrorResponse{ err.Error() })
	}

	return c.JSON(http.StatusOK, DataResponse{ d })
}