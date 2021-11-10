package getAll

import (
	"bytes"
	"net/http"
)

func NewHandler(service chatInterface) func(w http.ResponseWriter, r *http.Request) {
	handler := &Handle {
		chatService: service,
	}
	return handler.Add
}

type Handle struct {
	chatService chatInterface
}

type chatInterface interface {
	GetAllMessage() *bytes.Buffer
}

func (h *Handle) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request method"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(h.chatService.GetAllMessage().Bytes())
}
