package videonow

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"stats/videonow/model"
)


type ConnectParams struct{
	Host string
	Port string
	User string
	Pass string
	Db string
}

type VideoNowClient struct {
	Params *ConnectParams
	Client *sql.DB
	IsConnect int
}

func NewVideoNowClient(host string, port string, user string, pass string, db string) (*VideoNowClient) {

	mlClient, err :=  sql.Open("mysql", "root:r00-t@tcp(127.0.0.1:3306)/videonow")

	if err != nil {
		log.Fatalln("VideoNow connection error")
	}

	return &VideoNowClient{
		Client: mlClient,
		IsConnect: 1,
		Params: &ConnectParams{
			Host: host,
			Port: port,
			User: user,
			Pass: pass,
			Db: db,
		},
	}
}

// переподключаемся при разрывке связи
func (v *VideoNowClient) Ping() (error) {

	err := v.Client.Ping()

	if err != nil {

		mlClient, err := sql.Open("mysql", v.Params.User + ":" + v.Params.Pass + "@" + v.Params.Host + ":" + string(v.Params.Port) + "/" + v.Params.Db)

		if err != nil {
			log.Fatal(err)
			return err
		}

		v.Client = mlClient
	}

	return nil
}

// выбираем профиль
func (v *VideoNowClient) Profile(id string) (*model.Profile, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, site_id, is_active, is_test, is_bad, category_id FROM profile WHERE id = ?", id)
	proto := new(model.Profile)
	var siteId string = ""
	err := row.Scan(&proto.Id, &siteId, &proto.Active, &proto.Test, &proto.Bad, &proto.Category)
	if err != nil {
		return &model.Profile{}, err
	}

	if siteId != "" {
		site, err := v.Site(siteId)
		if err != nil {
			log.Fatalln("MYSQL: запрос сата %s: %s", siteId, err.Error())
		}
		proto.Site = *site
	}

	return proto, nil
}

// выбираем компанию
func (v *VideoNowClient) Campaign(id string) (*model.Campaign, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, belong, status, is_approved, discount  FROM campaign WHERE id = ?", id)
	proto := new(model.Campaign)
	err := row.Scan(&proto.Id, &proto.Belong, &proto.Status, &proto.Approved, &proto.Discount)

	if err != nil {
		return &model.Campaign{}, err
	}

	return proto, nil
}

// выбираем сайт
func (v *VideoNowClient) Site(id string) (*model.Site, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, user_id, is_active, category_id  FROM site WHERE id = ?", id)
	proto := new(model.Site)
	userId := ""
	err := row.Scan(&proto.Id, &userId, &proto.Active, &proto.Category)

	if err != nil {
		return &model.Site{}, err
	}

	if userId != "" {
		user, err := v.User(userId)
		if err != nil {
			log.Fatalln("MYSQL: запрос пользователя %s: %s", userId, err.Error())
		}
		proto.User = *user
	}

	return proto, nil
}

// выбираем пользователя
func (v *VideoNowClient) User(id string) (*model.User, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id FROM user WHERE id = ?", id)
	proto := new(model.User)
	err := row.Scan(&proto.Id)

	if err != nil {
		return &model.User{}, err
	}

	return proto, nil
}