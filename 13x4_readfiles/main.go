package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/google/uuid"
	"strings"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/index", index)
	// mux.HandleFunc("/read", read)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")

		check(err)
		defer mf.Close()

		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)

		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		wd, err := os.Getwd()
		check(err)

		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		check(err)
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)

		c = appendValue(w, c, fname)
	}
	
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)

	// tpl.ExecuteTemplate(w, "index.gohtml", cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("user-data")
	if err != nil {
		uuid := uuid.New()
		cookie = &http.Cookie{
			Name:	"user-data",
			Value:	uuid.String(),
		}
		http.SetCookie(w, cookie)
		fmt.Println("added new cookie")
	}

	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	str := c.Value

	if !strings.Contains(str, fname) {
		str += "|" + fname
	}

	c.Value = str
	http.SetCookie(w, c)
	return c
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}