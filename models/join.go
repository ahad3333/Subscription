package models

type UserSubscriptionJoin struct {
	Purchased Subscription `json:"purchased"`
	UnPurchased []*Subscription `json:"unpurchased"`

}


type GetUserSubscription struct{
	Status    	 string `json:"status"`
	Start_date   string `json:"start_date"`
	End_date     string `json:"End_date"`
	Price  		 float64 `json:"price"`
	Title_ru      string `json:"title_ru"`
	Title_en      string `json:"title_en"`
	Title_uz      string `json:"title_uz"`
	Duration   	  float64 `json:"duration"`
}