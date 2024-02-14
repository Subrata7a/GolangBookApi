package auth

import (
	"GolangBookApi/model"
	"crypto/subtle"
	"fmt"
	"net/http"
)

func BasicAuth(realm string, creds map[string]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w, realm)
				(&model.Error{}).GetError(w, http.StatusNetworkAuthenticationRequired, "StatusNetworkAuthenticationRequired", "Authentication Failed")
				return
			}

			credPass, credUserOk := creds[user]
			if !credUserOk || subtle.ConstantTimeCompare([]byte(pass), []byte(credPass)) != 1 {
				basicAuthFailed(w, realm)
				(&model.Error{}).GetError(w, http.StatusNetworkAuthenticationRequired, "StatusNetworkAuthenticationRequired", "Authentication Failed")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("Authenticated", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
