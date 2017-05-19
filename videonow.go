package videonow

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type ConnectParams struct{
	Host string
	Port string
	User string
	Pass string
	Db string
}

type VideoNow struct {
	Params *ConnectParams
	Client *sql.DB
}

func New(host string) (*VideoNow) {

	mlClient, err :=  sql.Open("mysql", host)

	if err != nil {
		log.Fatalln("VideoNow connection error")
	}

	return &VideoNow{
		Client: mlClient,
		Params: &ConnectParams{
			Host: host,
		},
	}
}

// закрываем соединение с mysql
func (v *VideoNow) Close() {
	v.Client.Close()
}

// выбираем профиль
func (v *VideoNow) Profile(id string) (*Profile, error){

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, site_id, is_active, is_test, is_bad, category_id FROM profile WHERE id = ?", id)
	proto := &Profile{}

	var siteId string = ""
	err := row.Scan(&proto.Id, &siteId, &proto.Active, &proto.Test, &proto.Bad, &proto.Category)
	if err != nil {
		return &Profile{}, err
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
func (v *VideoNow) Campaign(id string) (*Campaign, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, belong, status, is_approved, discount  FROM campaign WHERE id = ?", id)
	proto := &Campaign{}
	err := row.Scan(&proto.Id, &proto.Belong, &proto.Status, &proto.Approved, &proto.Discount)

	if err != nil {
		return &Campaign{}, err
	}

	return proto, nil
}

// выбираем сайт
func (v *VideoNow) Site(id string) (*Site, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id, user_id, is_active, category_id  FROM site WHERE id = ?", id)
	proto := &Site{}
	userId := ""
	err := row.Scan(&proto.Id, &userId, &proto.Active, &proto.Category)

	if err != nil {
		return &Site{}, err
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
func (v *VideoNow) User(id string) (*User, error){

	//v.Ping()

	// Prepare statement for reading data
	row := v.Client.QueryRow("SELECT id FROM user WHERE id = ?", id)
	proto := &User{}
	err := row.Scan(&proto.Id)

	if err != nil {
		return &User{}, err
	}

	return proto, nil
}