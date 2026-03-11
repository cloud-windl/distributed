package errno

const (
	CodeSuccess = 0

	CodeInvalidStudentID    = 4001
	CodeInvalidStudentBody  = 4002
	CodeCreateStudentFailed = 4003
	CodeUpdateStudentFailed = 4004
	CodeDeleteStudentFailed = 4005
	CodeStudentNotFound     = 4006
	CodeQueryStudentsFailed = 4007

	CodeInvalidGradeID    = 4101
	CodeInvalidGradeBody  = 4102
	CodeAddGradeFailed    = 4103
	CodeUpdateGradeFailed = 4104
	CodeDeleteGradeFailed = 4105
	CodeGradeNotFound     = 4106
	CodeQueryGradesFailed = 4107

	CodeLogReadBodyFailed = 5001
	CodeEmptyLogMessage   = 5002

	CodeInvalidRegistryReq = 6001
	CodeRegisterFailed     = 6002
	CodeReadBodyFailed     = 6003
	CodeEmptyServiceURL    = 6004
	CodeDeregisterFailed   = 6005
)
