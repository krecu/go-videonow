package videonow

import "time"
type Campaign struct {

	Id uint64			`json:"id",bson:"id"`//INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
	Name string			`json:"name",bson:"name"`//VARCHAR(80) DEFAULT '' NOT NULL,
	Type string			`json:"type",bson:"type"`//VARCHAR(30) NOT NULL,
	Status string 			`json:"status",bson:"status"`//ENUM('no_data', 'stop', 'work', 'end') DEFAULT 'no_data' NOT NULL,
	Archived bool 			`json:"archived",bson:"archived"`//ENUM('', 'on') NOT NULL,
	Belong string 			`json:"belong",bson:"belong"`//ENUM('owner', 'adv', 'external', 'service', 'owner_adv', 'owner_external', 'rtb'),
	IsApproved bool 		`json:"is_approved",bson:"is_approved"`//ENUM('null', 'new', 'ok', 'no') DEFAULT 'null' NOT NULL,
	IsPayment bool 			`json:"is_payment",bson:"is_payment"`//ENUM('no', 'yes') DEFAULT 'no' NOT NULL,
	IsApprovedText string 		`json:"is_approved_text",bson:"is_approved_text"`//TEXT NOT NULL,
	AutoStart bool 			`json:"auto_start",bson:"auto_start"`//ENUM('on', 'off') DEFAULT 'on',
	Discount float64 		`json:"discount",bson:"discount"`//FLOAT UNSIGNED DEFAULT '0' NOT NULL,
	OpenStatFlag bool 		`json:"open_stat_flag",bson:"open_stat_flag"`//ENUM('on', 'off') DEFAULT 'off' NOT NULL,
	PlannedStartDate time.Time 	`json:"planned_start_date",bson:"planned_start_date"`//DATE NOT NULL,
	PlannedStopDate  time.Time 	`json:"planned_stop_date",bson:"planned_stop_date"`//DATE NOT NULL,
	DateEnd time.Time 		`json:"date_end",bson:"date_end"`//DATE NOT NULL,
	DayBudget float64 		`json:"day_budget",bson:"day_budget"`//FLOAT DEFAULT '0' NOT NULL COMMENT 'Желаемый бюджет на один день',
	DayEvents uint64 		`json:"day_events",bson:"day_events"`//INT(11) NOT NULL COMMENT 'Желаемое количество событий в день',
	ShowsPerUser uint64 		`json:"shows_per_user",bson:"shows_per_user"`//TINYINT(3) unsigned DEFAULT '3' NOT NULL,
	Container string 		`json:"container",bson:"container"`//VARCHAR(30) NOT NULL,
	GeoType int 			`json:"geo_type",bson:"geo_type"`//TINYINT(3) unsigned DEFAULT '0' NOT NULL COMMENT 'Тип гео-таргетинга',
	CreativeType string 		`json:"creative_type",bson:"creative_type"`//ENUM('inner', 'wrapper') NOT NULL,
	Url string 			`json:"url",bson:"url"`//TEXT NOT NULL,
	LinkTitle string 		`json:"link_title",bson:"link_title"`//VARCHAR(255) NOT NULL,
	LinkTarget string 		`json:"link_target",bson:"link_target"`//ENUM('', '_self') NOT NULL,
	IsPriority bool 		`json:"is_priority",bson:"is_priority"`//ENUM('', 'on') NOT NULL,
	IsWordsTargeted bool 		`json:"is_words_targeted",bson:"is_words_targeted"`//ENUM('', 'on') NOT NULL,
	WordsTargeted string 		`json:"words_targeted",bson:"words_targeted"`//TEXT NOT NULL,
	BasePrice float64 		`json:"base_price",bson:"base_price"`//DECIMAL(10,4) COMMENT 'цена события до наценок за таргетинг',
	TimeTargeting bool 		`json:"time_targeting",bson:"time_targeting"`//TINYINT(1) unsigned DEFAULT '0' NOT NULL,
	TargetUrlType string 		`json:"target_url_type",bson:"target_url_type"`//ENUM('target', 'ban'),
	TargetListUrl string 		`json:"target_list_url",bson:"target_list_url"`//TEXT NOT NULL,
	DeviceTargeted bool 		`json:"device_targeted",bson:"device_targeted"`//TINYINT(1) unsigned DEFAULT '0' NOT NULL,
	Cpm float64 			`json:"cpm",bson:"cpm"`//DECIMAL(11,2) DEFAULT '0.00' NOT NULL COMMENT 'Cpm для external-кампаний',
	IsServerRequest bool 		`json:"is_server_request",bson:"is_server_request"`//TINYINT(1) unsigned DEFAULT '0' NOT NULL COMMENT 'Поддерживает ли внешняя кампания серверные запросы рекламы', // todo - пока ненужен
	AdServer int64 			`json:"ad_server",bson:"ad_server"`//TINYINT(4) unsigned, // todo - пока ненужен
	ServerUrl string 		`json:"server_url",bson:"server_url"`//VARCHAR(500), // todo - пока ненужен
	AdFoxKey string 		`json:"ad_fox_key",bson:"ad_fox_key"`//VARCHAR(50) DEFAULT '' NOT NULL, // todo - пока ненужен
	OwnerId int64 			`json:"owner_id",bson:"owner_id"`//INT(10) unsigned, // todo - пока ненужен
	Created time.Time 		`json:"created",bson:"created"`//TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ExtPriority int64 		`json:"ext_priority",bson:"ext_priority"`//nINT(10) unsigned DEFAULT '0',

	User User			`json:"user",bson:"user"`// refference by user_id
	Category []CampaignCategory	`json:"category",bson:"category"`//
	Region	 []Region		`json:"region",bson:"region"`
	Countries []Country		`json:"countries",bson:"countries"`
	Profiles []Profile		`json:"profiles",bson:"profiles"`

}