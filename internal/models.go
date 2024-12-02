package internal

import "time"

type Role int32

const (
	UnknownRole Role = iota
	AdminRole
	ClientRole
)

type DeviceInfo struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Model   string `json:"model"`
	Problem string `json:"problem"`
}

type Client struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	ID       int64
	Role     Role
	Password string
}

type ApplicationStatus int32

const (
	UnknownApplicationStatus ApplicationStatus = iota
	ReadyToBeIssuedApplicationStatus
	NewApplicationStatus
	InProgressApplicationStatus
)

type Application struct {
	ID         int64             `json:"id"`
	WorkerID   int64             `json:"worker_id"`
	CreatedAt  time.Time         `json:"created_at"`
	DeviceInfo DeviceInfo        `json:"device_info"`
	ClientInfo Client            `json:"client_info"`
	Status     ApplicationStatus `json:"status"`
}
