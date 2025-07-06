package response

import "gorm.io/gorm"

func ReturnErrorWithCode(err error, response *FinalResponse) {
	response.Error = &AppError{}
	switch err {
	case gorm.ErrRecordNotFound:
		response.Error.Code = "404"
		response.Error.Message = "record not found check variables that you pass"
	case gorm.ErrDuplicatedKey:
		response.Error.Code = "409"
		response.Error.Message = "duplicate key --> primary key constraint"
	case gorm.ErrForeignKeyViolated:
		response.Error.Code = "23505"
		response.Error.Message = "foreign key violation."
	case gorm.ErrInvalidDB:
		response.Error.Code = "252525"
		response.Error.Message = "database connection problem"
	default:
		response.Error.Code = "898989"
		response.Error.Message = "pta nahi kya problem hain. 😔"
	}
	response.Error.Details = string(err.Error())
}
