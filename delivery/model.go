package delivery

type Request struct {
	AppId           string `form:"app"`
	Country         string `form:"country"`
	OperatingStstem string `form:"os"`
}
