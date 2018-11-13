package session

import (
	"net/http"
)

type ContextKey string
type UnsetModeEnums string

// FuncIDGenerator comment
type FuncIDGenerator = func() string

const SESSION_ID_CONTEXT_KEY ContextKey = "session_id"
const DEFAULT_SESSION_NAME string = "gopher.sid"
const UNSET_KEEP UnsetModeEnums = "keep"
const UNSET_DESTROY UnsetModeEnums = "destroy"

// HTTPCookie comment
type HTTPCookie struct {
	http.Cookie
	// someExtraField string
}

// Cookie comment
type Cookie struct {
	HTTPCookie
}

// StoreInterface comment
type StoreInterface interface {
}

// Store comment
type Store struct {
	StoreInterface
}

// Session comment
type Session struct {
	Request *http.Request
	Cookie  *Cookie
}

// MemoryStore comment
type MemoryStore struct {
	*Store
}

// MiddlewareOptions comment
type MiddlewareOptions struct {
	Cookie              *HTTPCookie                  `json:"cookie"`
	GenIDFunction       *FuncIDGenerator             `json:"genid"`
	Name                *string                      `json:"name"`
	Store               *interface{ StoreInterface } `json:"store"`
	IsProxy             *bool                        `json:"proxy"`
	IsResave            *bool                        `json:"resave"`
	IsRolling           *bool                        `json:"rolling"`
	IsSaveUninitialized *bool                        `json:"saveUninitialized"`
	UnsetMode           *UnsetModeEnums              `json:"unset"`
	Secret              *[]string                    `json:"secret"`
}

// Middleware comment
type Middleware struct {
	Cookie              HTTPCookie  `json:"cookie"`
	Name                string      `json:"name"`
	Store               interface{} `json:"store"`
	IsProxy             bool        `json:"proxy"`
	IsResave            bool        `json:"resave"`
	IsRolling           bool        `json:"rolling"`
	IsSaveUninitialized bool        `json:"saveUninitialized"`
	Secret              []string    `json:"secret"`
	Router              http.Handler
	SessionID           string
	unsetDestroy        bool
	generateSessionID   FuncIDGenerator
}

/**
 * Determine if request is secure.
 *
 * @param {Object} req
 * @param {Boolean} [trustProxy]
 * @return {Boolean}
 * @private
 */

func isSecure(r *http.Request, trustProxy bool) bool {
	//   // socket is https server
	//   if (r.connection && r.connection.encrypted) {
	//     return true;
	//   }

	//   // do not trust proxy
	//   if (trustProxy === false) {
	//     return false;
	//   }

	//   // no explicit trust; try req.secure from express
	//   if (trustProxy !== true) {
	//     return req.secure === true
	//   }

	//   // read the proto from x-forwarded-proto header
	//   var header = req.headers['x-forwarded-proto'] || '';
	//   var index = header.indexOf(',');
	//   var proto = index !== -1
	//     ? header.substr(0, index).toLowerCase().trim()
	//     : header.toLowerCase().trim()

	//   return proto === 'https';
}
