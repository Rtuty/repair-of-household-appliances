package internal

import (
	"time"
)

type localStorage struct {
	Applications []Application
	Clients      []Client
	Users        []User
}

func NewLocalStorage() *localStorage {
	return &localStorage{
		Applications: make([]Application, 0),
		Clients:      make([]Client, 0),
	}
}

func (ls *localStorage) GetApplications() []Application {
	return ls.Applications
}

func (ls *localStorage) AddApplication(app Application) {
	app.ID = int64(len(ls.Applications) + 1)
	app.Status = NewApplicationStatus
	app.CreatedAt = time.Now()

	ls.Applications = append(ls.Applications, app)
}

func (ls *localStorage) UpdateApplicationStatus(id int64, newStatus ApplicationStatus) {
	for i := 0; i < len(ls.Applications); i++ {
		if ls.Applications[i].ID == id {
			ls.Applications[i].Status = newStatus
			return
		}
	}
}

func (ls *localStorage) DeleteApplication(id int64) {
	for i := 0; i < len(ls.Applications); i++ {
		if ls.Applications[i].ID == id {
			newApplications := make([]Application, len(ls.Applications)-1)

			copy(newApplications, ls.Applications[:i])
			copy(newApplications[i:], ls.Applications[i+1:])

			ls.Applications = newApplications

			return
		}
	}
}

func (ls *localStorage) AddUser(u User) {
	ls.Users = append(ls.Users, u)
}

func (ls *localStorage) GetUser(id int64) User {
	for _, user := range ls.Users {
		if user.ID == id {
			return user
		}
	}

	return User{}
}
