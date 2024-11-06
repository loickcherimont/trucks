package middlewares

import (
	"net/http"

	"github.com/loickcherimont/trucks/internal/models"
	"github.com/loickcherimont/trucks/internal/utils"
)

// MIDDLEWARES

// Prevent user uses /admin/* without authentication
func CheckLogging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := models.Store.Get(r, "session-name")
		utils.ProcessError(err, w)

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}
