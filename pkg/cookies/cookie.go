package cookies

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, value string, maxAge time.Time, cookieName string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    value,
		HttpOnly: true,
		Path:     "/",
		Expires:  maxAge,
	}
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func DeleteCookie(w http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
