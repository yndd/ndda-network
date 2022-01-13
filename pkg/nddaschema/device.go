package ndda

type Device interface {
	// children
	NewInterface(name string) Interface
	NewNetworkInstance(name string) NetworkInstance
	NewSystemPlatform() SystemPlatform
}

func NewDevice(p Schema, name string) Device {
	return &device{
		// parent
		parent: p,
		// children
		Interface:       make(map[string]Interface),
		NetworkInstance: make(map[string]NetworkInstance),
		// data key
	}
}

type device struct {
	// parent
	parent Schema
	// children
	Interface       map[string]Interface
	NetworkInstance map[string]NetworkInstance
	SystemPlatform  SystemPlatform
	// data
}

func (x *device) NewInterface(name string) Interface {
	if _, ok := x.Interface[name]; !ok {
		x.Interface[name] = NewInterface(x, name)
	}
	return x.Interface[name]
}

func (x *device) NewNetworkInstance(name string) NetworkInstance {
	if _, ok := x.NetworkInstance[name]; !ok {
		x.NetworkInstance[name] = NewNetworkInstance(x, name)
	}
	return x.NetworkInstance[name]
}

func (x *device) NewSystemPlatform() SystemPlatform {
	if x.SystemPlatform == nil {
		x.SystemPlatform = NewSystemPlatform(x)
	}
	return x.SystemPlatform
}
