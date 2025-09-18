package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shivamhw/learn-by-practice/golang/databases/sqllite/models"
)

var (
	db      *sql.DB
	userDb  *models.Users
	eventDb *models.Events
)

func main() {
	if len(os.Args) < 2 {
		panic("no cmd specified")
	}
	var err error
	cmd := os.Args[1]
	db, err = sql.Open("sqlite3", "data.db?_foreign_keys=on")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userDb = models.NewUsersDb(db)
	eventDb = models.NewEventDb(db)

	switch cmd {
	case "migrate":
		db_migrate(db, os.Args[2])
	case "user":
		sub_cmd := os.Args[2]
		switch sub_cmd {
		case "add":
			CreateUser()
		case "list":
			ListUser()
		case "get" :
			id, _:= strconv.Atoi(os.Args[3])
			CheckUser(int64(id))
		}
	case "event":
		sub_cmd := os.Args[2]
		switch sub_cmd {
		case "add":
			id, _:= strconv.Atoi(os.Args[3])
			CreateEvent(int64(id))
		case "list":
			ListEvent()
		case "get" :
			id, _:= strconv.Atoi(os.Args[3])
			CheckEvent(int64(id))
		}
	}
}

func CheckEvent(id int64) {
	event, err := eventDb.Get(id)
	if err != nil {
		fmt.Print("err: ", err)
		return
	}
	fmt.Printf("%d %s %s %s %d", event.Id, event.Name, event.Description, event.Date, event.Owner_id)
}

func CreateEvent(owner_id int64) {
	name := uuid.NewString()
	event := models.EventModel{
		Name:        fmt.Sprintf("evetn_%s", name[:4]),
		Description: "test description",
		Owner_id:    owner_id,
		Date:        time.Now(),
	}
	e, err := eventDb.Insert(event)
	if err != nil {
		fmt.Print("err: ", err)
		return
	}
	fmt.Print("New event created ", e)
}

func ListEvent() {
	events, err := eventDb.List()
	if err != nil {
		fmt.Print("err:", err)
		return
	}
	fmt.Println("nof events ", len(events))
	for _, e := range events {
		fmt.Printf("%d %s %s %s %d\n", e.Id, e.Name, e.Description, e.Date, e.Owner_id)
	}
}

func ListUser() {
	users, err := userDb.List()
	if err != nil {
		fmt.Printf("err : %v", err)
		return
	}
	fmt.Println("total users : ", len(users))
	for _, v := range users {
		fmt.Printf("%d %s %s %s\n", v.Id, v.Name, v.Email, v.Password)
	}
}

func CheckUser(userid int64) {
	user, err := userDb.Get(userid)
	if err != nil {
		fmt.Printf("failed checking user with %q", err)
	}
	fmt.Printf("User found with %v", user)
}

func CreateUser() {

	user := models.UserModel{
		Name:     "shivamwh",
		Email:    "main@mail.com" + uuid.NewString()[:4],
		Password: "test234",
	}
	cUser, err := userDb.Insert(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created User %v", cUser)
}
