package storage

import (
	"context"
	"add/models"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Subscription() SubscriptionRepoI
	UserSubscription()  UserSubscriptionRepoI
}

type UserRepoI interface {
	Insert(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimarKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context,book *models.UpdateUser) error
	Delete(ctx context.Context, req *models.UserPrimarKey) error 
}

type SubscriptionRepoI interface {
	Insert(context.Context, *models.CreateSubscription) (string, error)
	GetByID(context.Context, *models.SubscriptionPrimeryKey) (*models.Subscription, error)
	GetList(ctx context.Context, req *models.GetListSubscriptionRequest) (*models.GetListSubscriptionResponse, error)
	Update(ctx context.Context, req *models.UpdateSubscription) error
	Delete(ctx context.Context, req *models.SubscriptionPrimeryKey) error
}
type UserSubscriptionRepoI interface{
	Insert(ctx context.Context, req *models.CreateUserSubscription) (string, error)
	GetById(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.SubscriptionUser, error) 
	GetByIdUser(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.UserSubscriptionJoin, error)
	Update(ctx context.Context, req *models.UpdateSubscriptionUser) error
	GetUserSubscription(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (*models.GetUserSubscription, error)
	UpdateUserSubscriptionStatus(ctx context.Context) error
	CheckAccess(cxt context.Context, req *models.SubscriptionUserPrimeryKey) (bool, error)
	MakeSubscriptionActive(ctx context.Context, req *models.SubscriptionUserPrimeryKey) (string,error)
}