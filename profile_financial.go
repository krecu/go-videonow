package videonow

//SELECT
//IFNULL(ff.id, fs.id) AS id,
//p.use_content AS use_content,
//p.partner_plan_id AS partner_plan_id,
//IFNULL(pp.percent,0) AS partner_percent,
//IFNULL(ff.belong, fs.belong) AS belong,
//IFNULL(ff.percent, fs.percent) AS percent,
//IFNULL(ff.min_cpm, fs.min_cpm) AS min_cpm,
//IFNULL(ff.is_show, fs.is_show) AS is_show,
//IFNULL(ff.currency_id, fs.currency_id) AS currency_id,
//IFNULL(ff.country_id, fs.country_id) AS country_id,
//IFNULL(ff.user_id, fs.user_id) AS user_id
//FROM profile p
//LEFT JOIN partner_plan 	pp ON pp.id = p.partner_plan_id
//LEFT JOIN financial_condition ff ON ff.profile_id = p.id
//LEFT JOIN (
//SELECT
//f.*,
//pp.id AS profile
//FROM financial_condition f
//JOIN site s 		ON s.user_id 	=  f.user_id
//JOIN profile pp ON pp.site_id = s.id
//WHERE f.profile_id IS NULL
//) fs ON fs.profile = p.id
//
//WHERE p.id = ? GROUP BY belong

type ProfileFinancial struct {
	Id uint64		`json:"id",bson:"id"`//INT(11) DEFAULT '0' NOT NULL,
	PartnerPercent float64	`json:"partner_percent",bson:"partner_percent"`//INT(10) unsigned DEFAULT '0' NOT NULL,
	Belong string		`json:"belong",bson:"belong"`//INT(10) unsigned DEFAULT '0' NOT NULL,
	Percent float64		`json:"percent",bson:"percent"`//DECIMAL(10,4) unsigned NOT NULL,
	MinCpm float64		`json:"min_cpm",bson:"min_cpm"`//DECIMAL(10,4) unsigned NOT NULL,
	IsShow bool		`json:"is_show",bson:"is_show"`//DECIMAL(10,4) unsigned NOT NULL,
	CurrencyId uint64	`json:"currency_id",bson:"currency_id"`//DECIMAL(10,4) unsigned NOT NULL,
	CountryId uint64	`json:"country_id",bson:"country_id"`//DECIMAL(10,4) unsigned NOT NULL,
	UserId uint64		`json:"user_id",bson:"user_id"`//DECIMAL(10,4) unsigned NOT NULL,
}