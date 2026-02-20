package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserByEmail(ctx context.Context, email string, name string) (string, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) SaveUser(ctx context.Context, user models.User) error {
	sql := `INSERT INTO users(user_id, name, email) VALUES (:user_id, :name, :email)`
	_, err := u.db.NamedExecContext(ctx, sql, &user)
	return err
}

func (u *userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	sql := `SELECT user_id, name, email FROM users`
	users := []*models.User{}
	if err := u.db.SelectContext(ctx, &users, sql); err != nil {
		return nil, fmt.Errorf("query users error: %v", err)
	}

	return users, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string, name string) (string, error) {
	// sql := fmt.Sprintf(`
	// 	SELECT user_id, email, name
	// 	FROM users
	// 	WHERE email = '%s' AND name = '%s'
	// `, email, name)
	sql := `
		SELECT user_id, email, name
		FROM users
		WHERE email = $1 AND name = $2`
	//fmt.Println(sql)
	row := u.db.QueryRowContext(ctx, sql, email, name)
	var user models.User
	err := row.Scan(&user.UserID, &user.Email, &user.Name)
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}
