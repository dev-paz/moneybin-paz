package models

import (
	"fmt"

	dto "github.com/moneybin/moneybin-paz/dto"

	_ "github.com/lib/pq"
)

func ReadDonations() (*[]dto.Donation, error) {
	var d dto.Donation
	var donations []dto.Donation
	rows, err := db.Query(`SELECT donationid, username, userid, amount FROM donations`)
	if err != nil {
		fmt.Println("error querying")
	}
	for rows.Next() {
		rows.Scan(&d.ID, &d.UserName, &d.UserId, &d.Amount, &d.DonationCreatedTimestamp)
		donations = append(donations, d)
	}
	return &donations, nil
}

func CreateDonation(d *dto.Donation) error {
	fmt.Println(d)
	sqlStatement := `
	INSERT INTO donations (donation_id, user_name, user_id, amount, donation_created_timestamp)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING donation_id`
	id := 0
	err := db.QueryRow(sqlStatement, d.ID, d.UserName, d.UserId, d.Amount, d.DonationCreatedTimestamp).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
