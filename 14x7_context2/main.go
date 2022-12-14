package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/bar", bar)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)
		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
