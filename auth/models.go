package auth

import "golang.org/x/crypto/bcrypt"

//UserCredentials tem as informações de autenticação dos usuários da aplicação
type UserCredentials struct {
	ID       uint64
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *UserCredentials) hashAndSalt() (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err == nil {
		u.Password = string(hash)
	}
	return
}

//LoginData contém toda informação entregue no login
type LoginData struct {
	Token string `binding:"required"`
}
