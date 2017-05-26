package videonow

type CampaignCategory struct {
	Id uint64			`json:"id",bson:"id"`//INT(11) DEFAULT '0' NOT NULL,
	PayEvent uint64			`json:"pay_event",bson:"pay_event"`//INT(10) unsigned NOT NULL,
	PayEventCreated uint64		`json:"pay_event_created",bson:"pay_event_created"`//INT(10) unsigned NOT NULL,
	OwnerEvent uint64		`json:"owner_event",bson:"owner_event"`//INT(10) unsigned DEFAULT '0' NOT NULL,
	OwnerEventCreated uint64	`json:"owner_event_created",bson:"owner_event_created"`//INT(10) unsigned DEFAULT '0' NOT NULL,
	EventPrice float64		`json:"event_price",bson:"event_price"`//DECIMAL(10,4) unsigned NOT NULL,
	EventPriceReal float64		`json:"event_price_real",bson:"event_price_real"`//DECIMAL(10,4) unsigned NOT NULL,
}