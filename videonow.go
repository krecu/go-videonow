package videonow

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/krecu/go-cache"
	_ "github.com/davecgh/go-spew/spew"
	"encoding/json"
	"strconv"
	"time"
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
	Cache *cache.Redis
	_cache bool
}

func New(host string) (*VideoNow, error) {

	mlClient, err :=  sql.Open("mysql", host)

	if err != nil {
		return &VideoNow{}, err
	}

	return &VideoNow{
		Client: mlClient,
		Params: &ConnectParams{
			Host: host,
		},
	}, nil
}

// закрываем соединение с mysql
func (v *VideoNow) Close() {
	v.Client.Close()
}

// установка кеша
func (v *VideoNow) SetCache(provider *cache.Redis) {
	v.Cache = provider
	v._cache = true
}

// выбираем профиль
func (v *VideoNow) Profile(id string) (*Profile, error){

	var (
		Id, Title, CategoryId, ContainerTypeId, DeviceId, VpaidType, CashIn,
		PartnerPlanId, SiteId, IsActive, IsTest, IsBad, UseContent []byte
	)

	proto := &Profile{}


	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {

		data, err := v.Cache.Get("profile_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}

	}

	// если модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {

		// Prepare statement for reading data
		row := v.Client.QueryRow(`
		SELECT
			id, title, category_id, container_type_id, device_id, vpaid_type, cash_in, partner_plan_id,
			site_id, is_active, is_test, is_bad, category_id, use_content
		FROM profile WHERE id = ?
		`, id)

		// извлекаем
		err := row.Scan(&Id, &Title, &CategoryId, &ContainerTypeId, &DeviceId, &VpaidType, &CashIn, &PartnerPlanId,
			&SiteId, &IsActive, &IsTest, &IsBad, &CategoryId, &UseContent); if err != nil {
			log.Printf("VIDEONOW MYSQL: запрос профиля %s: %s", id, err.Error())
			return &Profile{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id

		_PartnerPlanId, _ := strconv.ParseUint(string(PartnerPlanId), 10, 64)
		proto.PartnerPlanId = _PartnerPlanId

		proto.Title = string(Title)

		if string(IsActive) == "1" {
			proto.IsActive = true
		} else {
			proto.IsActive = false
		}

		if string(IsTest) == "1" {
			proto.IsTest = true
		} else {
			proto.IsTest = false
		}

		if string(IsBad) == "1" {
			proto.IsBad = true
		} else {
			proto.IsBad = false
		}

		if string(UseContent) == "" {
			proto.UseContent = false
		} else {
			proto.UseContent = true
		}

		if CategoryId != nil {
			category, err := v.Category(string(CategoryId)); if err != nil {
				log.Printf("VIDEONOW MYSQL: запрос категории %s: %s", CategoryId, err.Error())
			}
			proto.Category = *category
		}


		financial, err := v.ProfileFinancial(id); if err != nil {
			log.Printf("VIDEONOW MYSQL: запрос финансирования %s: %s", id, err.Error())
		}
		proto.Financial = financial

		financialContent, err := v.ProfileFinancialContent(id); if err != nil {
			log.Printf("VIDEONOW MYSQL: запрос финансировани контента %s: %s", id, err.Error())
		}
		proto.FinancialContent = financialContent


		// если задан сайт то извлечем еще и его
		if SiteId != nil {
			site, err := v.Site(string(SiteId)); if err != nil {
				log.Fatalf("VIDEONOW MYSQL: запрос сата %s: %s", SiteId, err.Error())
			}
			proto.Site = *site
		}

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto);
			if err == nil {
				v.Cache.Set("profile_"+id, string(data), time.Minute * 1)
			}
		}

	}

	return proto, nil
}

// выбираем компанию
func (v *VideoNow) Campaign(id string) (*Campaign, error){

	var (
		err error
		rows *sql.Rows

		Id, Name, Type, Status, Archived, Belong, IsApproved, IsPayment, IsApprovedText,
		AutoStart, Discount, OpenStatFlag, PlannedStartDate, PlannedStopDate,
		DateEnd, DayBudget, DayEvents, ShowsPerUser, Container, GeoType, CreativeType,
		Url, LinkTitle, LinkTarget, IsPriority, IsWordsTargeted, WordsTargeted, BasePrice,
		TimeTargeting, TargetUrlType, TargetListUrl, DeviceTargeted, Cpm, IsServerRequest,
		AdServer, ServerUrl, AdFoxKey, OwnerId, Created, ExtPriority, userId []byte
	)

	proto := &Campaign{}

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {

		data, err := v.Cache.Get("campaign_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {

		// берем основные данные по компании
		row := v.Client.QueryRow(`

		SELECT
			id, name, type, status, archived, belong, is_approved, is_payment, is_approved_text,
			auto_start, discount, openstat_flag, planned_start_date, planned_stop_date,
			date_end, day_budget, day_events, shows_per_user, container, geo_type, creative_type,
			url, link_title, link_target, is_priority, is_words_targeted, words_targeted, base_price,
			time_targeting, target_url_type, target_list_url, device_targeted, cpm, is_server_request,
			ad_server, server_url, adfox_key, owner_id, created, ext_priority,
			user_id
		FROM campaign WHERE id = ?
		`, id)

		err = row.Scan(
			&Id, &Name, &Type, &Status, &Archived, &Belong, &IsApproved, &IsPayment, &IsApprovedText,
			&AutoStart, &Discount, &OpenStatFlag, &PlannedStartDate, &PlannedStopDate,
			&DateEnd, &DayBudget, &DayEvents, &ShowsPerUser, &Container, &GeoType, &CreativeType,
			&Url, &LinkTitle, &LinkTarget, &IsPriority, &IsWordsTargeted, &WordsTargeted, &BasePrice,
			&TimeTargeting, &TargetUrlType, &TargetListUrl, &DeviceTargeted, &Cpm, &IsServerRequest,
			&AdServer, &ServerUrl, &AdFoxKey, &OwnerId, &Created, &ExtPriority,
			&userId,
		)

		if err != nil {
			log.Printf("Error load campaign %s: %s", id, err)
			return &Campaign{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id
		proto.Name = string(Name)
		proto.Type = string(Type)
		proto.Status = string(Status)
		if string(Archived) == "on" {
			proto.Archived = true
		} else {
			proto.Archived = false
		}
		proto.Belong = string(Belong)
		if string(IsApproved) == "ok" {
			proto.IsApproved = true
		} else {
			proto.IsApproved = false
		}
		if string(IsPayment) == "yes" {
			proto.IsPayment = true
		} else {
			proto.IsPayment = false
		}

		proto.IsApprovedText = string(IsApprovedText)

		if string(AutoStart) == "on" {
			proto.AutoStart = true
		} else {
			proto.AutoStart = false
		}

		_discount, _ := strconv.ParseFloat(string(Discount),64)
		proto.Discount = _discount

		if string(OpenStatFlag) == "on" {
			proto.OpenStatFlag = true
		} else {
			proto.OpenStatFlag = false
		}

		if PlannedStartDate != nil {
			_PlannedStartDate, _ := time.Parse("2006-01-02", string(PlannedStartDate))
			proto.PlannedStopDate = _PlannedStartDate
		}
		if PlannedStopDate != nil {
			_PlannedStopDate, _ := time.Parse("2006-01-02", string(PlannedStopDate))
			proto.PlannedStopDate = _PlannedStopDate
		}
		if DateEnd != nil {
			_DateEnd, _ := time.Parse("2006-01-02", string(DateEnd))
			proto.DateEnd = _DateEnd
		}

		_DayBudget, _ := strconv.ParseFloat(string(DayBudget),64)
		proto.DayBudget = _DayBudget

		_DayEvents, _ := strconv.ParseUint(string(DayEvents), 10, 64)
		proto.DayEvents = _DayEvents

		_ShowsPerUser, _ := strconv.ParseUint(string(ShowsPerUser), 10, 64)
		proto.ShowsPerUser = _ShowsPerUser

		proto.Container = string(Container)

		proto.GeoType, _ = strconv.Atoi(string(GeoType))

		proto.CreativeType = string(CreativeType)
		proto.Url = string(Url)
		proto.LinkTitle = string(LinkTitle)
		proto.LinkTarget = string(LinkTarget)

		if string(IsPriority) == "on" {
			proto.IsPriority = true
		} else {
			proto.IsPriority = false
		}
		if string(IsWordsTargeted) == "on" {
			proto.IsWordsTargeted = true
		} else {
			proto.IsWordsTargeted = false
		}

		proto.WordsTargeted = string(WordsTargeted)

		_BasePrice, _ := strconv.ParseFloat(string(BasePrice),64)
		proto.BasePrice = _BasePrice

		if string(TimeTargeting) == "1" {
			proto.TimeTargeting = true
		} else {
			proto.TimeTargeting = false
		}

		proto.TargetUrlType = string(TargetUrlType)
		proto.TargetListUrl = string(TargetListUrl)

		if string(DeviceTargeted) == "1" {
			proto.DeviceTargeted = true
		} else {
			proto.DeviceTargeted = false
		}

		_Cpm, _ := strconv.ParseFloat(string(Cpm),64)
		proto.Cpm = _Cpm

		if Created != nil {
			_Created, _ := time.Parse("2006-01-02 15:04:05", string(Created))
			proto.Created = _Created
		}


		// дастаем данные по юзеру
		if userId != nil {
			user, _ := v.User(string(userId))
			proto.User = *user
		}

		// дастаем доступные только для этой компании профили
		rows, err = v.Client.Query(`
		SELECT
			p.id
		FROM profile p
		WHERE p.is_active = 1 AND p.is_bad != 1 AND p.is_test != 1 AND
			p.id NOT IN (SELECT
					profile_id
				     FROM ban_profile_campaign
				     LEFT JOIN campaign c ON c.id = campaign_id
				     WHERE c.id = ?
				     )
  		`, id)

		if err == nil {
			var profileId int
			for rows.Next() {
				err = rows.Scan(&profileId); if err == nil {
					proto.Profiles = append(proto.Profiles, profileId)
				}
			}
		}

		// дастаем все категории этой компании
		rows, err = v.Client.Query(`
		SELECT
			cc.category_id
		FROM campaign_category cc
		WHERE cc.campaign_id = ?
  		`, id)

		if err == nil {
			var ccId string
			for rows.Next() {
				err = rows.Scan(&ccId); if err == nil {
					cc, err := v.CampaignCategory(ccId, id); if err == nil {
						proto.Category = append(proto.Category, *cc)
					}
				}
			}
		}

		// извлекаем регионы компании
		rows, err = v.Client.Query(`
		SELECT
			region_id
		FROM campaign_regions
		WHERE campaign_id = ? GROUP BY region_id
		`, id)

		if err == nil {
			protoR := &Region{}
			for rows.Next() {
				err = rows.Scan(&protoR.Id); if err == nil {
					proto.Region = append(proto.Region, *protoR)
				}
			}
		}

		// извлекаем страны компании
		rows, err = v.Client.Query(`
		SELECT
			country_id
		FROM campaign_countries
		WHERE campaign_id = ?
		`, id)

		if err == nil {
			protoС := &Country{}
			for rows.Next() {
				err = rows.Scan(&protoС.Id); if err == nil {
					proto.Countries = append(proto.Countries, *protoС)
				}
			}
		}

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("campaign_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем сайт
func (v *VideoNow) Site(id string) (*Site, error){

	proto := &Site{}

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("site_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {

		var (
			Id, IsActive, CategoryId, UserId []byte
		)

		// Prepare statement for reading data
		row := v.Client.QueryRow("SELECT id, user_id, is_active, category_id  FROM site WHERE id = ?", id)
		err := row.Scan(&Id, &UserId, &IsActive, &CategoryId); if err != nil {
			log.Println(err)
			return &Site{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id

		if UserId != nil {
			user, err := v.User(string(UserId)); if err != nil {
				log.Printf("MYSQL: запрос пользователя %s: %s", UserId, err.Error())
			}
			proto.User = *user
		}

		if CategoryId != nil {
			category, err := v.Category(string(CategoryId)); if err != nil {
				log.Printf("MYSQL: запрос категории %s: %s", CategoryId, err.Error())
			}
			proto.Category = *category
		}

		if string(IsActive) == "1" {
			proto.IsActive = true
		} else {
			proto.IsActive = false
		}

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("site_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем пользователя
func (v *VideoNow) User(id string) (*User, error){

	proto := &User{}

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("user_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {

		var (
			Id, UserAdvType []byte
		)

		// Prepare statement for reading data
		row := v.Client.QueryRow(`
		SELECT
			id,
			user_adv_type
		FROM user WHERE id = ?`, id)

		err := row.Scan(&Id, &UserAdvType);
		if err != nil {
			return &User{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id

		proto.UserAdvType = string(UserAdvType)

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("user_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем категории сайтов
func (v *VideoNow) Category(id string) (*Category, error){

	proto := &Category{}

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("category_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {

		var (
			Id, Name, Rate []byte
		)

		// Prepare statement for reading data
		row := v.Client.QueryRow(`
		SELECT
			id,
			name,
			rate
		FROM category WHERE id = ?`, id)

		err := row.Scan(&Id, &Name, &Rate);
		if err != nil {
			return &Category{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id

		proto.Name = string(Name)

		_Rate, _ := strconv.ParseFloat(string(Rate),64)
		proto.Rate = _Rate

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("category_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем категории компаний
func (v *VideoNow) CampaignCategory(id string, cid string) (*CampaignCategory, error){

	proto := &CampaignCategory{}

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("campaign_category_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if proto.Id == 0 {
		var (
			Id []byte
			PayEvent []byte
			PayEventCreated []byte
			OwnerEvent []byte
			OwnerEventCreated []byte
			EventPrice []byte
			EventPriceReal []byte
		)

		// извлекаем категории компании
		row := v.Client.QueryRow(`
		SELECT
		category_id, pay_events, pay_events_created, owner_events, owner_events_created,
		event_price, event_price_real FROM campaign_category
		WHERE category_id = ? AND campaign_id = ? LIMIT 1`, id, cid)

		err := row.Scan(&Id, &PayEvent, &PayEventCreated, &OwnerEvent, &OwnerEventCreated, &EventPrice, &EventPriceReal)

		if err != nil {
			return &CampaignCategory{}, err
		}

		_id, _ := strconv.ParseUint(string(Id), 10, 64)
		proto.Id = _id

		_PayEvent, _ := strconv.ParseUint(string(PayEvent), 10, 64)
		proto.PayEvent = _PayEvent

		_PayEventCreated, _ := strconv.ParseUint(string(PayEventCreated), 10, 64)
		proto.PayEventCreated = _PayEventCreated

		_OwnerEvent, _ := strconv.ParseUint(string(OwnerEvent), 10, 64)
		proto.OwnerEvent = _OwnerEvent

		_OwnerEventCreated, _ := strconv.ParseUint(string(OwnerEventCreated), 10, 64)
		proto.OwnerEventCreated = _OwnerEventCreated

		_EventPrice, _ := strconv.ParseFloat(string(EventPrice),64)
		proto.EventPrice = _EventPrice

		_EventPriceReal, _ := strconv.ParseFloat(string(EventPriceReal),64)
		proto.EventPriceReal = _EventPriceReal

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("campaign_category_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем финонсирование по профилю
func (v *VideoNow) ProfileFinancial(id string) ([]ProfileFinancial, error){

	var proto []ProfileFinancial
	var rows *sql.Rows
	var err error

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("profile_financial_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if len(proto) == 0 {

		var (
			Id, PartnerPercent, Belong, Percent,
			MinCpm, IsShow, CurrencyId, CountryId, UserId []byte
		)

		// дастаем финансовые условия
		rows, err = v.Client.Query(`
		SELECT
			IFNULL(ff.id, fs.id) AS id,
			IFNULL(pp.percent,0) AS partner_percent,
			IFNULL(ff.belong, fs.belong) AS belong,
			IFNULL(ff.percent, fs.percent) AS percent,
			IFNULL(ff.min_cpm, fs.min_cpm) AS min_cpm,
			IFNULL(ff.is_show, fs.is_show) AS is_show,
			IFNULL(ff.currency_id, fs.currency_id) AS currency_id,
			IFNULL(ff.country_id, fs.country_id) AS country_id,
			IFNULL(ff.user_id, fs.user_id) AS user_id
		FROM profile p
		LEFT JOIN partner_plan 	pp ON pp.id = p.partner_plan_id
		LEFT JOIN financial_condition ff ON ff.profile_id = p.id
		LEFT JOIN (
				SELECT
					f.*,
					pp.id AS profile
				FROM financial_condition f
				JOIN site s 		ON s.user_id 	=  f.user_id
				JOIN profile pp ON pp.site_id = s.id
				WHERE f.profile_id IS NULL
		) fs ON fs.profile = p.id

		WHERE p.id = ? GROUP BY belong
  		`, id);

		if err != nil {
			return proto, err
		}

		for rows.Next() {

			item := &ProfileFinancial{}

			err = rows.Scan(&Id, &PartnerPercent, &Belong, &Percent,
				&MinCpm, &IsShow, &CurrencyId, &CountryId, &UserId); if err == nil {

				_id, _ := strconv.ParseUint(string(Id), 10, 64)
				item.Id = _id

				_PartnerPercent, _ := strconv.ParseFloat(string(PartnerPercent),64)
				item.PartnerPercent = _PartnerPercent

				item.Belong = string(Belong)

				_Percent, _ := strconv.ParseFloat(string(Percent),64)
				item.Percent = _Percent

				_MinCpm, _ := strconv.ParseFloat(string(MinCpm),64)
				item.MinCpm = _MinCpm

				if string(IsShow) == "0" {
					item.IsShow = false
				} else {
					item.IsShow = true
				}

				_CurrencyId, _ := strconv.ParseUint(string(CurrencyId), 10, 64)
				item.CurrencyId = _CurrencyId

				_CountryId, _ := strconv.ParseUint(string(CountryId), 10, 64)
				item.CountryId = _CountryId

				_UserId, _ := strconv.ParseUint(string(UserId), 10, 64)
				item.UserId = _UserId


				proto = append(proto, *item)
			}


		}


		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("profile_financial_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}

// выбираем финонсирование по профилю
func (v *VideoNow) ProfileFinancialContent(id string) ([]ProfileFinancialContent, error){

	var proto []ProfileFinancialContent
	var rows *sql.Rows
	var err error

	// если используеться кеш
	// попробуем извлечь данные
	if v._cache {
		data, err := v.Cache.Get("profile_financial_content_" + id)
		if data != nil && err == nil {
			json.Unmarshal(data, &proto);
		}
	}

	// если в кеше модель пустая то попробуем извлечь из БД
	if len(proto) == 0 {

		var (
			Id, Percent, MinCpm, CurrencyId, UserId, ContentUserId []byte
		)

		// дастаем финансовые условия
		rows, err = v.Client.Query(`
		SELECT
			IFNULL(ff.id, fs.id) AS id,
			IFNULL(ff.percent, fs.percent) AS percent,
			IFNULL(ff.min_cpm, fs.min_cpm) AS min_cpm,
			IFNULL(ff.currency_id, fs.currency_id) AS currency_id,
			IFNULL(ff.user_id, fs.user_id) AS user_id,
			IFNULL(ff.content_user_id, fs.content_user_id) AS content_user_id
		FROM profile p
		LEFT JOIN partner_plan pp ON pp.id = p.partner_plan_id
		LEFT JOIN financial_condition_content_user ff ON ff.profile_id = p.id
		LEFT JOIN (
			SELECT
				f.*,
				pp.id AS profile
			FROM financial_condition_content_user f
			JOIN site s ON s.user_id = f.user_id
			JOIN profile pp ON pp.site_id = s.id
			WHERE f.profile_id IS NULL
		) fs ON fs.profile = p.id

		WHERE p.id = ?
  		`, id);

		if err != nil {
			return proto, err
		}

		for rows.Next() {

			item := &ProfileFinancialContent{}

			err = rows.Scan(&Id, &Percent, &MinCpm, &CurrencyId, &UserId, &ContentUserId); if err == nil {

				_id, _ := strconv.ParseUint(string(Id), 10, 64)
				item.Id = _id

				_Percent, _ := strconv.ParseFloat(string(Percent),64)
				item.Percent = _Percent

				_MinCpm, _ := strconv.ParseFloat(string(MinCpm),64)
				item.MinCpm = _MinCpm

				_CurrencyId, _ := strconv.ParseUint(string(CurrencyId), 10, 64)
				item.CurrencyId = _CurrencyId

				_ContentUserId, _ := strconv.ParseUint(string(ContentUserId), 10, 64)
				item.ContentUserId = _ContentUserId

				_UserId, _ := strconv.ParseUint(string(UserId), 10, 64)
				item.UserId = _UserId

				proto = append(proto, *item)
			}
		}

		// если есть кеш то положим обьект в него
		if v._cache {
			data, err := json.Marshal(proto); if err == nil {
				v.Cache.Set("profile_financial_content_"+id, string(data), time.Minute * 1)
			}
		}
	}

	return proto, nil
}