package utils

// hashAccount 对账号进行脱敏处理
func HashAccount(account string) string {
	if len(account) <= 4 {
		return "****"
	}
	return account[:3] + "****" + account[len(account)-2:]
}