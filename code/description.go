package code

var ErrorMap = map[uint32]string{
	// !! Общие ошибки продукта !!
	ResultOK:               "Успешное выполнение операции",
	RequestProcessingError: "Внутренняя ошибка системы",
	ErrExecChanelDown:      "Канал исполнения недоступен. Попробуйте позже",
	// !! Ошибки Transaction-Manager
	// db_mercury (register ticket for bonus)
	ErrDataInputError: "Ошибка при вводе номера билета или кода выигрыша",
	// finance-manager (register ticket for bonus)
	ErrTicketHasBeenResold:        "Начисление бонусов невозможно, билет был продан другому участнику программы лояльности",
	ErrTicketWasBoughtWithBonuses: "Начисление бонусов невозможно, билет куплен за бонусы",
	ErrMissingCardOrPhoneNumber:   "Начисление бонусов невозможно, билет куплен без использования карты или номера телефона участника программы лояльности",
	// !! Ошибки Finance-Manager
	// Общие ошибки модуля
	ErrSapProcessingError:  "SAP недоступен",
	ErrInvalidJSON:         "Некорректное сообщение в ответе. Требуется JSON",
	ErrMethodNotFound:      "Запрашиваемый ресурс не найден",
	ErrUnexpectedErrorCode: "Неизвестный код ошибки бонусного сервера",
	// Calculate_price request
	ErrCalcPriceProductNotFound: "Продукт не найден",
	ErrCalcPriceCommonError:     "При обработке запроса возникла ошибка",
	// Consume points request
	ErrConsumeParticipationNotFound:            "Участие не найдено",
	ErrConsumeInvalidPin:                       "Неверный PIN",
	ErrConsumeProductNotFound:                  "Продукт не найден",
	ErrConsumeParticipationBlocked:             "Участие заблокировано",
	ErrConsumeBonusPayRestricted:               "Списание ограничено",
	ErrConsumeNotEnoughPoints:                  "Недостаточно бонусных баллов",
	ErrConsumeBonusPriceChanged:                "Стоимость продукта в бонусах изменена",
	ErrConsumePinLimit:                         "Превышено количество попыток ввода PIN-кода",
	ErrConsumePinExpired:                       "Срок действия PIN-кода истек",
	ErrConsumeProductIDCourseMissing:           "Для продукта отсутствует курс",
	ErrConsumeAuthorizedPurchaseAmountExceeded: "Превышена авторизованная сумма покупки",
	ErrConsumeCommonError:                      "При обработке запроса возникла ошибка",
	// Return points request
	ErrReturnParticipationNotFound: "Участие не найдено",
	ErrReturnParticipationBlocked:  "Участие заблокировано",
	ErrReturnOperationNotFound:     "Отменяемая операция не найдена",
	ErrReturnCommonError:           "При обработке запроса возникла ошибка",
	// Draw result request
	ErrDrawResultParticipationNotFound:           "Участие не найдено",
	ErrDrawResultTicketAlreadyRegistered:         "Билет уже был зарегистрирован в акции",
	ErrDrawResultProductNotFound:                 "Продукт не найден",
	ErrDrawResultParticipationBlocked:            "Участие заблокировано",
	ErrDrawResultOperationNotAvailableOnTerminal: "Операция недоступна на данном терминале",
	ErrDrawResultCommonError:                     "При обработке запроса возникла ошибка",
}

/*
   {ResultCode::PosProviderRequired,
    "Запрос может быть исполнен только через POS провайдер."},
   {ResultCode::BonusOperationsNotAllowedForProvider,
    "Бонусные операции недоступны для текущего провайдера"},
   {ResultCode::OrtaxPersonNotAuthorized,
    "Участник не авторизован в системе Ortax"},
   {ResultCode::OrtaxTransportError,
    "Невозможно отправить запрос в систему ORTAX"},
   {ResultCode::CorrectionsGateTaskAlreadyExist,
    "Задача с указанным uuid уже существует"},
   {ResultCode::CorrectionsGateTaskNotFound, "Задача не найдена"},
   {ResultCode::CorrectionsGateTaskTicketsNotFound,
    "Данные билетов для задачи с указанным uuid не найдены"},
*/

func GetDescription(sign uint32) string {
	description, ok := ErrorMap[sign]
	if !ok {
		return "Описание для данной ошибки не подготовлено. Обратитесь в службу технической поддержки"
	}
	return description
}
