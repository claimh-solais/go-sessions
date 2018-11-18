package session

import (
	"errors"
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
	// ensure a secret is available or bail
	if secret == nil || len(secret) == 0 {
		return nil, errors.New("secret option required for sessions")
	}
	mdw.Secret = secret

	var unsetDestroy = false
	if opts.UnsetMode != nil {
		unsetDestroy = (*opts.UnsetMode == UNSET_DESTROY)
	}
	mdw.unsetDestroy = unsetDestroy
	(*mdw.Store).SetSessionGenerator(GenerateNewSession)

	return mdw, nil
}

var storeReady bool = false

func (ctx *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer ctx.handler.ServeHTTP(w, r)

	// self-awareness
	if sessionID := r.Context().Value(REQUEST_CONTEXT_SESSION_ID); sessionID != "" {
		return
	}

	// Handle connection as if there is no session if
	// the store has temporarily disconnected etc
	if storeReady == false {
		log.Println("store is disconnected")
		return
	}

	// pathname mismatch
	// var originalPath = parseUrl.original(req).pathname || "/"
	// if (originalPath.indexOf(cookieOptions.path || '/') !== 0) {
	// 	return
	// }

	//   var originalHash;
	//   var originalId;
	//   var savedHash;
	//   var touched = false

	//   // expose store
	//   req.sessionStore = store;

	//   // get the session ID from the cookie
	//   var cookieId = req.sessionID = getcookie(req, name, secrets);

	//   // set-cookie
	//   onHeaders(res, function(){
	// 	if (!req.session) {
	// 	  debug('no session');
	// 	  return;
	// 	}

	// 	if (!shouldSetCookie(req)) {
	// 	  return;
	// 	}

	// 	// only send secure cookies via https
	// 	if (req.session.cookie.secure && !issecure(req, trustProxy)) {
	// 	  debug('not secured');
	// 	  return;
	// 	}

	// 	if (!touched) {
	// 	  // touch session
	// 	  req.session.touch()
	// 	  touched = true
	// 	}

	// 	// set cookie
	// 	setcookie(res, name, req.sessionID, secrets[0], req.session.cookie.data);
	//   });

	// 	if (shouldDestroy(req)) {
	// 	  // destroy session
	// 	  debug('destroying');
	// 	  store.destroy(req.sessionID, function ondestroy(err) {
	// 		if (err) {
	// 		  defer(next, err);
	// 		}

	// 		debug('destroyed');
	// 		writeend();
	// 	  });

	// 	  return writetop();
	// 	}

	// 	// no session to save
	// 	if (!req.session) {
	// 	  debug('no session');
	// 	  return _end.call(res, chunk, encoding);
	// 	}

	// 	if (!touched) {
	// 	  // touch session
	// 	  req.session.touch()
	// 	  touched = true
	// 	}

	// 	if (shouldSave(req)) {
	// 	  req.session.save(function onsave(err) {
	// 		if (err) {
	// 		  defer(next, err);
	// 		}

	// 		writeend();
	// 	  });

	// 	  return writetop();
	// 	} else if (storeImplementsTouch && shouldTouch(req)) {
	// 	  // store implements touch method
	// 	  debug('touching');
	// 	  store.touch(req.sessionID, req.session, function ontouch(err) {
	// 		if (err) {
	// 		  defer(next, err);
	// 		}

	// 		debug('touched');
	// 		writeend();
	// 	  });

	// 	  return writetop();
	// 	}
}
