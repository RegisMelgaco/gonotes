package common

//ErrorHolder - Model para enviar erros simples ao cliente
type ErrorHolder struct {
	Error string `json:"error"`
}
