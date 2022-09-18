package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session-id")
	if err != nil {
		c = &http.Cookie{
			Name:  "session-id",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + `|` + getCode(e)
	}

	http.SetCookie(w, c)
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This `+c.Value+`</a>
	  </body>
	</html>`)
}

func getCode(str string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
