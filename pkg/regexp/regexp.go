package regexp

import (
	"github.com/dlclark/regexp2"
)

// 正则匹配模式（编译时确保正则有效）
var (
	Account = regexp2.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,15}$`, 0)
	Pwd     = regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,32}$`, 0)
)

// Match 执行正则匹配，正则通过 MustCompile 预编译，MatchString 不会返回错误
func Match(re *regexp2.Regexp, str string) bool {
	res, err := re.MatchString(str)
	if err != nil {
		// 正则通过 MustCompile 创建，正常情况下不会出错
		return false
	}
	return res
}
