package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
bcrypt加密算法介绍
bcrypt算法对于同一个密码，每次生成的hash不一样
*/
func main() {
	pwd1 := "123456"
	p1, _ := GetPwd(pwd1)
	p2, _ := GetPwd(pwd1)
	//同一个字符串生成的密码不一样
	fmt.Printf("p1: %v\n", string(p1))
	fmt.Printf("p2: %v\n", string(p2))

	compareResult := ComparePwd(string(p2), pwd1)
	fmt.Printf("compareResult: %v\n", compareResult)

}

// 生成密码
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //加密处理
	return hash, err
}

//比较密码 hashedPassword为加密后的值，password当前需要比较的明文
func ComparePwd(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
