// ble
package goble

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const (
	AT_GET_DEVICE_NAME     = "AT+NAME?"
	AT_SET_DEVICE_NAME     = "AT+NAME%s"
	AT_RESET               = "AT+RESET"
	AT_GET_MODE            = "AT+MODE?"
	AT_SET_MODE            = "AT+MODE%d"
	AT_FACTORY             = "AT+RENEW"
	AT_GET_PIN             = "AT+PASS?"
	AT_SET_PIN             = "AT+PASS%d"
	AT_SOFWATE_VERSION     = "AT+VERS?"
	AT_LAST_DEVICE_ADDRESS = "AT+RADD?"
	AT_RSSI                = "AT+RSSI?"
	AT_CLEAR               = "AT+CLEAR"
	AT_GET_BOUND_MODE      = "AT+TYPE?"
	AT_SET_BOUND_MODE      = "AT+TYPE%d"
)

type BleMode int
type BleBondMode int

//Ble mode
const (
	TRANSMISSION  BleMode = 0
	REMOTE        BleMode = 1
	ZERO_PLUS_ONE BleMode = 2
)

//Ble bond mode
const (
	NOT_NEED_PIN_CODE BleBondMode = 0
	NEEDED_PIN_CODE   BleBondMode = 1
)

type Ble struct {
	fd *os.File
}

type BleResponse struct {
	Error  error
	Result string
	Param  interface{}
}

func New(dev string) (*Ble, error) {

	f, err := os.OpenFile(dev, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	if err != nil {
		return nil, err
	}

	t := syscall.Termios{
		Iflag:  syscall.IGNPAR,
		Cflag:  syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B9600,
		Cc:     [32]uint8{syscall.VMIN: 64, syscall.VTIME: 1},
		Ispeed: syscall.B9600,
		Ospeed: syscall.B9600,
	}

	fd := f.Fd()

	if _, _, errno := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&t)),
		0,
		0,
		0,
	); errno != 0 {
		return nil, errno
	}

	if err = syscall.SetNonblock(int(fd), false); err != nil {
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

func (t *Ble) Reset() *BleResponse {
	return t.write_read(AT_RESET)
}

func (t *Ble) SofwareVersion() *BleResponse {
	return t.write_read(AT_SOFWATE_VERSION)
}

func (t *Ble) ClearLastConnectedDevice() *BleResponse {
	return t.write_read(AT_CLEAR)
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
