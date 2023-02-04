package models

type SubscriptionUserPrimeryKey struct {
	Id                    string `json:"id"`
	User_id     		  string `json:"user_id"`
}

type SubscriptionUser struct {
	Id                    string `json:"id"`
	User_id     		  string `json:"user_id"`
	Subscription_id       string `json:"subscription_id"`
	Free_trial_start_date string `json:"free_trial_start_date"`
	Free_trial_end_date	  string `json:"free_trial_and_date"`
	Status 				  string `json:"status"`
	Send_notification     bool 	 `json:"send_notification"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}

type CreateUserSubscription struct {
	Id                    string `json:"id"`
	User_id     		  string `json:"user_id"`
	Subscription_id       string `json:"subscription_id"`
	Free_trial_start_date string `json:"free_trial_start_date"`
	Free_trial_end_date	  string `json:"free_trial_and_date"`
	Status 				  string `json:"status"`
	UpdatedAt             string `json:"updated_at"`
}


type UpdateSubscriptionUser struct {
	Id   				  string `json:"id"`
	User_id     		  string `json:"user_id"`
	Subscription_id       string `json:"subscription_id"`
	Status 				  string `json:"status"`
	UpdatedAt     		  string `json:"updated_at"`
}

type UpdateSubscriptionUserSwag struct {
	User_id     		  string `json:"user_id"`
	Subscription_id       string `json:"subscription_id"`
	Status 				  string `json:"status"`
}

type GetListSubscriptionUserRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListSubscriptionUserResponse struct {
	Count int32   `json:"count"`
	SubscriptionUsers []*SubscriptionUser `json:"SubscriptioUser"`
}

type PatchSubscriptionUser struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}


type Checkaccess struct {
	Send_notification     bool 	 `json:"send_notification"`
}

type MakeSubscriptionActive struct {
	Status string `json:"status"`
}