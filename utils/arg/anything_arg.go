package arg

type Arg struct {
	AnythingArray []interface{}
	AnythingMap   map[string]interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	String        *string
}
