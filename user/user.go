package user

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func ListUsers() []User {
	f, err := openFile()
	if err != nil {
		return []User{}
	}
	defer f.Close()
	var users []User
	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		return []User{}
	}
	return users
}

func AddUser(email string) (User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return User{}, err
	}
	f, err := openFile()
	if err != nil {
		return User{}, err
	}
	defer f.Close()
	var users []User
	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		return User{}, err
	}
	user := User{
		ID:    uid.String(),
		Email: email,
	}
	users = append(users, user)
	err = overwriteFile(users, f)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	f, err := openFile()
	if err != nil {
		return err
	}
	defer f.Close()
	var users []User
	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		return err
	}
	var filteredUsers []User
	for _, user := range users {
		if user.ID != id {
			filteredUsers = append(filteredUsers, user)
		}
	}
	err = overwriteFile(filteredUsers, f)
	if err != nil {
		return err
	}
	return nil
}

func openFile() (*os.File, error) {
	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func overwriteFile(users []User, f *os.File) error {
	b, err := json.Marshal(users)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	err = f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
