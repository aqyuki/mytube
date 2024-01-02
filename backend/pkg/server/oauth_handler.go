package server

import (
	"log/slog"
	"net/http"

	"github.com/aqyuki/mytube/backend/pkg/logging"
	mid "github.com/aqyuki/mytube/backend/pkg/middleware"
	"github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/labstack/echo/v4"
)

func (s *Server) setOAuthRouting(e *echo.Echo) {
	oauthGroup := e.Group("/oauth")

	oauthGroup.POST("/register", s.registerHandler)
	oauthGroup.POST("/login", s.loginHandler)
	oauthGroup.POST("/logout", s.logoutHandler)
	oauthGroup.DELETE("/delete", s.deleteAccountHandler)
}

type OAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OAuthSuccessResponse struct {
	OAuth   bool `json:"oauth"`
	Deleted bool `json:"deleted"`
}

type OAuthErrorResponse struct {
	Error string `json:"error"`
}

func (s *Server) registerHandler(c echo.Context) error {
	logger := mid.UnwrapLogger(c)

	var req OAuthRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind request body", slog.Any("error", err))
		return c.JSON(http.StatusBadRequest, OAuthErrorResponse{
			Error: "Request body is invalid",
		})
	}

	ctx := logging.WithLogger(c.Request().Context(), mid.UnwrapLogger(c))
	user, err := s.accountService.Register(ctx, req.Username, req.Password)
	if err != nil {
		logger.Error("Failed to register account", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to register account",
		})
	}

	// save to session
	content := session.Content{
		Username:  user.Username,
		OAuth:     true,
		CreatedAt: user.CreatedAt.Unix(),
	}
	if err := s.sessionManager.SaveContent(c, &content); err != nil {
		logger.Error("Failed to save session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to save session",
		})
	}

	logger.Info("Successfully register account", slog.Any("username", user.Username))
	return c.JSON(http.StatusOK, OAuthSuccessResponse{
		OAuth:   true,
		Deleted: false,
	})
}

func (s *Server) loginHandler(c echo.Context) error {
	logger := mid.UnwrapLogger(c)
	data, err := s.sessionManager.GetContent(c)
	if err != nil {
		logger.Error("Failed to get session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to get session",
		})
	} else if data.OAuth {
		logger.Info("Already logged in", slog.Any("username", data.Username))
		return c.JSON(http.StatusOK, OAuthSuccessResponse{
			OAuth:   true,
			Deleted: false,
		})
	}

	var req OAuthRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind request body", slog.Any("error", err))
		return c.JSON(http.StatusBadRequest, OAuthErrorResponse{
			Error: "Request body is invalid",
		})
	}

	ctx := logging.WithLogger(c.Request().Context(), mid.UnwrapLogger(c))
	user, err := s.accountService.Login(ctx, req.Username, req.Password)
	if err != nil {
		logger.Error("Failed to login system", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to login system",
		})
	}

	// save to session
	content := session.Content{
		Username:  user.Username,
		OAuth:     true,
		CreatedAt: user.CreatedAt.Unix(),
	}
	if err := s.sessionManager.SaveContent(c, &content); err != nil {
		logger.Error("Failed to save session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to save session",
		})
	}

	logger.Info("Successfully login system", slog.Any("username", user.Username))
	return c.JSON(http.StatusOK, OAuthSuccessResponse{
		OAuth:   true,
		Deleted: false,
	})
}

func (s *Server) logoutHandler(c echo.Context) error {
	logger := mid.UnwrapLogger(c)
	data, err := s.sessionManager.GetContent(c)
	if err != nil {
		logger.Error("Failed to get session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to get session",
		})
	} else if !data.OAuth {
		logger.Info("Already logged out")
		return c.JSON(http.StatusOK, OAuthSuccessResponse{
			OAuth:   false,
			Deleted: false,
		})
	}

	if err := s.sessionManager.DeleteContent(c); err != nil {
		logger.Error("Failed to delete session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to delete session",
		})
	}

	logger.Info("Successfully logout system", slog.Any("username", data.Username))
	return c.JSON(http.StatusOK, OAuthSuccessResponse{
		OAuth:   false,
		Deleted: false,
	})
}

func (s *Server) deleteAccountHandler(c echo.Context) error {
	logger := mid.UnwrapLogger(c)
	data, err := s.sessionManager.GetContent(c)
	if err != nil {
		logger.Error("Failed to get session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to get session",
		})
	} else if !data.OAuth {
		logger.Info("Already logged out")
		return c.JSON(http.StatusOK, OAuthSuccessResponse{
			OAuth:   false,
			Deleted: false,
		})
	}

	ctx := logging.WithLogger(c.Request().Context(), mid.UnwrapLogger(c))
	if err := s.accountService.Delete(ctx, data.Username); err != nil {
		logger.Error("Failed to delete account", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to delete account",
		})
	}

	if err := s.sessionManager.DeleteContent(c); err != nil {
		logger.Error("Failed to delete session", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, OAuthErrorResponse{
			Error: "Internal server error : Failed to delete session",
		})
	}
	return c.JSON(http.StatusOK, OAuthSuccessResponse{
		OAuth:   false,
		Deleted: true,
	})
}
