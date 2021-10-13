package cp_go

const (
	ReferToCardIssuer = 5001 // Отказ эмитента проводить онлайн операцию

	// DoNotHonor Отказ эмитента без объяснения причин
	// Данный код возвращается по следующим причинам:
	// — неверно указан код CVV на картах MasterCard;
	// — внутренние ограничения банка, выпустившего карту;
	// — карта заблокирована или еще не активирована;
	// — на карте не включены интернет-платежи или не подключен 3DS.
	DoNotHonor = 5005

	Error                    = 5006 // Отказ сети проводить операцию или неправильный CVV код
	InvalidTransaction       = 5012 // Карта не предназначена для онлайн платежей
	AmountError              = 5013 // Слишком маленькая или слишком большая сумма операции
	FormatError              = 5030 // Ошибка на стороне эквайера — неверно сформирована транзакция
	BankNotSupportedBySwitch = 5031 // Неизвестный эмитент карты
	SuspectedFraud           = 5034 // Отказ эмитента — подозрение на мошенничество
	LostCard                 = 5041 // Карта потеряна
	StolenCard               = 5043 // Карта украдена
	InsufficientFunds        = 5051 // Недостаточно средств
	ExpiredCard              = 5054 // Карта просрочена или неверно указан срок действия

	// TransactionNotPermitted Ограничение на карте
	// Данный код возвращается по следующим причинам:
	//  — внутренние ограничения банка, выпустившего карту;
	//  — карта заблокирована или еще не активирована;
	//  — на карте не включены интернет-платежи или не подключен 3DS.
	TransactionNotPermitted = 5057

	ExceedWithdrawalFrequency = 5065 // Превышен лимит операций по карте
	IncorrectCVV              = 5082 // Неверный CVV код
	Timeout                   = 5091 // Эмитент недоступен
	CannotReachNetwork        = 5092 // Эмитент недоступен
	SystemError               = 5096 // Ошибка банка-эквайера или сети
	UnableToProcess           = 5204 // Операция не может быть обработана по прочим причинам
	AuthenticationFailed      = 5206 // 3-D Secure авторизация не пройдена
	AuthenticationUnavailable = 5207 // 3-D Secure авторизация недоступна
	AntiFraud                 = 5300 // Лимиты эквайера на проведение операций
)

var statusText = map[int]string{
	ReferToCardIssuer:         "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	DoNotHonor:                "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	Error:                     "Проверьте правильность введенных данных карты или воспользуйтесь другой картой",
	InvalidTransaction:        "Воспользуйтесь другой картой или свяжитесь с банком, выпустившим карту",
	AmountError:               "Проверьте корректность суммы",
	FormatError:               "Повторите попытку позже",
	BankNotSupportedBySwitch:  "Воспользуйтесь другой картой",
	SuspectedFraud:            "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	LostCard:                  "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	StolenCard:                "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	InsufficientFunds:         "Недостаточно средств на карте",
	ExpiredCard:               "Проверьте правильность введенных данных карты или воспользуйтесь другой картой",
	TransactionNotPermitted:   "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	ExceedWithdrawalFrequency: "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	IncorrectCVV:              "Неверно указан код CVV",
	Timeout:                   "Повторите попытку позже или воспользуйтесь другой картой",
	CannotReachNetwork:        "Повторите попытку позже или воспользуйтесь другой картой",
	SystemError:               "Повторите попытку позже",
	UnableToProcess:           "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	AuthenticationFailed:      "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	AuthenticationUnavailable: "Свяжитесь с вашим банком или воспользуйтесь другой картой",
	AntiFraud:                 "Воспользуйтесь другой картой",
}

// ErrorStatusText returns a text of error codes that determine the reason for the refusal to make a payment
// string if the code is unknown.
func ErrorStatusText(code int) string {
	return statusText[code]
}

const (
	SUCCESS            = 0
	UNKNOWN_INVOICE_ID = 10
	INVALID_ACCOUNT_ID = 11
	INVALID_AMOUNT     = 12
	REJECTED           = 13
	EXPIRED            = 20
)

const (
	KZT = "KZT"
	RUB = "RUB"
	EUR = "EUR"
	USD = "USD"
	GBP = "GBP"
	UAH = "UAH"
	BYR = "BYR"
	BYN = "BYN"
	AZN = "AZN"
	CHF = "CHF"
	CZK = "CZK"
	CAD = "CAD"
	PLN = "PLN"
	SEK = "SEK"
	TRY = "TRY"
	CNY = "CNY"
	INR = "INR"
)
