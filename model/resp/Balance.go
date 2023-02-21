package resp

type CheckBalance struct {
	Success bool
	Code    int
	Balance float32
	Error   bool
}
