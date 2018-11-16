package session

import "log"

// SetSession comment
func (ctx *Store) SetSession(sess *Session) {
	log.Println(sess.Cookie.String())
}
