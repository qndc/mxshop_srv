package comm

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

var options = &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}

// Encode 加密
func Encode(pwd string) string {

	// 默认加密
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)

	// 自定义加密
	salt, encodedPwd := password.Encode(pwd, options)

	//加密方法二
	//Md5 := md5.New()
	//_, _ = io.WriteString(Md5, code)
	//hex.EncodeToString(Md5.Sum(nil))
	return fmt.Sprintf("sha512-%s-%s", salt, encodedPwd)
}

// Verify 验密
func Verify(rawPwd string, encodedPwd string) bool {
	strs := strings.Split(encodedPwd, "-")
	check := password.Verify(rawPwd, strs[1], strs[2], options)
	return check
}
