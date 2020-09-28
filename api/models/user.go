package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/events"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// UserRepository defines methods for Posts to interact with the database
type UserRepository interface {
	Get(meta http.Params) ([]domain.User, error)
	GetById(id int) (domain.User, error)
	GetOwner() (domain.User, error)
	Create(u *domain.User) (domain.User, error)
	Update(u *domain.User) (domain.User, error)
	Delete(id int) error
	CheckToken(token string) (domain.User, error)
	Exists(id int) bool
	ExistsByEmail(email string) bool
	Total() (int, error)
}

// UserStore defines the data layer for Users
type UserStore struct {
	db *sqlx.DB
}

// newUser - Construct
func newUser(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// Get all users
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no posts available.
func (s *UserStore) Get(meta http.Params) ([]domain.User, error) {
	const op = "UserRepository.Get"

	var u []domain.User
	q := fmt.Sprintf("SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id LEFT JOIN roles ON user_roles.role_id = roles.id ORDER BY users.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, meta.Page * meta.Limit)
	if err := s.db.Select(&u, q); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get users", Operation: op, Err: err}
	}

	if len(u) == 0 {
		return []domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "No users available", Operation: op}
	}

	return u, nil
}

// GetById returns a user by Id
// Returns errors.NOTFOUND if the user was not found by the given Id.
func (s *UserStore) GetById(id int) (domain.User, error) {
	const op = "UserRepository.GetById"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id LEFT JOIN roles ON user_roles.role_id = roles.id WHERE users.id = ?", id); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post with the ID: %d", id), Operation: op}
	}
	return u, nil
}

// GetOwner gets the owner of the site with the Id of 6
func (s *UserStore) GetOwner() (domain.User, error) {
	const op = "UserRepository.GetOwner"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id LEFT JOIN roles ON user_roles.role_id = roles.id WHERE roles.id = 6 LIMIT 1"); err != nil {
		return domain.User{}, fmt.Errorf("Could not get the owner of the site")
	}
	return u, nil
}

// Create user
func (s *UserStore) Create(u *domain.User) (domain.User, error) {
	const op = "UserRepository.Create"

	if exists := s.ExistsByEmail(u.Email); exists {
		return domain.User{}, fmt.Errorf("Could not create the user, the email, %v has been taken", u.Email)
	}

	hashedPassword, err := encryption.HashPassword(u.Password)
	if err != nil {
		log.Error(err)
		return domain.User{}, err
	}

	token := encryption.GenerateUserToken(u.FirstName + u.LastName, u.Email)

	userQ := "INSERT INTO users (uuid, first_name, last_name, email, password, website, facebook, twitter, linked_in, instagram, token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.db.Exec(userQ, uuid.New().String(), u.FirstName, u.LastName, u.Email, hashedPassword, u.Website, u.Facebook, u.Twitter, u.Linkedin, u.Instagram, token)
	if err != nil {
		log.Error(err)
		return domain.User{}, fmt.Errorf("Could not create the user with the name: %v %v", u.FirstName, u.LastName)
	}

	id, err := c.LastInsertId()
	if err != nil {
		log.Error(err)
		return domain.User{}, fmt.Errorf("Could not get the newly created user ID with the name: %v %v", u.FirstName, u.LastName)
	}
	u.Id = int(id)

	roleQ := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err = s.db.Exec(roleQ, id, u.Role.Id)
	if err != nil {
		log.Error(err)
		return domain.User{}, fmt.Errorf("Could not create the user role for user with the name: %v %v", u.FirstName, u.LastName)
	}

	if u.Role.Id != 6 {
		ve, err := events.NewVerifyEmail()
		if err != nil {
			log.Error(err)
			return domain.User{}, err
		}

		// TODO: Get app title from options model
		err = ve.Send(u, "Verbis")
		if err != nil {
			log.Error(err)
			return domain.User{}, err
		}
	}

	return *u, nil
}

// Update user
func (s *UserStore) Update(u *domain.User) (domain.User, error) {
	const op = "UserRepository.Update"

	_, err := s.GetById(u.Id)
	if err != nil {
		log.Info(err)
		return domain.User{}, err
	}

	hashedPassword, err := encryption.HashPassword(u.Password)
	if err != nil {
		log.Error(err)
		return domain.User{}, err
	}

	userQ := "UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ?, website = ?, facebook = ?, twitter = ?, linked_in = ?, instagram = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(userQ, u.FirstName, u.LastName, u.Email, hashedPassword, u.Website, u.Facebook, u.Twitter, u.Linkedin, u.Instagram, u.Id)
	if err != nil {
		log.Error(err)
		return domain.User{}, fmt.Errorf("Could not update the user with the name: %v %v", u.FirstName, u.LastName)
	}

	roleQ := "UPDATE users SET (role_id) VALUES (?)"
	_, err = s.db.Exec(roleQ, u.Role.Id)
	if err != nil {
		log.Error(err)
		return domain.User{}, fmt.Errorf("Could not update the user role for user with the name: %v %v", u.FirstName, u.LastName)
	}

	return *u, nil
}

// Delete user
func (s *UserStore) Delete(id int) error {
	const op = "UserRepository.Delete"

	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM users WHERE id = ?", id); err != nil {
		return fmt.Errorf("Could not delete user with the ID of %v - %w", id, err)
	}

	return nil
}

// Get the user by Token
func (s *UserStore) CheckToken(token string) (domain.User, error) {
	const op = "UserRepository.CheckToken"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id LEFT JOIN roles ON user_roles.role_id = roles.id WHERE users.token = ?", token); err != nil {
		return domain.User{}, fmt.Errorf("Could not get user with token: %v", token)
	}
	return u, nil
}

// Check if the user record exists by ID
func (s *UserStore) Exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM users WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Check if the user record exists by email
func (s *UserStore) ExistsByEmail(email string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM users WHERE email = ?)", email).Scan(&exists)
	return exists
}

// Get the total number of posts
func (s *UserStore) Total() (int, error) {
	const op = "UserRepository.Total"
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total); err != nil {
		return -1, fmt.Errorf("Could not get the total number of users")
	}
	return total, nil
}


