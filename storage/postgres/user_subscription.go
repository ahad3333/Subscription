package postgres

import (
	"add/models"
	"add/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserSubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewUserSubscriptionRepo(db *pgxpool.Pool) *UserSubscriptionRepo {
	return &UserSubscriptionRepo{
		db: db,
	}

}

func (r *UserSubscriptionRepo) Insert(ctx context.Context, req *models.CreateUserSubscription) (string, error) {
	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO subscription_user (
		id,
		user_id,
		subscription_id,
		free_trial_start_date,
		free_trial_end_date,
		status,
		send_notification,
		created_at
	)
		select 
				$1,
				$3,
				s.id,
				now(),
				now() + free_trial * '1 days' :: interval,
				case
				when s.free_trial > 0 then 'FREE_TRIAL'
				else 'INACTIVE'
				end as status,
				false,
				created_at
		from subscription as s 
		where s.id = $2
`

	_, err := r.db.Exec(ctx, query, id, req.Subscription_id, req.User_id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserSubscriptionRepo) GetByIdUser(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.UserSubscriptionJoin, error) {
	query := `
		select 
			s.id,
			s.title_ru,
			s.title_en,
			s.title_uz,
			s.price,
			s.duration_type,
			s.duration,
			s.is_active,
			s.free_trial
		from subscription_user as su 
		join subscription as s on s.id = su.subscription_id
		where su.user_id = $1`

	var (
		id            sql.NullString
		title_ru      sql.NullString
		title_en      sql.NullString
		title_uz      sql.NullString
		price         sql.NullFloat64
		duration_type sql.NullString
		duration      sql.NullFloat64
		is_active     sql.NullBool
		free_trial    sql.NullFloat64
	)

	err := r.db.QueryRow(ctx, query, req.User_id).
		Scan(
			&id,
			&title_ru,
			&title_en,
			&title_uz,
			&price,
			&duration_type,
			&duration,
			&is_active,
			&free_trial,
		)

	if err != nil {
		return nil, err
	}

	purchased := &models.Subscription{
		Id:            id.String,
		Title_ru:      title_ru.String,
		Title_en:      title_en.String,
		Title_uz:      title_uz.String,
		Price:         price.Float64,
		Duration_type: duration_type.String,
		Duration:      duration.Float64,
		Is_activen:    true,
		Free_trial:    free_trial.Float64,
	}
	queryun := `
		select 
			s.title_ru,
			s.title_en,
			s.title_uz,
			s.price,
			s.duration_type,
			s.duration,
			s.is_active,
			s.free_trial
		from subscription_user as su 
		join subscription as s on s.id = su.subscription_id
		join users as u on u.id = su.user_id
		where su.user_id !=u.id
	`
	UserSubscriptionJoin := models.UserSubscriptionJoin{}

	rows, err := r.db.Query(ctx, queryun)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&id,
			&title_ru,
			&title_en,
			title_uz,
			&price,
			&duration_type,
			&duration,
			&is_active,
			&free_trial,
		)
		unpurchased := models.Subscription{
			Id:            id.String,
			Title_ru:      title_ru.String,
			Title_en:      title_en.String,
			Title_uz:      title_uz.String,
			Price:         price.Float64,
			Duration_type: duration_type.String,
			Duration:      duration.Float64,
			Is_activen:    false,
			Free_trial:    free_trial.Float64,
		}

		UserSubscriptionJoin.UnPurchased = append(UserSubscriptionJoin.UnPurchased, &unpurchased)
	}
	UserSubscriptionJoin.Purchased = *purchased

	return &UserSubscriptionJoin, nil
}

func (r *UserSubscriptionRepo) GetUserSubscription(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.GetUserSubscription, error) {
	query := `
		select 
			uh.status,
			uh.start_date,
			uh.end_date,
			uh.price,
			s.title_en,
			s.title_ru,
			s.title_uz,
			s.duration
		from subscription_user as su
		join user_subscription_history as uh on uh.user_subscription_id = su.id
		join subscription as s on s.id = su.subscription_id
		WHERE su.id = $1 
		Order by  uh.status, uh.start_date, uh.end_date, uh.price DESC
		`

	var (
		Status     sql.NullString
		Start_date sql.NullString
		End_date   sql.NullString
		Price      sql.NullFloat64
		Title_ru   sql.NullString
		Title_en   sql.NullString
		Title_uz   sql.NullString
		Duration   sql.NullFloat64
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&Status,
			&Start_date,
			&End_date,
			&Price,
			&Title_en,
			&Title_ru,
			&Title_uz,
			&Duration,
		)

	if err != nil {
		return nil, err
	}

	subscription := &models.GetUserSubscription{
		Status:     Status.String,
		Start_date: Start_date.String,
		End_date:   End_date.String,
		Price:      Price.Float64,
		Title_ru:   Title_ru.String,
		Title_en:   Title_en.String,
		Title_uz:   Title_uz.String,
		Duration:   Duration.Float64,
	}

	return subscription, nil
}

func (r *UserSubscriptionRepo) GetById(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.SubscriptionUser, error) {
	query := `
			SELECT
				id,
				user_id,
				subscription_id,
				free_trial_start_date,
				free_trial_end_date,
				created_at,
				updated_at,
				status,
				send_notification
			FROM
				subscription_user
			WHERE id = $1`

	var (
		id                    sql.NullString
		user_id               sql.NullString
		subscription_id       sql.NullString
		free_trial_start_date sql.NullString
		free_trial_end_date   sql.NullString
		created_at            sql.NullString
		updated_at            sql.NullString
		status                sql.NullString
		send_notification     sql.NullBool
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&user_id,
			&subscription_id,
			&free_trial_start_date,
			&free_trial_end_date,
			&created_at,
			&updated_at,
			&status,
			&send_notification,
		)

	if err != nil {
		return nil, err
	}

	subscriptionuser := &models.SubscriptionUser{
		Id:                    id.String,
		User_id:               user_id.String,
		Subscription_id:       subscription_id.String,
		Free_trial_start_date: free_trial_start_date.String,
		Free_trial_end_date:   free_trial_end_date.String,
		CreatedAt:             created_at.String,
		UpdatedAt:             updated_at.String,
		Status:                status.String,
		Send_notification:     send_notification.Bool,
	}

	return subscriptionuser, nil
}

func (r *UserSubscriptionRepo) Update(ctx context.Context, req *models.UpdateSubscriptionUser) error {

	query := `
		UPDATE 
			subscription_user
		SET 
			user_id = $2,
			subscription_id = $3,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.User_id,
		req.Subscription_id,
	)

	if err != nil {
		return err
	}

	// var price float64
	// querystatus := `
	// 		select
	// 			s.price
	// 		from subscription as s
	// 		join subscription_user as su on su.subscription_id =s.id
	// 		WHERE su.id = $1
	// 		`
	// err = r.db.QueryRow(ctx, querystatus, req.Id).Scan(
	// 	&price,
	// )
	// if err != nil {
	// 	return err
	// }

	// if price > 0 {
	// 	status := `
	// 	UPDATE
	// 		subscription_user
	// 	SET
	// 	status = 'ACTIVE'
	// 	`
	// 	_, err = r.db.Exec(ctx, status)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (r *UserSubscriptionRepo) UpdatePatch(ctx context.Context, req *models.PatchSubscriptionUser) error {

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Data) == 0 {
		return errors.New("no updates provided")
	}

	req.Data["id"] = req.Id

	for key := range req.Data {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Data)-1 {
			set += ", "
		}
		ind++
	}

	query = `
			UPDATE
				subscription_user
		` + set + ` , updated_at = now()
			WHERE
				id = :id
		`

	query, args := helper.ReplaceQueryParams(query, req.Data)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserSubscriptionRepo) UpdateUserSubscriptionStatus(ctx context.Context) error {

	query := `
		UPDATE 
			subscription_user as s
		SET status  =  case  
			WHEN status = 'FREE' then 'FREE_TRIAL'
			WHEN status = 'ACTIVE' then 'INACTIVE'
		End
		where id in (
			select user_subscription_id from user_subscription_history 
		Where status ='ACTIVE' or status ='FREE_TRIAL' and 
			now() > end_date)
	`
	_, err := r.db.Exec(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserSubscriptionRepo) CheckAccess(cxt context.Context, req *models.SubscriptionUserPrimeryKey) (bool, error) {
	var status models.SubscriptionUser
	query := `
	select
		status
	from subscription_user 
	where user_id = $1`
	r.db.QueryRow(cxt, query, req.User_id).Scan(
		status.Status,
	)

	if status.Status == "ACTIVE" || status.Status == "FREE_TRIAL" {
		status.Send_notification = true
	}
	return status.Send_notification, nil
}

func (r *UserSubscriptionRepo) MakeSubscriptionActive(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (string,error) {

	query := `
		UPDATE 
			subscription_user
		SET 
			status = 'ACTIVE'
			updated_at = now()
		WHERE user_id = $1
	`
	_,err :=r.db.Exec(ctx, query, req.User_id)
	if err != nil {
		return "Error updedte status:",err
	}

	return "Update status:",nil
}
