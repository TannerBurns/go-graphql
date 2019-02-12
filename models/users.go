package models

import "fmt"

// User shape
type UserInfo struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

// GetUsersByName is called within our user query for graphql
func (d *Database) GetUsersByName(name string) []UserInfo {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	// Create User struct for holding each row's data
	var r UserInfo
	// Create slice of Users for our response
	users := []UserInfo{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}
