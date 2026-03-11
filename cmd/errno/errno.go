package errno

const (
	CodeSuccess = 0

	//grades
	CodeInvalidStudentID = 4001
	CodeInvalidGradeBoyd = 4002
	CodeAddGradeFailed   = 4003
	CodeStudentNotFound  = 4004

	//log
	CodeLogReadBodyFailed = 5001
	CodeEmptyLogMessage   = 5002

	//registry
	CodeInvalidRegistryReq = 6001
	CodeRegisterFailed     = 6002
	CodeReadBodyFailed     = 6003
	CodeEmptyServiceURL    = 6004
	CodeDeregisterFailed   = 6005
)
