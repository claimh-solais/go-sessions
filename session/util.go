package session

import (
	"crypto/rand"
	"net/http"
)

/**
 * Generate a session ID for a new session.
 *
 * @function generateSessionId
 * @return {String}
 * @private
 */
func generateSessionID(size int) (string, error) {
	size = 10
	buffer := make([]byte, size)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:]), nil
}

/**
 * Determine if request is secure.
 *
 * @param {Object} req
 * @param {Boolean} [trustProxy]
 * @return {Boolean}
 * @private
 */
func requestIsSecure(r *http.Request, trustProxy bool) bool {
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
	return false
}
