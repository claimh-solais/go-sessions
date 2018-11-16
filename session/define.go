package session

import (
	"flag"
	"net/http"
)

type UnsetMode string

const (
	UNSET_KEEP    UnsetMode = "keep"
	UNSET_DESTROY UnsetMode = "destroy"
)

const DEFAULT_SESSION_NAME string = "gopher.sid"

var DISABLE_RESAVE = flag.Bool("DISABLE_RESAVE", false, "")
var ENABLE_RESAVE = flag.Bool("ENABLE_RESAVE", true, "")
var DISABLE_SAVE_UNINITIALIZED = flag.Bool("DISABLE_SAVE_UNINITIALIZED", false, "")
var ENABLE_SAVE_UNINITIALIZED = flag.Bool("ENABLE_SAVE_UNINITIALIZED", true, "")

type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
)

// type HTTPCookieOptions struct {
// 	Name  string
// 	Value string

// 	Path       string    // optional
// 	Domain     string    // optional
// 	Expires    time.Time // optional
// 	RawExpires string    // for reading cookies only

// 	// MaxAge=0 means no 'Max-Age' attribute specified.
// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// 	// MaxAge>0 means Max-Age attribute present and given in seconds
// 	MaxAge   int
// 	Secure   bool
// 	HTTPOnly bool
// 	SameSite SameSite
// 	Raw      string
// 	Unparsed []string // Raw text of unparsed attribute-value pairs
// }

type HTTPCookie struct {
	http.Cookie
	SameSite SameSite
}

// StoreInterface comment
type StoreInterface interface {
	SetSession(*Session)
}

// Store comment
type Store struct {
	StoreInterface
	Session *Session
}

// Session comment
type Session struct {
	ID      string
	Request *http.Request
	Cookie  *HTTPCookie
}

// MemoryStore comment
type MemoryStore struct {
	StoreInterface
	Store
}

// MiddlewareOptions comment
type MiddlewareOptions struct {
	Cookie              *HTTPCookie                  `json:"cookie"`
	Name                string                       `json:"name"`
	Store               *interface{ StoreInterface } `json:"store"`
	IsTrustProxy        bool                         `json:"proxy"`
	IsResave            *bool                        `json:"resave"`
	IsRolling           bool                         `json:"rolling"`
	IsSaveUninitialized *bool                        `json:"saveUninitialized"`
	UnsetMode           *UnsetMode                   `json:"unset"`
	Secret              *[]string                    `json:"secret"`
}

// Middleware comment
type Middleware struct {
	SessionID    string
	Store        *interface{ StoreInterface }
	handler      http.Handler
	unsetDestroy bool
	MiddlewareOptions
	IsResave            bool
	IsSaveUninitialized bool
	Secret              []string
}
