package session

import (
	"net/http"
)

type HTTPCookie struct {
	*http.Cookie
	// someExtraField string
}

type Cookie struct {
}

type MemoryStore struct {
}

type Session struct {
}

type Store struct {
}
