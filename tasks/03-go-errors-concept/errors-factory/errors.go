package errors

// NewError возвращает новое значение-ошибку, текст которой является msg.
// Две ошибки с одинаковым текстом, созданные через NewError, не равны между собой:
//
//	NewError("end of file") != NewError("end of file")
type err struct {
	msg string
}

func (e *err) Error() string {
	return e.msg
}

func NewError(msg string) error {
	return &err{msg: msg}
}
