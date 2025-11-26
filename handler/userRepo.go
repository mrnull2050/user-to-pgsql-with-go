package handler

import (
	"database/sql"
	"log"
	"pgsql/model"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) Create(um model.UserModel) error {
	_, err := u.DB.Exec(`INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`, um.Name, um.Email, um.Age)
	if err != nil {
		log.Println("Exec error:", err)
		return err
	}
	return err
}

func (u *UserRepo) GetAll() ([]model.UserModel, error) {
	rows, err := u.DB.Query(`SELECT id , name , email , age FROM users`)
	if err != nil {
		log.Fatal("error while getting query ")
		return nil, err
	}
	defer rows.Close()

	var users []model.UserModel

	for rows.Next() {
		var u model.UserModel
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		users = append(users, u)
	}
	return users, nil
}
func (r *UserRepo) GetById(id int) (model.UserModel, error) {
	var u model.UserModel
	err := r.DB.QueryRow(`SELECT id , name , email , age FROM users WHERE id=$1`, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	return u, err
}
func (r *UserRepo) Updata(u model.UserModel) error {
	_, err := r.DB.Exec(`UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4
    `, u.Name, u.Email, u.Age, u.ID)
	return err
}
func (r *UserRepo) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
