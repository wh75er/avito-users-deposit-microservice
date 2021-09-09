package http

import (
	"bank-microservice/internal/errors"
	"bank-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TransactionHandler struct {
	TransactionUsecase models.TransactionUsecase
	Logger *logrus.Logger
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewTransactionHandler(transactionUsecase models.TransactionUsecase, e *echo.Echo, logger *logrus.Logger) {
	handler := TransactionHandler {
		transactionUsecase,
		logger,
	}

	e.POST("/api/v1/deposits/:userUuid/transactions", handler.CreateTransaction)
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	var payload models.Transaction

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{ err.Error() })
	}

	var targetUuid string
	if err := echo.PathParamsBinder(c).String("userUuid", &targetUuid).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{ err.Error() })
	}

	if err := h.TransactionUsecase.CreateTransaction(targetUuid, &payload); err != nil {
		return c.JSON(errors.GetHttpError(err), ErrorResponse{ err.Error() })
	}

	return c.NoContent(http.StatusNoContent)
}