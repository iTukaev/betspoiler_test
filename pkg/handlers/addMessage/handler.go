package addMessage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Input struct {
	Name string `json:"name"`
	Message string `json:"message"`
}

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
	AddMessage(name string, message string) error
}

func (h *Handle) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request method"))
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	inputPayload := &Input{}
	if err := json.Unmarshal(content, inputPayload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := h.chatService.AddMessage(inputPayload.Name, inputPayload.Message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("message of user %s added", inputPayload.Name)))
}
