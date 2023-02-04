package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Insert(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO users (
		id,
		first_name,
		last_name,
		email,
		phone,
		created_at
	) VALUES ($1, $2, $3, $4, $5, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Phone,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepo) GetByID(ctx context.Context, req *models.UserPrimarKey) (*models.User, error) {

	query := `
		select 
			id,
			first_name,
			last_name,
			email,
			phone,
			updated_at,
			created_at
		from 
		users
		where id = $1
	`

	var (
		id        sql.NullString
		firstName sql.NullString
		lastName  sql.NullString
		email     sql.NullString
		phone     sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&firstName,
			&lastName,
			&email,
			&phone,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Id:        id.String,
		FirstName: firstName.String,
		LastName:  lastName.String,
		Phone:     phone.String,
		Email:     email.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return user, nil
}

func (r *UserRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   models.GetListUserResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			email,
			phone,
			updated_at,
			created_at
		FROM users
	
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListUserResponse{}, err
	}
	var (
		id        sql.NullString
		firstName sql.NullString
		lastName  sql.NullString
		email     sql.NullString
		phone     sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {

		err = rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&email,
			&phone,
			&createdAt,
			&updatedAt,
		)
		user := models.User{
			Id:        id.String,
			FirstName: firstName.String,
			LastName:  lastName.String,
			Phone:     phone.String,
			Email:     email.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		}
		if err != nil {
			return &models.GetListUserResponse{}, err
		}

		resp.Users = append(resp.Users, &user)

	}
	return &resp, nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.UpdateUser) error {
	query := `
		UPDATE 
			users 
		SET 
			first_name = $2,
			last_name = $3,
			email = $4,
			phone = $5,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, req *models.UserPrimarKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
