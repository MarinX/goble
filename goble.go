// ble
package goble

import (
	"fmt"
	"github.com/MarinX/serial"
	"strings"
	"time"
)

func New(dev string) (*Ble, error) {

	c := &serial.Config{Name: dev, Baud: 9600, ReadTimeout: 1 * time.Second}
	f, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return &Ble{
		fd: f,
	}, nil
}

func (t *Ble) Close() error {
	return t.fd.Close()
}

func (t *Ble) GetDeviceName() *BleResponse {
	return t.write_read(AT_GET_DEVICE_NAME)
}

func (t *Ble) SetDeviceName(name string) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_DEVICE_NAME, name))
}

func (t *Ble) SetMode(m BleMode) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_MODE, int(m)))
}

func (t *Ble) GetMode() *BleResponse {
	return t.write_read(AT_GET_MODE)
}

func (t *Ble) GetPin() *BleResponse {
	return t.write_read(AT_GET_PIN)
}

func (t *Ble) SetPin(pin int) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_PIN, pin))
}

func (t *Ble) GetBoundMode() *BleResponse {
	return t.write_read(AT_GET_BOUND_MODE)
}

func (t *Ble) SetBoundMode(mode BleBondMode) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_BOUND_MODE, int(mode)))
}

func (t *Ble) GetLastConnectedDeviceAddress() *BleResponse {
	return t.write_read(AT_LAST_DEVICE_ADDRESS)
}

func (t *Ble) GetRSSI() *BleResponse {
	return t.write_read(AT_RSSI)
}

func (t *Ble) GetPIO(pio BlePIOPin) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_GET_PIO, pio))
}

func (t *Ble) SetPIO(pio BlePIOPin, value BlePIO) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_PIO, pio, int(value)))
}

func (t *Ble) GetRole() *BleResponse {
	return t.write_read(AT_GET_ROLE)
}

func (t *Ble) SetRole(role BleRole) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_ROLE, role))
}

func (t *Ble) Reset() *BleResponse {
	return t.write_read(AT_RESET)
}

func (t *Ble) SofwareVersion() *BleResponse {
	return t.write_read(AT_SOFWATE_VERSION)
}

func (t *Ble) ClearLastConnectedDevice() *BleResponse {
	return t.write_read(AT_CLEAR)
}

func (t *Ble) SetMajor(major int) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_MAJOR, major))
}

func (t *Ble) SetMinor(minor int) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_MAJOR, minor))
}

func (t *Ble) SetAdvertasingInterval(interval int) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_ADVERTASING, interval))
}

func (t *Ble) SetNoneConnectable() *BleResponse {
	return t.write_read(AT_NONE_CONNECTABLE)
}

func (t *Ble) SetBeaconMode(mode BleBeaconMode) *BleResponse {
	return t.write_read(fmt.Sprintf(AT_SET_BEACON, mode))
}

func (t *Ble) GetBeaconMode() *BleResponse {
	return t.write_read(AT_GET_BEACON)
}
func (t *Ble) FactoryReset() *BleResponse {
	return t.write_read(AT_FACTORY)
}

func (t *Ble) write_read(cmd string) *BleResponse {

	if _, err := t.fd.Write([]byte(cmd)); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	buff := make([]byte, 512)

	if _, err := t.fd.Read(buff); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	sep := strings.Split(string(buff), ":")
	if len(sep) == 2 {
		return &BleResponse{
			Result: string(buff),
			Param:  sep[1],
		}
	}
	return &BleResponse{
		Result: string(buff),
	}

}

func (t *Ble) Write_Read(cmd string) *BleResponse {

	if _, err := t.fd.Write([]byte(cmd)); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	buff := make([]byte, 512)

	if _, err := t.fd.Read(buff); err != nil {
		return &BleResponse{
			Error: err,
		}
	}

	sep := strings.Split(string(buff), ":")
	if len(sep) == 2 {
		return &BleResponse{
			Result: string(buff),
			Param:  sep[1],
		}
	}
	return &BleResponse{
		Result: string(buff),
	}

}
