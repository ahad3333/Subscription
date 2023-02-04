package models

type UserSubscriptioHistoryPrimarKey struct {
	Id    string `json:"user_id"`
}



type CreateUserSubscriptioHistory struct {
	UserSubscriptionId string `json:"user_subscription_id"`
	Status    	string `json:"status"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"End_date"`
	Price  		float64 `json:"price"`
}

type UserSubscriptioHistory struct {
	Id          string `json:"id"`
	UserSubscriptionId string `json:"user_subscription_id"`
	Status    	string `json:"status"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"End_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Price  		float64 `json:"price"`
}

type UpdateUserSubscriptioHistory struct {
	UserSubscriptionId string `json:"user_subscription_id"`
	Status    	string `json:"status"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"End_date"`
	Price  		float64 `json:"price"`
}

type GetListUserSubscriptioHistoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserSubscriptioHistoryResponse struct {
	Count int32   `json:"count"`
	UserSubscriptioHistory []*UserSubscriptioHistory `json:"User_Subscriptio_History"`
}