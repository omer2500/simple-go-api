package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"../../db/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	// Insert request body into struct
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Fprintf(w, "You didnt provide all the details we need")
	}
	// validate the user
	isUserValidated := vlidateUserInsert(u)
	if len(isUserValidated) > 0 {
		jres, _ := json.Marshal(isUserValidated)
		w.Write(jres)
		return
	}
	// hash user password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// assign more fields to User struct
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.UserId = "oemrsdaJSDAGJ%#(#@"
	u.Password = string(hash)
	res, err := db.Init().Collection("users").InsertOne(context.Background(), u)
	log.Println(err)
	log.Println(res)
}
