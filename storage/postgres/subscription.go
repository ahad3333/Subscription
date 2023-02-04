package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepo(db *pgxpool.Pool) *SubscriptionRepo {
	return &SubscriptionRepo{
		db: db,
	}

}

func (r *SubscriptionRepo) Insert(ctx context.Context, req *models.CreateSubscription) (string, error) {
	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO subscription (
		id,
		title_ru,
		title_en,
		title_uz,
		price,
		duration_type,
		duration,
		free_trial,
		created_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8,now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Title_ru,
		req.Title_en,
		req.Title_uz,
		req.Price,
		req.Duration_type,
		req.Duration,
		req.Free_trial,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SubscriptionRepo) GetByID(ctx context.Context, req *models.SubscriptionPrimeryKey) (*models.Subscription, error) {
	query := `
			SELECT
				id,
				title_ru,
				title_en,
				title_uz,
				price,
				duration_type,
				duration,
				free_trial,
				updated_at,
				created_at,
				delete_at
			FROM
				subscription
			WHERE id = $1`

	var (
		id            sql.NullString
		title_ru      sql.NullString
		title_en      sql.NullString
		title_uz      sql.NullString
		price         sql.NullFloat64
		duration_type sql.NullString
		duration      sql.NullFloat64
		free_trial    sql.NullFloat64
		createdAt     sql.NullString
		updatedAt     sql.NullString
		deleteAt      sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&title_ru,
			&title_en,
			&title_uz,
			&price,
			&duration_type,
			&duration,
			&free_trial,
			&createdAt,
			&updatedAt,
			&deleteAt,
		)

	if err != nil {
		return nil, err
	}

	subscription := &models.Subscription{
		Id:            id.String,
		Title_ru:      title_ru.String,
		Title_en:      title_en.String,
		Title_uz:      title_uz.String,
		Price:         price.Float64,
		Duration_type: duration_type.String,
		Duration:      duration.Float64,
		Free_trial:    free_trial.Float64,
		CreatedAt:     createdAt.String,
		UpdatedAt:     updatedAt.String,
		DeleteAt:      deleteAt.String,
	}

	return subscription, nil
}

func (r *SubscriptionRepo) GetList(ctx context.Context, req *models.GetListSubscriptionRequest) (*models.GetListSubscriptionResponse, error) {

	var (
		resp   models.GetListSubscriptionResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
				id,
				title_ru,
				title_en,
				title_uz,
				price,
				duration_type,
				duration,
				free_trial,
				updated_at,
				created_at ,
				delete_at
		FROM
			subscription
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
		return &models.GetListSubscriptionResponse{}, err
	}
	var (
		id            sql.NullString
		title_ru      sql.NullString
		title_en      sql.NullString
		title_uz      sql.NullString
		price         sql.NullFloat64
		duration_type sql.NullString
		duration      sql.NullFloat64
		free_trial    sql.NullFloat64
		createdAt     sql.NullString
		updatedAt     sql.NullString
		deleteAt      sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&id,
			&title_ru,
			&title_en,
			&title_uz,
			&price,
			&duration_type,
			&duration,
			&free_trial,
			&createdAt,
			&updatedAt,
			&deleteAt,
		)
		if err != nil {
			return &models.GetListSubscriptionResponse{}, err
		}

		subscription := &models.Subscription{
			Id:            id.String,
			Title_ru:      title_ru.String,
			Title_en:      title_en.String,
			Title_uz:      title_uz.String,
			Price:         price.Float64,
			Duration_type: duration_type.String,
			Duration:      duration.Float64,
			Free_trial:    free_trial.Float64,
			CreatedAt:     createdAt.String,
			UpdatedAt:     updatedAt.String,
			DeleteAt:      deleteAt.String,
		}

		resp.Subscriptions = append(resp.Subscriptions, subscription)

	}

	return &resp, nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, req *models.UpdateSubscription) error {

	query := `
		UPDATE 
			subscription
		SET 
			title_ru = $2,
			title_en = $3,
			title_uz = $4,
			price = $5,
			duration_type = $6,
			duration = $7,
			free_trial = $8,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.Title_ru,
		req.Title_en,
		req.Title_uz,
		req.Price,
		req.Duration_type,
		req.Duration,
		req.Free_trial,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, req *models.SubscriptionPrimeryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM subscription WHERE id = $1 ", req.Id)

	if err != nil {
		return err
	}

	return nil
}
