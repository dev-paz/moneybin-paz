package models

import (
	"fmt"

	dto "github.com/moneybin/moneybin-paz/dto"

	_ "github.com/lib/pq"
)

func ReadUsers() (*[]dto.User, error) {
	var u dto.User
	var users []dto.User
	rows, err := db.Query(`SELECT user_id, user_name, email, password, last_signed_in, signup_timestamp FROM users`)
	if err != nil {
		fmt.Println("error querying")
	}
	for rows.Next() {
		err = rows.Scan(&u.UserID, &u.UserName, &u.Email, &u.Password, &u.LastLoggedIn, &u.SignUpTimestamp)
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
	INSERT INTO users (user_id, user_name, email, password, last_logged_in, signup_timestamp)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING user_id`
	id := ""
	err := db.QueryRow(sqlStatement, u.UserID, u.UserName, u.Email, u.Password, u.LastLoggedIn, u.SignUpTimestamp).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func ReadUser (email string) (*dto.User, error) {
  sqlStatement := `SELECT user_id, user_name, email, password, last_logged_in, signup_timestamp FROM users WHERE email=$1;`
  var u dto.User

  row := db.QueryRow(sqlStatement, email)
  err := row.Scan(&u.UserID, &u.UserName, &u.Email, &u.Password, &u.LastLoggedIn, &u.SignUpTimestamp)
  if err != nil {
    return nil, err
  }
  return &u, nil
}
