package session

import (
	"log"
	"net/http"
)

// NewMiddleware create new session middleware
func NewMiddleware(handler http.Handler, opts *MiddlewareOptions) (*Middleware, error) {
	mdw := &Middleware{handler: handler}

	// get the cookie options
	var cookieOptions = &HTTPCookie{}
	if opts.Cookie != nil {
		cookieOptions = opts.Cookie
	}
	mdw.Cookie = cookieOptions

	// get the session cookie name
	var sessionName = DEFAULT_SESSION_NAME
	if opts.Name != "" {
		sessionName = opts.Name
	}
	mdw.Name = sessionName

	// get the session store
	var store interface{ StoreInterface }
	if opts.Store != nil {
		store = *opts.Store
	} else {
		log.Println(`Warning: connect.session() MemoryStore is not designed for a production environment, as it will leak memory, and will not scale past a single process.`)
		var memoryStore interface{ StoreInterface } = &MemoryStore{}
		store = memoryStore
	}
	mdw.Store = &store

	// get the trust proxy setting
	mdw.IsTrustProxy = opts.IsTrustProxy

	// get the resave session option
	var resave bool
	if opts.IsResave != nil {
		resave = *opts.IsResave
	} else {
		log.Println("undefined resave option; provide resave option")
		resave = true
	}
	mdw.IsResave = resave

	// get the rolling session option
	mdw.IsRolling = opts.IsRolling

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

	var unsetDestroy = false
	if opts.UnsetMode != nil {
		unsetDestroy = (*opts.UnsetMode == UNSET_DESTROY)
	}
	mdw.unsetDestroy = unsetDestroy

	return mdw, nil
}

func (ctx *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*ctx.Store).SetSession(&Session{
		ID:      generateSessionID(),
		Request: r,
		Cookie: &HTTPCookie{
			Cookie: http.Cookie{
				Name:   "asdasd",
				Value:  generateSessionID(),
				Path:   "/",
				Domain: ".claimh.loc",
				Secure: requestIsSecure(r, ctx.IsTrustProxy),
			},
		},
	})

	ctx.handler.ServeHTTP(w, r)
}
