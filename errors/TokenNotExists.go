package errors

type TokenNotExists struct{}

func (t *TokenNotExists) Error() string {
	return "token not exists"
}
