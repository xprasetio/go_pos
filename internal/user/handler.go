package user_domain

import (
	"context"
	"net/http"
	"pos-go/pkg/logger"
	"pos-go/pkg/response"
	pkgvalidator "pos-go/pkg/validator"
	"pos-go/shared/constants"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		logger.Logger.WithError(err).Warn("invalid request body register")
		return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	if req.RoleID == 0 {
		req.RoleID = 2
	}
	if err := pkgvalidator.ValidateRegisterRequest(req); err != nil {
		logger.Logger.WithError(err).Warn("validation error register")
		detail := pkgvalidator.FormatValidationError(err)
		if len(detail) > 0 {
			return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, detail)
		}
		return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	if err := h.Service.Register(context.Background(), req.Name, req.Email, req.Password, req.RoleID); err != nil {
		logger.Logger.WithError(err).Error("failed to register user")
		return response.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	logger.Logger.WithField("email", req.Email).Info("register success")
	return response.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, nil)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		logger.Logger.WithError(err).Warn("invalid request body login")
		return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	if err := pkgvalidator.ValidateLoginRequest(req); err != nil {
		logger.Logger.WithError(err).Warn("validation error login")
		detail := pkgvalidator.FormatValidationError(err)
		if len(detail) > 0 {
			return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, detail)
		}
		return response.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}
	user, token, err := h.Service.Login(&req)
	if err != nil {
		logger.Logger.WithError(err).Warn("invalid credentials login")
		return response.SendResponseHttp(c, http.StatusUnauthorized, constants.ErrNotFound, nil)
	}
	logger.Logger.WithField("email", req.Email).Info("login success")

	resp := map[string]interface{}{
		"user":         user,
		"access_token": token,
	}
	return response.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, resp)
}

// TODO: Tambahkan handler user di sini
