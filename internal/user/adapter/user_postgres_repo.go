package adapter

import (
	"Clean-Sweep-Solutions_/internal/errors"
	"Clean-Sweep-Solutions_/internal/user/domain"
	"github.com/jackc/pgx"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) domain.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) Save(user *domain.User) error {
	queryStatement :=
		`INSERT INTO users(full_name, email, phone, password, role, created_at, updated_at) 
		VALUES ($1,$2,$3,$4)`

	_, err := u.db.Exec(queryStatement, user.FullName, user.Email, user.Phone,
		user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return errors.ErrFailedExecQuery
	}

	return nil
}
