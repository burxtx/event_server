package event

import "time"

// Session is aggregate root of session
type Session struct {
	Key        string
	Data       string
	ExpireDate time.Time
	// https://docs.djangoproject.com/en/3.2/topics/http/sessions/#django.contrib.sessions.base_session.AbstractBaseSession.get_decoded
	UserID UserID // thus providing an option to query the database for all active sessions for an account
}

// Repository provides access to a session store.
type SessionRepository interface {
	Create(user *Session) error
	Find(ID UserID) (*Session, error)
	List() []*Session
}
