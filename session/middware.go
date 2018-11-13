package session

import (
	"context"
	"log"
	"net/http"
)

// NewMiddleware create new session middleware
func NewMiddleware(handler http.Handler, opts *MiddlewareOptions) (*Middleware, error) {
	mdw := &Middleware{}

	// get the cookie options
	var cookieOptions HTTPCookie
	if opts.Cookie != nil {
		cookieOptions = *opts.Cookie
	} else {
		cookieOptions = *opts.Cookie
	}
	mdw.Cookie = cookieOptions

	// get the session id generate function
	var generateSessionIDFunction FuncIDGenerator
	if opts.GenIDFunction != nil {
		generateSessionIDFunction = *opts.GenIDFunction
	} else {
		generateSessionIDFunction = generateSessionID
	}
	mdw.generateSessionID = generateSessionIDFunction

	// get the session cookie name
	var sessionName string
	if opts.Name == nil {
		sessionName = DEFAULT_SESSION_NAME
	} else {
		sessionName = *opts.Name
	}
	mdw.Name = sessionName

	// get the session store
	var store StoreInterface
	if opts.Store != nil {
		store = *opts.Store
	} else {
		log.Println(`Warning: connect.session() MemoryStore is not designed for a production environment, as it will leak memory, and will not scale past a single process.`)
		store = &MemoryStore{}
	}
	mdw.Store = store

	// get the trust proxy setting
	var trustProxy bool
	if opts.IsProxy == nil {
		trustProxy = false
	} else {
		trustProxy = *opts.IsProxy
	}
	mdw.IsProxy = trustProxy

	// get the resave session option
	var resaveSession bool
	if opts.IsResave == nil {
		log.Println("undefined resave option; provide resave option")
		resaveSession = true
	} else {
		resaveSession = *opts.IsResave
	}
	mdw.IsResave = resaveSession

	// get the rolling session option
	var rollingSessions bool
	if opts.IsRolling == nil {
		rollingSessions = false
	} else {
		rollingSessions = *opts.IsRolling
	}
	mdw.IsRolling = rollingSessions

	// get the save uninitialized session option
	var saveUninitializedSession bool
	if opts.IsSaveUninitialized == nil {
		log.Println("undefined saveUninitialized option; provide saveUninitialized option")
		saveUninitializedSession = true
	} else {
		saveUninitializedSession = *opts.IsSaveUninitialized
	}
	mdw.IsSaveUninitialized = saveUninitializedSession

	// get the cookie signing secret
	var secret []string
	defaultSecret := []string{"Penn"}

	if opts.Secret == nil {
		log.Println("req.secret; provide secret option")
		secret = defaultSecret
	} else {
		if len(secret) == 0 {
			log.Println("secret option array must contain one or more strings")
			secret = defaultSecret
		} else {
			secret = *opts.Secret
		}
	}
	mdw.Secret = secret

	var unsetDestroy bool
	if opts.UnsetMode != nil {
		unsetDestroy = (*opts.UnsetMode == UNSET_DESTROY)
	}
	mdw.unsetDestroy = unsetDestroy

	return mdw, nil
}

func (ctx *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// // generates the new session
	sessionID := ctx.generateSessionID()
	requestCtx := context.WithValue(
		r.Context(),
		SESSION_ID_CONTEXT_KEY,
		sessionID,
	)

	// if ctx.Cookie.Secure == true {
	// 	ctx.Cookie.Secure = isSecure(r, ctx.IsProxy)
	// }
	// session := &Session{
	// 	Request: r,
	// 	Cookie:  &Cookie{ctx.Cookie},
	// }

	ctx.Router.ServeHTTP(w, r.WithContext(requestCtx))
	// We can modify the response here
}

/**
 * Generate a session ID for a new session.
 *
 * @function generateSessionId
 * @return {String}
 * @private
 */
func generateSessionID() string {
	return ""
}

// func (ctx *SessionMiddleware) () {

// }
