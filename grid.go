package main

import "net/http"

func (cfg APIConfig) getGrid(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from current grid"))
}
