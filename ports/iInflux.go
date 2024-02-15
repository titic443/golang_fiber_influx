package port

type IAmita interface {
	Write(map[string]string, map[string]interface{}, string) error
	Select1H(string) (interface{}, error)
}
