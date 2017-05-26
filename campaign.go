package videonow

import "time"
type Campaign struct {

	Id uint64			//INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
	Name string			//VARCHAR(80) DEFAULT '' NOT NULL,
	Type string			//VARCHAR(30) NOT NULL,
	Status string 			//ENUM('no_data', 'stop', 'work', 'end') DEFAULT 'no_data' NOT NULL,
	Archived bool 			//ENUM('', 'on') NOT NULL,
	Belong string 			//ENUM('owner', 'adv', 'external', 'service', 'owner_adv', 'owner_external', 'rtb'),
	IsApproved bool 		//ENUM('null', 'new', 'ok', 'no') DEFAULT 'null' NOT NULL,
	IsPayment bool 			//ENUM('no', 'yes') DEFAULT 'no' NOT NULL,
	IsApprovedText string 		//TEXT NOT NULL,
	AutoStart bool 			//ENUM('on', 'off') DEFAULT 'on',
	Discount float64 		//FLOAT UNSIGNED DEFAULT '0' NOT NULL,
	OpenStatFlag bool 		//ENUM('on', 'off') DEFAULT 'off' NOT NULL,
	PlannedStartDate time.Time 	//DATE NOT NULL,
	PlannedStopDate  time.Time 	//DATE NOT NULL,
	DateEnd time.Time 		//DATE NOT NULL,
	DayBudget float64 		//FLOAT DEFAULT '0' NOT NULL COMMENT 'Желаемый бюджет на один день',
	DayEvents uint64 		//INT(11) NOT NULL COMMENT 'Желаемое количество событий в день',
	ShowsPerUser uint64 		//TINYINT(3) unsigned DEFAULT '3' NOT NULL,
	Container string 		//VARCHAR(30) NOT NULL,
	GeoType int 			//TINYINT(3) unsigned DEFAULT '0' NOT NULL COMMENT 'Тип гео-таргетинга',
	CreativeType string 		//ENUM('inner', 'wrapper') NOT NULL,
	Url string 			//TEXT NOT NULL,
	LinkTitle string 		//VARCHAR(255) NOT NULL,
	LinkTarget string 		//ENUM('', '_self') NOT NULL,
	IsPriority bool 		//ENUM('', 'on') NOT NULL,
	IsWordsTargeted bool 		//ENUM('', 'on') NOT NULL,
	WordsTargeted string 		//TEXT NOT NULL,
	BasePrice float64 		//DECIMAL(10,4) COMMENT 'цена события до наценок за таргетинг',
	TimeTargeting bool 		//TINYINT(1) unsigned DEFAULT '0' NOT NULL,
	TargetUrlType string 		//ENUM('target', 'ban'),
	TargetListUrl string 		//TEXT NOT NULL,
	DeviceTargeted bool 		//TINYINT(1) unsigned DEFAULT '0' NOT NULL,
	Cpm float64 			//DECIMAL(11,2) DEFAULT '0.00' NOT NULL COMMENT 'Cpm для external-кампаний',
	IsServerRequest bool 		//TINYINT(1) unsigned DEFAULT '0' NOT NULL COMMENT 'Поддерживает ли внешняя кампания серверные запросы рекламы', // todo - пока ненужен
	AdServer int64 			//TINYINT(4) unsigned, // todo - пока ненужен
	ServerUrl string 		//VARCHAR(500), // todo - пока ненужен
	AdFoxKey string 		//VARCHAR(50) DEFAULT '' NOT NULL, // todo - пока ненужен
	OwnerId int64 			//INT(10) unsigned, // todo - пока ненужен
	Created time.Time 		//TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ExtPriority int64 		//nINT(10) unsigned DEFAULT '0',

	User User			// refference by user_id
	Category []CampaignCategory	//
	Region	 []Region
	Countries []Country
	Profiles []Profile

}