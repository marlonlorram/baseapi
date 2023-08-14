package entities

// Err representa um erro com uma mensagem descritiva.
type Err struct {
	message string
}

// Error retorna a descrição do erro.
func (e *Err) Error() string {
	return e.message
}

// notFound é usado quando um recurso não é encontrado.
type notFound struct {
	Err
}

// NewNotFound cria um erro indicando que um recurso não foi encontrado.
func NewNotFound(msg string) error {
	return &notFound{Err{message: msg}}
}

// IsNotFound verifica se o erro é do tipo não encontrado.
func IsNotFound(err error) bool {
	_, ok := err.(*notFound)
	return ok
}

// unauthorized é usado para operações que pecisam de permissão.
type unauthorized struct {
	Err
}

// NewUnauthorized cria um erro indicando falta de autorização.
func NewUnauthorized(msg string) error {
	return &unauthorized{Err{message: msg}}
}

// IsUnauthorized verifica se o erro é devido a uma operação não autorizada.
func IsUnauthorized(err error) bool {
	_, ok := err.(*unauthorized)
	return ok
}

// badRequest é usado para solicitações inválidas.
type badRequest struct {
	Err
}

// NewBadRequest cria um erro indicando que uma solicitação é inválida.
func NewBadRequest(msg string) error {
	return &badRequest{Err{message: msg}}
}

// IsBadRequest verifica se o erro é resultado de uma solicitação inválida.
func IsBadRequest(err error) bool {
	_, ok := err.(*badRequest)
	return ok
}

// forbidden é usado para operações proibidas.
type forbidden struct {
	Err
}

// NewForbidden cria um erro indicando que uma operação é proibida.
func NewForbidden(msg string) error {
	return &forbidden{Err{message: msg}}
}

// IsForbidden verifica se o erro indica uma operação proibida.
func IsForbidden(err error) bool {
	_, ok := err.(*forbidden)
	return ok
}
