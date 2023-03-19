package externals

type DeviceType int

const (
	DeviceTypeDesktop = iota
	DeviceTypeMobile
)

type Device struct {
	Name string     `json:"name"`
	Type DeviceType `json:"type"`
}
