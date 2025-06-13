package controller

import (
	"database/sql"
	"fmt"
	"rest-explore/internal/config"
	"rest-explore/internal/model"
)

func CreateDataRegistration(data model.RegistrationData) (*model.RegistrationData, error) {

	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(`CREATE TABLE IF NOT EXISTS user (
		name VARCHAR(100),
		email VARCHAR(100) PRIMARY KEY,
		password VARCHAR(100),
		mobile_no VARCHAR(15)
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("Table created", result)

	result, err = db.Exec(`INSERT INTO user (name, email, password, mobile_no) VALUES (?, ?, ?, ?)`,
		data.Name, data.Email, data.Password, data.MobileNo)
	if err != nil {
		return nil, fmt.Errorf("failed to insert data into table: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to insert into table (No row aff): %v", err)
	}
	fmt.Printf("Insert successful, %d row(s) affected\n", rowsAffected)

	return &data, nil
}

func Insert(user model.RegistrationData) (*model.RegistrationData, error) {
	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}

	// Ensure table exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (
		name VARCHAR(100),
		email VARCHAR(100) PRIMARY KEY,
		password VARCHAR(100),
		mobile_no VARCHAR(15)
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	// Insert data
	_, err = db.Exec(`INSERT INTO user (name, email, password, mobile_no) VALUES (?, ?, ?, ?)`,
		user.Name, user.Email, user.Password, user.MobileNo)
	if err != nil {
		return nil, fmt.Errorf("failed to insert: %v", err)
	}

	return &user, nil
}

func GetByName(name string) (*model.RegistrationData, error) {
	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`SELECT name, email, password, mobile_no FROM user WHERE name = ?`, name)

	var data model.RegistrationData
	err = row.Scan(&data.Name, &data.Email, &data.Password, &data.MobileNo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with name: %s", name)
		}
		return nil, err
	}
	return &data, nil
}

// Get all users
func GetAll() []*model.RegistrationData {
	db, err := config.NewDBConnection()
	if err != nil {
		return []*model.RegistrationData{}
	}
	rows, err := db.Query(`SELECT name, email, password, mobile_no FROM user`)
	if err != nil {
		return []*model.RegistrationData{}
	}
	defer rows.Close()

	var allData []*model.RegistrationData
	for rows.Next() {
		var data model.RegistrationData
		err := rows.Scan(&data.Name, &data.Email, &data.Password, &data.MobileNo)
		if err == nil {
			allData = append(allData, &data)
		}
	}
	return allData
}

// Update existing user by email
func Update(data model.RegistrationData) (*model.RegistrationData, error) {
	db, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`UPDATE user SET name = ?, password = ?, mobile_no = ? WHERE email = ?`,
		data.Name, data.Password, data.MobileNo, data.Email)
	if err != nil {
		return nil, fmt.Errorf("update failed: %v", err)
	}
	return &data, nil
}

// Delete user by email
func Delete(email string) error {
	db, err := config.NewDBConnection()
	if err != nil {
		return err
	}

	result, err := db.Exec(`DELETE FROM user WHERE email = ?`, email)
	if err != nil {
		return fmt.Errorf("delete failed: %v", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no user found with email: %s", email)
	}
	return nil
}
