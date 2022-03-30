package abstract

type Abstract interface {
	GetInterfaceName(itfcName string) (string, error)
}
