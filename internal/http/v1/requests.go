package v1

type SkipLimit struct {
	Skip  *int `query:"skip"`
	Limit *int `query:"limit"`
}
