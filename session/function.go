package session

import "net/http"

// GenerateNewSession comment
func GenerateNewSession(r *http.Request, isTrustProxy bool) (*Session, error) {
	sessionID, err := generateSessionID(12)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:      sessionID,
		Request: r,
		Cookie: &HTTPCookie{
			Cookie: http.Cookie{
				Name:   "asdasd",
				Value:  sessionID,
				Path:   "/",
				Domain: ".claimh.loc",
				Secure: requestIsSecure(r, isTrustProxy),
			},
		},
	}, nil
}
