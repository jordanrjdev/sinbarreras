package implementation

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JGurus/sin-barreras-backend/models"
)

const (
	migrateUser = `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		username VARCHAR(25) NOT NULL,
		avatar VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	 )`
	createUser    = "INSERT INTO users (username, avatar) VALUES(?,?)"
	updateUser    = "UPDATE users SET username = ?, avatar = ?, updated_at = now() WHERE id = ? AND deleted_at IS NULL"
	deleteUser    = "UPDATE users SET deleted_at = now() WHERE id = ?"
	getAll        = "SELECT * FROM users WHERE deleted_at IS NULL"
	getByID       = "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"
	getByUsername = "SELECT * FROM users WHERE username = ? AND deleted_at IS NULL"
)

//UserImpl .
type UserImpl struct {
	db *sql.DB
}

//NewUserImpl .
func NewUserImpl(db *sql.DB) *UserImpl {
	return &UserImpl{db}
}

//Migrate .
func (dao *UserImpl) Migrate() error {
	stmt, err := dao.db.Prepare(migrateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración exitosa")
	return nil
}

//Create .
func (dao *UserImpl) Create(u *models.User) error {
	stmt, err := dao.db.Prepare(createUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Username, u.Avatar)
	if err != nil {
		return err
	}
	fmt.Println("Usuario creado exitosamente")
	return nil
}

//Update .
func (dao *UserImpl) Update(id int, u *models.User) error {
	stmt, err := dao.db.Prepare(updateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Username, u.Avatar, id)
	if err != nil {
		return err
	}
	fmt.Println("Usuario actualizado exitosamente")
	return nil
}

//Delete .
func (dao *UserImpl) Delete(id int) error {
	stmt, err := dao.db.Prepare(deleteUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	println("Usuario eliminado exitosamente")
	return nil
}

//GetAll .
func (dao *UserImpl) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	stmt, err := dao.db.Prepare(getAll)
	if err != nil {
		return users, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	fmt.Println("Usuarios consultados exitosamente")
	return users, nil
}

//GetByID .
func (dao *UserImpl) GetByID(id int) (models.User, error) {
	user := models.User{}
	stmt, err := dao.db.Prepare(getByID)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Username, &user.Avatar, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetByUsername return a User by username
func (dao UserImpl) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	stm, err := dao.db.Prepare(getByUsername)
	if err != nil {
		return user, err
	}
	defer stm.Close()
	row := stm.QueryRow(username)
	err = row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
