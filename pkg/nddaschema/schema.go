package ndda

type Schema interface {
	NewDevice(name string) Device
}

func NewSchema() Schema {
	return &schema{
		// parent nil
		// children
		devices: make(map[string]Device),
		// data key
	}
}

type schema struct {
	// parent is nil
	// children
	devices map[string]Device
	// data is nil
}

func (x *schema) NewDevice(name string) Device {
	if _, ok := x.devices[name]; !ok {
		x.devices[name] = NewDevice(x, name)
	}
	return x.devices[name]
}
