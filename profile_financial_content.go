package videonow

//SELECT
//IFNULL(ff.id, fs.id) AS id
//IFNULL(ff.percent, fs.percent) AS percent,
//IFNULL(ff.min_cpm, fs.min_cpm) AS min_cpm,
//IFNULL(ff.currency_id, fs.currency_id) AS currency_id,
//IFNULL(ff.user_id, fs.user_id) AS user_id,
//IFNULL(ff.content_user_id, fs.content_user_id) AS content_user_id
//FROM profile p
//LEFT JOIN partner_plan pp ON pp.id 			  = p.partner_plan_id
//LEFT JOIN financial_condition_content_user ff ON ff.profile_id = p.id
//LEFT JOIN (
//SELECT
//f.*,
//pp.id AS profile
//FROM financial_condition_content_user f
//JOIN site s 		ON s.user_id 	=  f.user_id
//JOIN profile pp ON pp.site_id = s.id
//WHERE f.profile_id IS NULL
//) fs ON fs.profile = p.id
//
//WHERE p.id = 1813580;

type ProfileFinancialContent struct {
	Id uint64		`json:"id",bson:"id"`//INT(11) DEFAULT '0' NOT NULL,
	Percent float64		`json:"percent",bson:"percent"`//DECIMAL(10,4) unsigned NOT NULL,
	MinCpm float64		`json:"min_cpm",bson:"min_cpm"`//DECIMAL(10,4) unsigned NOT NULL,
	CurrencyId uint64	`json:"currency_id",bson:"currency_id"`//DECIMAL(10,4) unsigned NOT NULL,
	UserId uint64		`json:"user_id",bson:"user_id"`//DECIMAL(10,4) unsigned NOT NULL,
	ContentUserId uint64	`json:"content_user_id",bson:"content_user_id"`//DECIMAL(10,4) unsigned NOT NULL,
}