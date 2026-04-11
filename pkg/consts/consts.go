package consts

// User status constants
const (
	UserStatusNormal    = 0 // 正常
	UserStatusCancelled = 1 // 已注销
	UserStatusDisabled  = 2 // 已注销（软删除状态）
)

// User gender constants
const (
	GenderFemale  = 0 // 女
	GenderMale    = 1 // 男
	GenderUnknown = 2 // 未知
)

// Post status constants
const (
	PostStatusNormal   = 0 // 正常
	PostStatusDeleted  = 1 // 已删除
)
