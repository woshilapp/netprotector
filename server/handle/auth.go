package handle

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/woshilapp/netprotector/server/global"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user global.ApiUser
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hash := fmt.Sprintf("%x", sha1.Sum([]byte(user.Password)))
		auth := false
		token := ""

		for _, u := range global.Users {
			if u.Username == user.Username && u.PasswordHash == hash {
				auth = true
				global.TokenLock.Lock()
				token, err = global.GenerateRandomString(16)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					global.TokenLock.Unlock()
					return
				}
				global.Tokens = append(global.Tokens, token)
				global.TokenLock.Unlock()
				break
			}
		}

		if auth {
			data := struct {
				Token  string
				Status int
			}{
				Token:  token,
				Status: 1,
			}

			encoded, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(encoded)
		} else {
			data := struct {
				Token  string
				Status int
			}{
				Token:  "",
				Status: 0,
			}

			encoded, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(encoded)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var token global.Token
		err := json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		global.TokenLock.Lock()
		global.Tokens = slices.DeleteFunc(global.Tokens, func(s string) bool { return s == token.Token })
	}
}
