package videonow

//CREATE TABLE category
//(
//id SMALLINT(6) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
//name VARCHAR(50) DEFAULT '' NOT NULL,
//iab_category_id TINYINT(3) unsigned,
//iab_subcategory_id TINYINT(3) unsigned,
//imho_category_id INT(10) unsigned,
//rate FLOAT(5,2) unsigned DEFAULT '1.00'
//);

type Category struct {
	Id uint64		`json:"id",bson:"id"`//id SMALLINT(6) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	Name string		`json:"name",bson:"name"`//name VARCHAR(50) DEFAULT '' NOT NULL,
	Rate float64		`json:"rate",bson:"rate"`// rate FLOAT(5,2) unsigned DEFAULT '1.00'
}