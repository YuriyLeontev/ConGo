package congo

import (
	"time"
)

type Account struct {
	Id          int       `json:"-"`
	Email       string    `json:"email"`
	Fname       *string   `json:"fname,omitempty"`
	Sname       *string   `json:"sname,omitempty"`
	Phone       *string   `json:"phone"`
	Sex         string    `json:"sex"`
	Birth       time.Time `json:"birth"`
	Joined      time.Time `json:"joined"`
	Country     *string   `json:"country,omitempty"`
	City        *string   `json:"city,omitempty"`
	Status_user string    `json:"status_user"`
	Interests   []string  `json:"interests,omitempty"`
	Premium     Premium   `json:"premium,omitempty"`
	Likes       Likes     `json:"likes,omitempty"`
}

type Premium struct {
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
}

type Likes struct {
	Id int       `json:"id"`
	Ts time.Time `json:"ts"`
}

type Filter struct {
	Filter   string // Имя параметра фильрации
	Method   string // Метод фильтрации
	Parametr string // Параметр
}
