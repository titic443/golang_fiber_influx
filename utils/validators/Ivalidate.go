package validators

type IValidateBody interface {
	MapType([]byte) []interface{}
}
