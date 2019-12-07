package models

import (
	"fmt"

	dto "github.com/moneybin/moneybin-paz/dto"

	_ "github.com/lib/pq"
)

func ReadUsers() (*[]dto.User, error) {
	//
	var u dto.User
	var users []dto.User
	rows, err := db.Query(`SELECT user_id, user_name, email , last_signed_in, signup_timestamp FROM users`)
	if err != nil {
		fmt.Println("error querying")
	}
	for rows.Next() {
		err = rows.Scan(&u.UserID, &u.UserName, &u.Email, &u.LastLoggedIn, &u.SignUpTimestamp)
		if err != nil {
			// handle this error
			panic(err)
		}
		users = append(users, u)
	}
	return &users, nil
}

func CreateUser(u *dto.User) error {
	sqlStatement := `
	INSERT INTO users (user_id, user_name, email, last_logged_in, signup_timestamp)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING user_id`
	id := ""
	err := db.QueryRow(sqlStatement, u.UserID, u.UserName, u.Email, u.LastLoggedIn, u.SignUpTimestamp).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func ReadUser(user_id string) (*dto.User, error) {

	sqlStatement := `SELECT user_id, user_name, email, last_logged_in, signup_timestamp FROM users WHERE user_id=$1;`
	var u dto.User
	row := db.QueryRow(sqlStatement, user_id)
	err := row.Scan(&u.UserID, &u.UserName, &u.Email, &u.LastLoggedIn, &u.SignUpTimestamp)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &u, nil
}
