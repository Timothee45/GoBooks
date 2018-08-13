package error

type Error struct {
	Code 		int
	Label		string
}

var IncorrectId = Error{Code: 415, Label: "Incorrect id"}
var BadToken = Error{Code: 409, Label: "Bad Token"}

func GenerateError(code int, label string) Error {
	return Error{Code: code, Label: label}
}
