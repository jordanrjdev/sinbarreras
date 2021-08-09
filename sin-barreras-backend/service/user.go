package service

import "github.com/JGurus/sin-barreras-backend/models"

type User struct {
	db UserServices
}

func NewUser(db UserServices) User {
	return User{db}
}

func (s User) Migrate() error {
	err := s.db.Migrate()
	if err != nil {
		return nil
	}
	return nil
}

func (s User) Create(u *models.User) error {
	user := u
	err := s.db.Create(user)
	if err != nil {
		return nil
	}
	return nil
}

func (s User) Update(id int, u *models.User) error {
	user := u
	err := s.db.Update(id, user)
	if err != nil {
		return nil
	}
	return nil
}

func (s User) Delete(id int) error {
	err := s.db.Delete(id)
	if err != nil {
		return nil
	}
	return nil
}

func (s User) GetAll() ([]models.User, error) {
	var err error
	var users = make([]models.User, 0)
	users, err = s.db.GetAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

//GetByID .
func (s User) GetByID(id int) (models.User, error) {
	user := models.User{}
	user, err := s.db.GetByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByUsername .
func (s User) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	user, err := s.db.GetByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}
