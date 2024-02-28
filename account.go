package congo

import (
	"time"
)

type Account struct {
	Id         int       `json:"-"`
	Email      string    `json:"email"`
	Fname      *string   `json:"fname,omitempty"`
	Sname      *string   `json:"sname,omitempty"`
	Phone      *string   `json:"phone"`
	Sex        string    `json:"sex"`
	Birth      time.Time `json:"birth"`
	Country_id *string   `json:"country_id,omitempty"`
	City_id    *string   `json:"city_id,omitempty"`
	Joned      time.Time `json:"joned"`
	Status_id  string    `json:"status_id"`
	Interests  []string  `json:"interests,omitempty"`
	Premium    Premium   `json:"premium,omitempty"`
	Likes      Likes     `json:"likes,omitempty"`
}

type Premium struct {
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
}

type Likes struct {
	Id int       `json:"id"`
	Ts time.Time `json:"ts"`
}
