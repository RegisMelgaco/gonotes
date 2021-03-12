package common

//StoreError monta um ErrorHolder com o erro provido
func StoreError(err error) ErrorHolder {
	return ErrorHolder{Error: err.Error()}
}
