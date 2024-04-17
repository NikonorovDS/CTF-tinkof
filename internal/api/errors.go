package api

import "errors"

var (
	errDBSomethingWrong = errors.New("ошибка БД")
	errIdNotFound       = errors.New("ID не найден")

	errWrongParamType = errors.New("неверный тип параметра")
	errWrongJsonData  = errors.New("неверный формат json данных")

	errNotEnoughMoney    = errors.New("недостаточно денег")
	errTicketAlreadyUsed = errors.New("билет уже съеден")

	errPermissionDenied = errors.New("доступ запрещен")

	errEmptyUsernameOrPassword    = errors.New("имя пользователя или пароль не могут быть пустыми")
	errTooShortUsernameOrPassword = errors.New("имя пользователя и пароль должны быть не меньше 9 символов")
	errUsernameAlreadyExists      = errors.New("имя пользователя уже занято")
	errInvalidLoginCreds          = errors.New("неверные данные для входа")
	errFailedClearSession         = errors.New("не удалось отчистить сессию")

	errInterviewFailed = errors.New("вы не прошли интервью")
)
