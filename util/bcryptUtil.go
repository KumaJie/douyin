package util

import "golang.org/x/crypto/bcrypt"

// HashPwd 使用bcrypt对密码进行加密
func HashPwd(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPwd 比较两段密码是否相同
func VerifyPwd(dbPwd string, inputPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(inputPwd)); err != nil {
		return false
	}
	return true
}
