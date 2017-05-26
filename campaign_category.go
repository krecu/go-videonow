package videonow

type CampaignCategory struct {
	Id uint64			//INT(11) DEFAULT '0' NOT NULL,
	PayEvent uint64			//INT(10) unsigned NOT NULL,
	PayEventCreated uint64		//INT(10) unsigned NOT NULL,
	OwnerEvent uint64		//INT(10) unsigned DEFAULT '0' NOT NULL,
	OwnerEventCreated uint64	//INT(10) unsigned DEFAULT '0' NOT NULL,
	EventPrice float64		//DECIMAL(10,4) unsigned NOT NULL,
	EventPriceReal float64		//DECIMAL(10,4) unsigned NOT NULL,
}