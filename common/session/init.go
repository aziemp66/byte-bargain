package session

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SessionManager is a helper struct that provides methods for managing sessions
type SessionManager struct {
	store  cookie.Store
	option sessions.Options
}

// NewSessionManager creates a new SessionManager with the given store and options
func NewSessionManager(secretKey []byte) *SessionManager {
	store := cookie.NewStore(secretKey)
	option := sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	store.Options(option)

	return &SessionManager{
		store:  store,
		option: option,
	}
}

// UseSession adds the session middleware to the gin engine
func (sm *SessionManager) GetSessionHandler() gin.HandlerFunc {
	return sessions.Sessions(fmt.Sprintf("freya_session_%s", uuid.NewString()), sm.store)
}

// GetSession returns the session for the current request
func (sm *SessionManager) GetSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sm.option)
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
