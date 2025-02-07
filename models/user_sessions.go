package models

import (
	"fmt"

	dto "github.com/moneybin/moneybin-paz/dto"
)

func CreateUserSession(us *dto.UserSession) error {
	sqlStatement := `
	INSERT INTO user_sessions (user_id, refresh_token)
	VALUES ($1, $2)
	RETURNING user_id`
	id := ""
	fmt.Println(us)
	err := db.QueryRow(sqlStatement, us.UserID, us.RefreshToken).Scan(&id)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func ReadUserSession(user_id string) (*dto.UserSession, error) {
	sqlStatement := `SELECT user_id, refresh_token FROM user_sessions WHERE user_id=$1;`
	var us dto.UserSession
	row := db.QueryRow(sqlStatement, user_id)
	err := row.Scan(&us.UserID, &us.RefreshToken)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &us, nil
}
