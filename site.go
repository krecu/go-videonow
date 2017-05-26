package videonow

//CREATE TABLE site
//(
//id INT(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
//user_id INT(11) unsigned DEFAULT '0' NOT NULL,
//is_active TINYINT(1) DEFAULT '1' NOT NULL,
//reject_reason TEXT NOT NULL,
//title VARCHAR(255),
//description TEXT NOT NULL,
//url VARCHAR(64),
//xml_url VARCHAR(255) NOT NULL,
//update_interval TINYINT(4) DEFAULT '0' NOT NULL,
//created DATETIME NOT NULL,
//default_profile_partner_plan_id INT(10) unsigned NOT NULL,
//category_id INT(11),
//connection_type_id TINYINT(4) NOT NULL,
//google_slotname VARCHAR(255) NOT NULL COMMENT 'Параметр для рекламной сети WMGroup, у них свои id для наших сайтов',
//yandex_id VARCHAR(255) NOT NULL COMMENT 'Параметр для Yandex, у них свои id для наших сайтов',
//comment TEXT NOT NULL
//);
//CREATE INDEX user_id ON site (user_id);

type Site struct {
	Id uint64		// id INT(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	IsActive bool		//is_active TINYINT(1) DEFAULT '1' NOT NULL,
	User User		//user_id INT(11) unsigned DEFAULT '0' NOT NULL,
	Category Category	//category_id INT(11),
}