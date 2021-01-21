package model

import (
	"encoding/json"
	"errors"
	"go_poj/pkg/errno"
	"log"
)

type User struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (user *User) SelectUserByName(name string) error {
	stmt, err := DB.Prepare("SELECT user_name, password FROM user WHERE user_name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (user *User) Validate() error {
	if user.Username == "" || user.Password == "" {
		return errors.New(errno.ErrValidation.Message)
	}
	return nil
}

func (user *User) Create() (int64, error) {
	id, err := Insert("INSERT INTO user (user_name, password) value (?,?)", user.Username, user.Password)
	if err != nil {
		return 1, err
	}

	return id, nil
}

func (user *User) UserToJson() string {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}

func (user *User) JsonToUser(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		return err
	}
	return nil
}
