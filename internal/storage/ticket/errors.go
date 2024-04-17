package ticket

import "errors"

var (
	errIdNotFound        = errors.New("ID не найден")
	errPermissionDenied  = errors.New("доступ запрещен")
	errTicketAlreadyUsed = errors.New("билет уже съеден")
	errNotEnoughMoney    = errors.New("недостаточно денег")
)
