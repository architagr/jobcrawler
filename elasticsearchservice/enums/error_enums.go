package enums

type ErrorCode int

const genericErrorCode int = 1001
const (
	ERROR_CODE_REQUEST_PARAM          ErrorCode = ErrorCode(iota + genericErrorCode)
	ERROR_CODE_REQUEST_INTERNAL_ERROR ErrorCode = ErrorCode(iota + genericErrorCode)
	ERROR_CODE_FILE_OPEN              ErrorCode = ErrorCode(iota + genericErrorCode)
)

const userprofileErrorCode int = 2001
const (
	ERROR_CODE_USERNAME_EXIST ErrorCode = ErrorCode(iota + userprofileErrorCode)
)

const dbErrorCode int = 2001
const (
	ERROR_CODE_ADD_NEW_USER    ErrorCode = ErrorCode(iota + dbErrorCode)
	ERROR_CODE_UPDATE_NEW_USER ErrorCode = ErrorCode(iota + dbErrorCode)
	ERROR_CODE_GET_USER        ErrorCode = ErrorCode(iota + dbErrorCode)
	ERROR_CODE_USER_NOT_FOUND  ErrorCode = ErrorCode(iota + dbErrorCode)
)
