package videonow

type User struct {
	Id uint64		`json:"id",bson:"id"`
	UserAdvType string	`json:"user_adv_type",bson:"user_adv_type"`
}