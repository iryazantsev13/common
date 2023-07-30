package code

const (
	// !! Общие ошибки продукта !!
	ResultOK               uint32 = 0
	RequestProcessingError        = 46
	ErrExecChanelDown             = 193
	// !!------- Ошибки Transaction-Manager -------!!
	// db_mercury (register ticket for bonus)
	ErrDataInputError = 19007
	// finance-manager (register ticket for bonus)
	ErrTicketHasBeenResold        = 19004
	ErrTicketWasBoughtWithBonuses = 19005
	ErrMissingCardOrPhoneNumber   = 19006
	// !!------- Ошибки Finance-Manager -------!!
	// Общие ошибки модуля
	ErrSapProcessingError  = 18099
	ErrInvalidJSON         = 19000
	ErrMethodNotFound      = 19001
	ErrUnexpectedErrorCode = 19002
	// Calculate_price request
	ErrCalcPriceProductNotFound = 20001
	ErrCalcPriceCommonError     = 20099
	// Consume points request
	ErrConsumeParticipationNotFound            = 21001
	ErrConsumeInvalidPin                       = 21002
	ErrConsumeProductNotFound                  = 21003
	ErrConsumeParticipationBlocked             = 21004
	ErrConsumeBonusPayRestricted               = 21005
	ErrConsumeNotEnoughPoints                  = 21006
	ErrConsumeBonusPriceChanged                = 21007
	ErrConsumePinLimit                         = 21008
	ErrConsumePinExpired                       = 21009
	ErrConsumeProductIDCourseMissing           = 21010
	ErrConsumeAuthorizedPurchaseAmountExceeded = 21011
	ErrConsumeCommonError                      = 21099
	// Return points request
	ErrReturnParticipationNotFound = 22001
	ErrReturnParticipationBlocked  = 22002
	ErrReturnOperationNotFound     = 22003
	ErrReturnCommonError           = 22099
	// Draw result request
	ErrDrawResultParticipationNotFound           = 23001
	ErrDrawResultTicketAlreadyRegistered         = 23002
	ErrDrawResultProductNotFound                 = 23003
	ErrDrawResultParticipationBlocked            = 23004
	ErrDrawResultOperationNotAvailableOnTerminal = 23005
	ErrDrawResultCommonError                     = 23099
)
