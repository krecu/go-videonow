package videonow

type Js struct {

	Version string			`json:"version",bson:"version"`//VARCHAR(80) DEFAULT '' NOT NULL,
	Data map[string]string		`json:"data",bson:"data"`//VARCHAR(30) NOT NULL,
}