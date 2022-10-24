package main

import "time"

type DailyReports struct {
    Id    int `json:id`
	Date  string `json:date`
	Student_id  *int `json:student_id`
	Attend int `json:attend`
	Temperature *string `json:temperature`
	SomeoneToPickUp *string `json:someoneToPickUp`
	TimeToPickUp *string `json:timeToPickUp`
	Message *string `json:message`
    Center_id *int `json:center_id`
    Student_name *string `json:student_name`
}

type TeacherMessage struct {
    Id int `json:id`
    Staff_id int `json:staff_id`
    Message *string `json:message`
    Datetime string `json:datetime`
    Student_id int `json:student_id`
}

type Students struct {
    Id int `json:id`
    Center_id int `json:center_id`
    Name string`json:name`
    ContactTell *string `json:contactTell`
    Grade int `json:grade`
    Email *string `json:email`
    Status int `json:status`
    Rfid *string `json:rfid`
    CenterName string `json:centerName`
}

type StuInAndOutSensors struct {
    Id int  `json:id`
    Datetime string  `json:datetime`
    Rfid string  `json:rfid`
    Sensor_id int  `json:sensor_id`
    Name string  `json:name`
    Place string  `json:place`
}

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
