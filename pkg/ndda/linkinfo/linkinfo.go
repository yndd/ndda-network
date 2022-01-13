package linkinfo

type LinkInfoOption func(*LinkInfo)

func NewLinkInfo(opts ...LinkInfoOption) *LinkInfo {
	i := &LinkInfo{
		DeviceNames:  make(map[int]*string),
		ItfceNames:   make(map[int]*string),
		Ipv4Prefixes: make([]*string, 0),
		Ipv6Prefixes: make([]*string, 0),
	}

	for _, f := range opts {
		f(i)
	}

	return i
}

type LinkInfo struct {
	Name         *string
	DeviceNames  map[int]*string
	ItfceNames   map[int]*string
	Ipv4Prefixes []*string
	Ipv6Prefixes []*string
}

func (x *LinkInfo) GetLinkName() string {
	return *x.Name
}

func (x *LinkInfo) GetDeviceNames() map[int]*string {
	return x.DeviceNames
}

func (x *LinkInfo) GetDeviceName(i int) string {
	return *x.DeviceNames[i]
}

func (x *LinkInfo) GetItfceNames() map[int]*string {
	return x.ItfceNames
}

func (x *LinkInfo) GetItfceName(i int) string {
	return *x.ItfceNames[i]
}

func (x *LinkInfo) GetIpv4Prefixes() []*string {
	return x.Ipv4Prefixes
}

func (x *LinkInfo) GetIpv6Prefixes() []*string {
	return x.Ipv6Prefixes
}

func (x *LinkInfo) SetLinkName(s string) {
	x.Name = &s
}
