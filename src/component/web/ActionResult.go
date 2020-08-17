package web

type ActionResultInterface interface {
	Content() interface{}
	StatusCode() int
	Error() error
}

type ActionResult struct {
	content    interface{}
	statusCode int
	httpError  error
}

func NewActionResult(content interface{}, statusCode int, httpError error) *ActionResult {
	return &ActionResult{
		content,
		statusCode,
		httpError,
	}
}

func (r *ActionResult) Content() interface{} {
	return r.content
}

func (r *ActionResult) StatusCode() int {
	return r.statusCode
}

func (r *ActionResult) Error() error {
	return r.httpError
}

var DefaultSuccessContent = struct {
	Ok bool `json:"ok"`
}{Ok: true}

var DefaultFailContent = struct {
	Ok bool `json:"ok"`
}{Ok: false}
