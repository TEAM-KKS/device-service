package driver

import (
	"encoding/json"
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/stianeikeland/go-rpio"
	"time"
)

type DynamicLed struct {
	DefaultDeviceProfile
}

type Pin struct {
	Addr int	`json:"addr"`
	State bool `json:"state"`
}

var Pins map[int]bool

func init() {
	Pins = make(map[int]bool)
	Pins[19] = false
	setDynamicPinLow(19)
	Pins[13] = false
	setDynamicPinLow(13)
	Pins[26] = false
	setDynamicPinLow(26)
}

func (l *DynamicLed) get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	now := time.Now().UnixNano()
	res = make([]*dsModels.CommandValue, 1)
	states, err := json.Marshal(Pins)
	if err == nil {
		fmt.Println(string(states))
		res[0] = dsModels.NewStringValue(reqs[0].DeviceResourceName, now, string(states))
	} else {
		res[0] = dsModels.NewStringValue(reqs[0].DeviceResourceName, now, "")
	}
	return
}

func (l *DynamicLed) set(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	if params[0].DeviceResourceName == "states" {
		if value, err := params[0].StringValue(); err == nil {
			pin := Pin{}
			//fmt.Println("Received this from AppService: ", value)
			jerr := json.Unmarshal([]byte(value), &pin)
			//fmt.Println("Jerr: ", jerr)
			err := rpio.Open()
			//fmt.Println("err: ", err)
			//led := rpio.Pin(pin.Addr) // pin for light sensor
			//err := rpio.Open()
			if err == nil && jerr == nil {
				//fmt.Println("pin->Addr: ", pin.Addr)
				led := rpio.Pin(pin.Addr)
				led.Output()
				if pin.State {
					led.High()
				} else {
					led.Low()
				}
				fmt.Println(pin)
				Pins[pin.Addr] = pin.State
			}
			rpio.Close()
		}
	}
	return nil
}

func setDynamicPinLow(pin int) {
	led := rpio.Pin(pin)
	err := rpio.Open()
	if err == nil {
		led.Output()
		led.Low()
		rpio.Close()
	}
}

