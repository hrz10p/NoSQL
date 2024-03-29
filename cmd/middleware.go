package main

import (
	"context"
	"main/pkg/cookies"
	"main/pkg/models"
	"main/pkg/services"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contextKey string

var contextKeyUser = contextKey("activeUser")

type Middle struct {
	Service *services.Service
}

func NewMiddle(Service *services.Service) *Middle {
	return &Middle{
		Service: Service,
	}
}

func (app *Middle) Authenticate(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.GetCookie(r, "session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		objId, _ := primitive.ObjectIDFromHex(cookie.Value)

		user, err := app.Service.EmployerService.Get(objId)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *Middle) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		emp := getUserFromContext(r)
		if emp.Username == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func (app *Middle) RecoverPanic(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				w.Header().Set("Connection", "close")
// 				logger.GetLogger().Error(err.(error).Error())
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			}
// 		}()

// 		next.ServeHTTP(w, r)
// 	})
// }

// func (app *Middle) LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		logger.GetLogger().Info(fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.RequestURI))
// 		next.ServeHTTP(w, r)
// 	})
// }

// func (app *Middle) SecureHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("X-XSS-Protection", "1; mode=block")
// 		next.ServeHTTP(w, r)
// 	})
// }

func getUserFromContext(r *http.Request) models.Employer {
	user, ok := r.Context().Value(contextKeyUser).(models.Employer)
	if !ok {
		return models.Employer{}
	}
	return user
}
