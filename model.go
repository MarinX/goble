// model
package goble

import (
	"github.com/MarinX/serial"
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
	AT_GET_STATUS          = "AT+PIO1?"
	AT_SET_STATUS          = "AT+PIO1%d"
	AT_GET_PIO             = "AT+PIO%s?"
	AT_SET_PIO             = "AT+PIO%s%d"
	AT_SET_ROLE            = "AT+ROLE%d"
	AT_GET_ROLE            = "AT+ROLE?"
	AT_SET_MAJOR           = "AT+MARJ0x%d"
	AT_SET_MINOR           = "AT+MINO0xFA01"
	AT_ADVERTASING         = "AT+ADVI%d"
	AT_NONE_CONNECTABLE    = "AT+ADTY3"
	AT_GET_BEACON          = "AT+IBEA?"
	AT_SET_BEACON          = "AT+IBEA%d"
)

type BleMode int
type BleBondMode int
type BlePIO int
type BlePIOPin string
type BleRole int
type BleBeaconMode int

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

//Ble pin output
const (
	LOW  BlePIO = 0
	HIGH BlePIO = 1
)

//Ble pio pins
const (
	PIN_2 BlePIOPin = "2"
	PIN_3 BlePIOPin = "3"
	PIN_4 BlePIOPin = "4"
	PIN_5 BlePIOPin = "5"
	PIN_6 BlePIOPin = "6"
	PIN_7 BlePIOPin = "7"
	PIN_8 BlePIOPin = "8"
	PIN_9 BlePIOPin = "9"
	PIN_A BlePIOPin = "A"
	PIN_B BlePIOPin = "B"
)

//Ble role
const (
	ROLE_SLAVER BleRole = 0
	ROLE_MASTER BleRole = 1
)

//Ble beacon mode
const (
	BEACON_DISABLED BleBeaconMode = 0
	BEACON_ENABLE   BleBeaconMode = 1
)

type Ble struct {
	fd *serial.Port
}

type BleResponse struct {
	Error  error
	Result string
	Param  interface{}
}
