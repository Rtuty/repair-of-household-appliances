package internal

import (
	"encoding/json"
	"net/http"
)

type storage interface {
	GetApplications() []Application
	AddApplication(app Application)
	UpdateApplicationStatus(id int64, newStatus ApplicationStatus)
	DeleteApplication(id int64)
}

type Handlers struct {
	storage storage
}

func GetHandlers(storage storage) *Handlers {
	return &Handlers{storage: storage}
}

func (h *Handlers) CreateApplicationForm(w http.ResponseWriter, r *http.Request) {
	var app Application

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	h.storage.AddApplication(app)

	w.Write([]byte("Заявка успешно создана"))
}

func (h *Handlers) GetApplications(w http.ResponseWriter, _ *http.Request) {
	applications, err := json.Marshal(h.storage.GetApplications())
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Write(applications)
}
