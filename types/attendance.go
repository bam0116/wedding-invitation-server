package types

import "time"

type AttendanceCreate struct {
	Side  string `json:"side"`
	Name  string `json:"name"`
	Meal  string `json:"meal"`
	Count int    `json:"count"`
}

type Attendance struct {
    ID        int       `json:"id"`
    Side      string    `json:"side"`
    Name      string    `json:"name"`
    Meal      string    `json:"meal"`
    Count     int       `json:"count"`
    CreatedAt time.Time `json:"created_at"`
}