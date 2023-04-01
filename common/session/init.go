package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionHandler is a helper struct that provides methods for working with sessions
type SessionHandler struct {
	store sessions.Store
}

// NewSessionHandler creates a new SessionHandler with the given store
func NewSessionHandler(store sessions.Store) *SessionHandler {
	return &SessionHandler{store}
}

// GetSession returns the session for the current request
func (sh *SessionHandler) GetSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sessions.Options{HttpOnly: true, MaxAge: 86400, SameSite: http.SameSiteStrictMode})
	return session
}

// IsAuthenticated returns true if the user is authenticated
func (sh *SessionHandler) IsAuthenticated(c *gin.Context) bool {
	session := sh.GetSession(c)
	return session.Get("authenticated") != nil && session.Get("authenticated").(bool)
}

// SetAuthenticated sets the authenticated status in the session
func (sh *SessionHandler) SetAuthenticated(c *gin.Context) {
	session := sh.GetSession(c)
	session.Set("authenticated", true)
	session.Save()
}

// ClearAuthenticated clears the authenticated status from the session
func (sh *SessionHandler) ClearAuthenticated(c *gin.Context) {
	session := sh.GetSession(c)
	session.Delete("authenticated")
	session.Save()
}
