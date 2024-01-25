package main

import (
	"main/pkg/services"
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

// func (app *Middle) Authenticate(next http.Handler) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := cookies.GetCookie(r, "session")
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), contextKeyUser, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func (app *Middle) RequireAuthentication(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		emp := getEmployerFromContext(r)
// 		app := getApplicantFromContext(r)
// 		if emp == nil && app == nil {
// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

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

// func getEmployerFromContext(r *http.Request) *models.Employer {
// 	user, ok := r.Context().Value(contextKeyUser).(models.Employer)
// 	if !ok {
// 		return nil
// 	}
// 	return &user
// }

// func getApplicantFromContext(r *http.Request) *models.Applicant {
// 	user, ok := r.Context().Value(contextKeyUser).(models.Applicant)
// 	if !ok {
// 		return nil
// 	}
// 	return &user
// }
