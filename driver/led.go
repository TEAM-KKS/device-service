package driver

import (
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/stianeikeland/go-rpio"
	"sync"
	"time"
)

type Led struct {
	DefaultDeviceProfile
	mux sync.Mutex
}

var pin = 17
var state = false

func init() {
	setPinLow(pin)
}

func (l *Led) get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	l.mux.Lock()
	defer l.mux.Unlock()
	now := time.Now().UnixNano()
	res = make([]*dsModels.CommandValue, 1)
	res[0], _ = dsModels.NewBoolValue(reqs[0].DeviceResourceName, now, state)
	return
}

func (l *Led) set(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	l.mux.Lock()
	defer l.mux.Unlock()
	value, err := params[0].BoolValue()
	if err == nil {
		err := rpio.Open()
		if err == nil {
			led := rpio.Pin(pin)
			led.Output()
			if value {
				led.High()
			} else {
				led.Low()
			}
			state = value
			rpio.Close()
		} else {
			return err
		}
	} else {
		return err
	}
	return err
}

func setPinLow(pin int) {
	led := rpio.Pin(pin)
	err := rpio.Open()
	if err == nil {
		led.Output()
		led.Low()
		rpio.Close()
	}
}

