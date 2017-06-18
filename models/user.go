package models

import (
	"fmt"
	"time"
)

type User struct {
	ID         int
	Name       string
	UID        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsSpeaking bool
	IsActive   bool
}

func FindUser(uid string) (*User, error) {
	selectPart := "id, name, uid, created_at, updated_at, is_speaking, is_active"
	sqlQuery := fmt.Sprintf("SELECT %s FROM users WHERE uid = '%s'", selectPart, uid)
	rows, err := DB.Query(sqlQuery)

	if err != nil {
		return nil, err
	}

	user := User{}

	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.UID, &user.CreatedAt, &user.UpdatedAt, &user.IsSpeaking, &user.IsActive)
	} else {
		return nil, nil
	}

	return &user, nil
}
