package delivery

type CheckRequest struct {
	CheckIn  string `json:"check_in" form:"check_in"`
	CheckOut string `json:"check_out" form:"check_out"`
}
