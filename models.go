package yandex_taxi_go

// Vehicle Данные ТС
type Vehicle struct {
	Id               string   // Идентификатор ТС
	Amenities        []string // Удобства в ТС
	Brand            string   // Марка ТС
	Callsign         string   // Позывной
	Category         []string // Список категорий ТС
	Color            string   // Цвет ТС
	Model            string   // Модель ТС
	Number           string   // Государственный регистрационный номер
	RegistrationCert string   // Номер свидетельства о регистрации ТС (Обязательное поле для России)
	Status           string   // Статус ТС
	Vin              string   // VIN (Обязательное поле для России)
	Year             int      // Год выпуска ТС
}

type DriverProfileAccount struct {
	Id           string // Идентификатор счета
	Balance      string // Текущий баланс (сумма с фиксированной точностью)
	BalanceLimit string // Лимит по счету
	Currency     string // Валюта в формате ISO 4217
	Type         string // Тип счета
}

type DriverProfileCurrentStatus struct {
	Status          string // Текущее состояние водителя
	StatusUpdatedAt string // Время последнего обновления текущего состояния водителя в формате ISO 8601.
}

type DriverLicense struct {
	IssueDate        string
	ExpirationDate   string
	Number           string
	NormalizedNumber string
	Country          string
	BirthDate        string
}

type DriverProfileData struct {
	Id               string        // Идентификатор профиля водителя
	CheckMessage     string        // Прочее (доступно сотрудникам парка)
	Comment          string        // ...
	CreatedDate      string        // Дата создания профиля в формате ISO 8601
	DriverLicense    DriverLicense // Водительское удостоверение
	EmploymentType   string        // Тип занятости водителя
	FirstName        string        // Имя
	HasContractIssue bool          // Существуют проблемы с подтверждением занятости
	LastName         string        // Фамилия
	MiddleName       string        // Отчество
	ParkId           string        // Идентификатор партнёра
	Phones           []string      // Номер телефона
	WorkRuleId       string        // Идентификатор условия работы
	WorkStatus       string        // Статус работы водителя
}

type DriverProfilePark struct {
	Id   string // Идентификатор партнёра
	City string // Город партнера
	Name string // Название партнера
}

type DriverProfile struct {
	Accounts      []DriverProfileAccount      //Список счетов, которые связаны с водителем.
	Car           *Vehicle                    // Данные ТС
	CurrentStatus *DriverProfileCurrentStatus // ..
	Profile       *DriverProfileData          // Профиль водителя
}

type GetCarsListArgs struct {
	ParkID string
	Page   int
	Limit  int
}

type GetCarsListResult struct {
	Total  int       // Общее число автомобилей, удовлетворяющих запросу
	Offset int       // Отступ, начиная с которого возвращаются автомобили в ответе
	Limit  int       // Ограничение сверху на число автомобилей в ответе
	Cars   []Vehicle // Данные ТС
}

type GetDriverProfilesArgs struct {
	Offset    int
	Limit     int
	QueryText string
	ParkId    string
}

type GetDriverProfilesResult struct {
	Total          int                 // Общее число автомобилей, удовлетворяющих запросу
	Offset         int                 // Отступ, начиная с которого возвращаются автомобили в ответе
	Limit          int                 // Ограничение сверху на число автомобилей в ответе
	DriverProfiles []DriverProfile     // Список профилей
	Parks          []DriverProfilePark // Список партнеров
}
