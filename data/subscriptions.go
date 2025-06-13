package data

type SubscriptionsResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Meta          Meta           `json:"meta"`
}

type Subscription struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Meta struct {
	Count    int `json:"count"`
	Offset   int `json:"offset"`
	PgaeSize int `json:"pageSize"`
	NextPage int `json:"nextPage"`
}
