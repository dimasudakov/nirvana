package model

// ExceptionAttributes - структура, представляющая атрибуты исключения.
type ExceptionAttributes struct {
	ClientID string `json:"client_id,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
}
