package models

type SubscriptionPrimeryKey struct {
	Id string `json:"id"`
}

type CreateSubscription  struct{
	Title_ru      string `json:"title_ru"`
	Title_en      string `json:"title_en"`
	Title_uz      string `json:"title_uz"`
	Price		  float64 `json:"price"`
	Duration_type string `json:"duration_type"`
	Duration   	  float64 `json:"duration"`
	Free_trial    float64 `json:"free_trial"`
	UpdatedAt     string `json:"updated_at"`
}

type Subscription struct {
	Id            string `json:"id"`
	Title_ru      string `json:"title_ru"`
	Title_en      string `json:"title_en"`
	Title_uz      string `json:"title_uz"`
	Price		  float64 `json:"price"`
	Duration_type string `json:"duration_type"`
	Duration   	  float64 `json:"duration"`
	Is_activen    bool   `json:"is_active"`
	Free_trial    float64 `json:"free_trial"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeleteAt      string `json:"delete_at"`
}


type UpdateSubscription struct {
	Id   string `json:"id"`
	Title_ru      string `json:"title_ru"`
	Title_en      string `json:"title_en"`
	Title_uz      string `json:"title_uz"`
	Price		  float64 `json:"price"`
	Duration_type string `json:"duration_type"`
	Duration   	  float64 `json:"duration"`
	Free_trial    float64 `json:"free_trial"`
	UpdatedAt     string `json:"updated_at"`
}

type UpdateSubscriptionSwag struct {
	Title_ru      string `json:"title_ru"`
	Title_en      string `json:"title_en"`
	Title_uz      string `json:"title_uz"`
	Price		  float64 `json:"price"`
	Duration_type string `json:"duration_type"`
	Duration   	  float64 `json:"duration"`
	Free_trial    float64 `json:"free_trial"`
}

type GetListSubscriptionRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListSubscriptionResponse struct {
	Count int32   `json:"count"`
	Subscriptions []*Subscription `json:"Subscriptio"`
}
