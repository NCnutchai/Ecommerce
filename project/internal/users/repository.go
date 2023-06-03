package users

import (
	"project/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetIDByUsername(username string) (*int, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByUsernameAndPassword(username string, password string) (*User, error)
	InsertUser(user RegisterUser) (*User, error)
}

type userRepository struct {
	db *database.PostgresDB
}

func NewUserRepository(db *database.PostgresDB) Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetIDByUsername(username string) (*int, error) {
	var id int
	err := r.db.Conn.QueryRow(`
		SELECT id FROM auth_user Where username = $1`, username).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *userRepository) GetUserByUsername(username string) (*User, error) {
	var user User

	err := r.db.Conn.QueryRow(`
		SELECT id, username, password, last_login, firstname, lastname, phone_number, email, is_superuser, is_active, created_at, updated_at 
		FROM auth_user Where username =$1`, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Firstname, &user.Lastname,
		&user.PhoneNumber, &user.Email, &user.IsSuperuser, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Conn.Exec(`UPDATE  auth_user SET last_login = current_timestamp Where username = $1`, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsernameAndPassword(username string, password string) (*User, error) {
	var user User

	err := r.db.Conn.QueryRow(`
        SELECT id, username, password, last_login, firstname, lastname, phone_number, email, is_superuser, is_active, created_at, updated_at 
        FROM auth_user WHERE username = $1`, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Firstname, &user.Lastname,
		&user.PhoneNumber, &user.Email, &user.IsSuperuser, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) InsertUser(register RegisterUser) (*User, error) {
	var user User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Conn.Exec(`
        INSERT INTO auth_user (username, password, firstname, lastname, phone_number, email)
        VALUES ($1, $2, $3, $4, $5, $6)`,
		register.Username, hashedPassword, register.Firstname, register.Lastname, register.PhoneNumber, register.Email)
	if err != nil {
		return nil, err
	}

	err = r.db.Conn.QueryRow(`
        SELECT id, username, password, last_login, firstname, lastname, phone_number, email, is_superuser, is_active, created_at, updated_at 
        FROM auth_user WHERE username = $1`, register.Username).Scan(
		&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Email, &user.IsSuperuser, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
