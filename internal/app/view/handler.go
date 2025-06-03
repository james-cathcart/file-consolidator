package view

import "net/http"

type MainViewHandler struct {
	
}

func NewMainView() http.Handler {
	return &MainViewHandler{}
}

func (h *MainViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
