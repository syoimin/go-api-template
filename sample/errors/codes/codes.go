package codes

type Code string

const (
	OK                    Code = "OK"
	Unexpected            Code = "UnexpectedError"
	BadRequest            Code = "BadRequest"
	Unauthorized          Code = "UnauthorizedError"
	Forbidden             Code = "ForbiddenError"
	NotFound              Code = "NotFoundError"
	DatabaseError         Code = "DatabaseError"
	ParameterIllegalError Code = "ParameterIllegalError"
	AuthForbiddenError    Code = "AuthForbiddenError"
)
