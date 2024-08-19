package regexp

import (
	"github.com/dlclark/regexp2"
)

// 正则匹配模式
var (
	Account = regexp2.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,15}$`, 0)
	Pwd     = regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,32}$`, 0)
)

func Match(re *regexp2.Regexp, str string) bool {
	res, _ := re.MatchString(str)
	if res {
		return true
	} else {
		return false
	}
}
