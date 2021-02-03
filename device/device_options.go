package device

import "golang.zx2c4.com/wireguard/conn"

var defaults = []Option{
	WithLogger(NewLogger(LogLevelDebug, "" /* nothing to prepend */)),
	WithBindListener(func(_ *Device, port uint16) (conn.Bind, uint16, error) { return conn.CreateBind(port) }),
}

// Option represents a single option passed NewDevice
type Option func(*Device)

// WithLogger is used to pass the given logger to the device
func WithLogger(logger *Logger) Option {
	return func(device *Device) { device.log = logger }
}

// WithBindListener sets the provided function as the listener for "create bind" events
func WithBindListener(fn func(*Device, uint16) (conn.Bind, uint16, error)) Option {
	return func(device *Device) { device.onCreateBind = fn }
}
