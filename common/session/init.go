package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionManager is a helper struct that provides methods for managing sessions
type SessionManager struct {
	store   cookie.Store
	options sessions.Options
}

// NewSessionManager creates a new SessionManager with the given store and options
func NewSessionManager(store sessions.Store, options sessions.Options) *SessionManager {
	return &SessionManager{store, options}
}

// GetSession returns the session for the current request
func (sm *SessionManager) GetSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sm.options)
	return session
}

// SetSessionValue sets a value in the session
func (sm *SessionManager) SetSessionValue(c *gin.Context, key string, value interface{}) error {
	session := sm.GetSession(c)
	session.Set(key, value)
	return session.Save()
}

// GetSessionValue retrieves a value from the session
func (sm *SessionManager) GetSessionValue(c *gin.Context, key string) interface{} {
	session := sm.GetSession(c)
	return session.Get(key)
}

// DeleteSessionValue removes a value from the session
func (sm *SessionManager) DeleteSessionValue(c *gin.Context, key string) error {
	session := sm.GetSession(c)
	session.Delete(key)
	return session.Save()
}

// ClearSession removes all values from the session
func (sm *SessionManager) ClearSession(c *gin.Context) error {
	session := sm.GetSession(c)
	session.Clear()
	return session.Save()
}
