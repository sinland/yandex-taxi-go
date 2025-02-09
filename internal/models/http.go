package models

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// CarsListQueryParkCar Поисковые ограничения на автомобиль
type CarsListQueryParkCar struct {
	Amenities  []string `json:"amenities,omitempty"`  // Удобства в ТС
	Categories []string `json:"categories,omitempty"` // Список категорий ТС
	Id         []string `json:"id,omitempty"`
	Status     []string `query:"status,omitempty"`
	IsRental   bool     `query:"is_rental,omitempty"`
}

// CarsListQueryPark ...
type CarsListQueryPark struct {
	Id  string                `json:"id"`            // Идентификатор партнёра
	Car *CarsListQueryParkCar `json:"car,omitempty"` // Поисковые ограничения на автомобиль
}

// CarsListQuery Поисковые ограничения
type CarsListQuery struct {
	Park CarsListQueryPark `json:"park"`           //
	Text string            `json:"text,omitempty"` // Текстовый поисковый запрос по данным автомобиля
}

// CarsListFields Данные, которые необходимо извлечь
type CarsListFields struct {
	Car []string `json:"car"` // Данные ТС, которые необходимо извлечь
}

// Vehicle Данные ТС
type Vehicle struct {
	Id               string   `json:"id"`                // Идентификатор ТС
	Amenities        []string `json:"amenities"`         // Удобства в ТС
	Brand            string   `json:"brand"`             // Марка ТС
	Callsign         string   `json:"callsign"`          // Позывной
	Category         []string `json:"category"`          // Список категорий ТС
	Color            string   `json:"color"`             // Цвет ТС
	Model            string   `json:"model"`             // Модель ТС
	Number           string   `json:"number"`            // Государственный регистрационный номер
	RegistrationCert string   `json:"registration_cert"` // Номер свидетельства о регистрации ТС (Обязательное поле для России)
	Status           string   `json:"status"`            // Статус ТС
	Vin              string   `json:"vin"`               // VIN (Обязательное поле для России)
	Year             int      `json:"year"`              // Год выпуска ТС
}

type DriverProfileRequestSortOrderField struct {
	Direction string `json:"direction"` // Направление сортировки ('asc', 'desc')
	Field     string `json:"field"`     // Поле, по которому сортируются значения
}

// DriverProfileListRequestFields Поля профиля, которые необходимо извлечь. Если не указано, то возвращаются все
// поля профиля. Чтобы исключить определенный блок полей, передайте пустой массив для соответствующего раздела.
// Например, чтобы исключить информацию об автомобиле, укажите "car": []
type DriverProfileListRequestFields struct {
	Account       []string `json:"account,omitempty"` // Данные счёта, которые необходимо извлечь
	Car           []string `json:"car,omitempty"`     // Данные ТС, которые необходимо извлечь
	CurrentStatus []string `json:"current_status,omitempty"`
	DriverProfile []string `json:"driver_profile,omitempty"`
	Park          []string `json:"park,omitempty"`
	UpdatedAt     bool     `json:"updated_at"`
}

type DriverProfilesListRequestQueryParkAccountLastTransactionDate struct {
	From string `json:"from"` // Время от в формате ISO 8601
	To   string `json:"to"`   // Время от в формате ISO 8601
}

type DriverProfilesListRequestQueryParkAccount struct {
	LastTransactionDate *DriverProfilesListRequestQueryParkAccountLastTransactionDate `json:"last_transaction_date,omitempty"`
}

type DriverProfilesListRequestQueryParkCurrentStatus struct {
	Status []string `json:"status"`
}

type DriverProfilesListRequestQueryParkDriverProfile struct {
	Id         []string `json:"id,omitempty"`           // Идентификатор профиля водителя, ex: 2111ade6gk054dfdb9iu8c8cc9460mks
	WorkRuleID []string `json:"work_rule_id,omitempty"` // Идентификатор условия работы, ex: bc43tre6ba054dfdb7143ckfgvcby63e
	WorkStatus []string `json:"work_status,omitempty"`  // Статус работы водителя (working,not_working,fired)
}

type DriverProfilesListRequestQueryParkUpdatedAt struct {
	From string `json:"from"` // Время от в формате ISO 8601
	To   string `json:"to"`   // Время от в формате ISO 8601
}

type DriverProfilesListRequestQueryPark struct {
	Id            string                                           `json:"id"`                       // Идентификатор партнёра
	Account       *DriverProfilesListRequestQueryParkAccount       `json:"account,omitempty"`        // Фильтры по данным счёта
	CurrentStatus *DriverProfilesListRequestQueryParkCurrentStatus `json:"current_status,omitempty"` // Фильтр по текущему состоянию водителя
	DriverProfile *DriverProfilesListRequestQueryParkDriverProfile `json:"driver_profile,omitempty"` // Фильтры по данным водительского профиля
	UpdatedAt     *DriverProfilesListRequestQueryParkUpdatedAt     `json:"updated_at,omitempty"`     // Фильтры по времени последнего обновления; Полуинтервал времени, хотя бы один конец должен быть указан
}

type DriverProfilesListRequestQuery struct {
	Park *DriverProfilesListRequestQueryPark `json:"park,omitempty"` // Параметры партнера
	Text string                              `json:"text,omitempty"` // Произвольный текстовый поисковый запрос
}

type DriverProfilePark struct {
	Id   string `json:"id"`   // Идентификатор партнёра
	City string `json:"city"` // Город партнера
	Name string `json:"name"` // Название партнера
}

// DriverProfileAccount Информация о счете
type DriverProfileAccount struct {
	Id           string `json:"id"`            // Идентификатор счета
	Balance      string `json:"balance"`       // Текущий баланс (сумма с фиксированной точностью)
	BalanceLimit string `json:"balance_limit"` // Лимит по счету
	Currency     string `json:"currency"`      // Валюта в формате ISO 4217
	Type         string `json:"type"`          // Тип счета
}

type DriverProfileCurrentStatus struct {
	Status          string `json:"status"`            // Текущее состояние водителя
	StatusUpdatedAt string `json:"status_updated_at"` // Время последнего обновления текущего состояния водителя в формате ISO 8601.
}

type DriverLicense struct {
	IssueDate        string `yaml:"issue_date"`
	ExpirationDate   string `yaml:"expiration_date"`
	Number           string `yaml:"number"`
	NormalizedNumber string `yaml:"normalized_number"`
	Country          string `yaml:"country"`
	BirthDate        string `yaml:"birth_date"`
}

type DriverProfileModel struct {
	Id               string        `json:"id"`                 // Идентификатор профиля водителя
	CheckMessage     string        `json:"check_message"`      // Прочее (доступно сотрудникам парка)
	Comment          string        `json:"comment"`            // ...
	CreatedDate      string        `json:"created_date"`       // Дата создания профиля в формате ISO 8601
	DriverLicense    DriverLicense `json:"driver_license"`     // Водительское удостоверение
	EmploymentType   string        `json:"employment_type"`    // Тип занятости водителя
	FirstName        string        `json:"first_name"`         // Имя
	HasContractIssue bool          `json:"has_contract_issue"` // Существуют проблемы с подтверждением занятости
	LastName         string        `json:"last_name"`          // Фамилия
	MiddleName       string        `json:"middle_name"`        // Отчество
	ParkId           string        `json:"park_id"`            // Идентификатор партнёра
	Phones           []string      `json:"phones"`             // Номер телефона
	WorkRuleId       string        `json:"work_rule_id"`       // Идентификатор условия работы
	WorkStatus       string        `json:"work_status"`        // Статус работы водителя
}

type DriverProfile struct {
	Accounts      []DriverProfileAccount      `json:"accounts"`       // Список счетов, которые связаны с водителем. Информация о счете
	Car           *Vehicle                    `json:"car"`            // Данные ТС
	CurrentStatus *DriverProfileCurrentStatus `json:"current_status"` // ...
	DriverProfile *DriverProfileModel         `json:"driver_profile"` // Профиль водителя
}

// ------------

// CarsListRequest Запрос на получение списка автомобилей
type CarsListRequest struct {
	Limit  int             `json:"limit"`
	Offset int             `json:"offset"`
	Query  CarsListQuery   `json:"query"`
	Fields *CarsListFields `json:"fields,omitempty"`
}

// CarsListResponse ...
type CarsListResponse struct {
	Total  int       `json:"total"`  // Общее число автомобилей, удовлетворяющих запросу
	Offset int       `json:"offset"` // Отступ, начиная с которого возвращаются автомобили в ответе
	Limit  int       `json:"limit"`  // Ограничение сверху на число автомобилей в ответе
	Cars   []Vehicle `json:"cars"`   // Данные ТС
}

// DriverProfilesRequest ...
type DriverProfilesRequest struct {
	SortOrder []DriverProfileRequestSortOrderField `json:"sort_order,omitempty"` // Массив полей для управления порядком профилей в ответе
	Offset    int                                  `json:"offset"`               // Смещение относительно начала списка
	Limit     int                                  `json:"limit"`                // Запрашиваемое число элементов списка
	Fields    *DriverProfileListRequestFields      `json:"fields,omitempty"`     // Поля профиля, которые необходимо извлечь
	Query     DriverProfilesListRequestQuery       `json:"query"`                // Фильтры, объединяются через логическое "И"
}

// DriverProfilesResponse ...
type DriverProfilesResponse struct {
	Total          int                 `json:"total"`           // Общее количество элементов списка
	Offset         int                 `json:"offset"`          // Запрошённое смещение относительно начала списка
	Limit          int                 `json:"limit"`           // Запрошённое число элементов списка
	Parks          []DriverProfilePark `json:"parks"`           // Список партнеров
	DriverProfiles []DriverProfile     `json:"driver_profiles"` // Список профилей
}
