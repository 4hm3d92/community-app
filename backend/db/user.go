package db

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/4hm3d92/community-app/backend/models"
	"github.com/4hm3d92/community-app/backend/utils"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}

	rows, err := db.Pool.Query(context.Background(), "SELECT id, username, name, role, enabled, created_on, updated_on FROM users ORDER BY id DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Role, &user.Enabled, &user.CreatedOn, &user.UpdatedOn)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
	passwordHash, err := utils.GenerateFromPassword(user.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, username, password, role, enabled) VALUES ($1, $2, $3, $4, $5);`
	_, err = db.Pool.Exec(context.Background(), query, user.Name, user.Username, passwordHash, user.Role, user.Enabled)
	if err != nil {
		return err
	}

	return nil
}

/*
func (db Database) GetUserById(userId int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1;`
	row := db.Pool.QueryRow(context.Background(), query, userId)
	switch err := row.Scan(&user.ID, &user.Name, &user.Role, &user.LastName); err {
	case pgx.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}
*/

func (db Database) GetUserByUsername(user *models.UserLoginRequest) (models.User, error) {
	userData := models.User{}
	query := `SELECT id, password, role FROM users WHERE username = $1;`
	row := db.Pool.QueryRow(context.Background(), query, &user.Username)
	switch err := row.Scan(&userData.ID, &userData.Password, &userData.Role); err {
	case pgx.ErrNoRows:
		return userData, ErrNoMatch
	default:
		return userData, err
	}
}

/*
func (db Database) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Pool.Exec(query, userId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
*/

func (db Database) UpdateUser(userId int, userData models.User) error {
	query := `UPDATE users SET name=$1, role=$2, enabled=$3 WHERE id=$3 RETURNING id, name, updated_on;`
	_, err := db.Pool.Exec(context.Background(), query, userData.Name, userData.Role, userData.Enabled, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNoMatch
		}
		return err
	}

	return nil
}

func (db Database) SetPassword(userData models.User) error {
	//user := models.User{}
	query := `UPDATE users SET password=$1 WHERE id=$2;`
	_, err := db.Pool.Exec(context.Background(), query, userData.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNoMatch
		}
		return err
	}

	return nil
}
