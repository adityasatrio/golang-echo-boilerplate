package errcode

const (
	EBL10001 = 10001
	EBL10002 = 10002
)

var businessLogicReason = map[int]string{
	EBL10001: "EBL10001: system parameter key already exist",
	EBL10002: "EBL10002: your explanation of error EBL = error business logic",
}

func BusinessLogicReason(code int) string {
	return businessLogicReason[code]
}
