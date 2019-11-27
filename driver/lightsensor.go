package driver

import (
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/stianeikeland/go-rpio"
	"sync"
	"time"
	"os"
)

type LightSensor struct {
	DefaultDeviceProfile
	mux sync.Mutex
}

func (r *LightSensor) get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	err = rpio.Open()
	now := time.Now().UnixNano()
	count := 0
	if err == nil {
		pin := rpio.Pin(4) // pin for light sensor
		pin.Low()
		time.Sleep(time.Millisecond * 100)
		pin.Input()
		count = 0
		for pin.Read() == rpio.Low && count < 50{
			count++
		}
		if count >= 50 {
			count = 1
		} else {
			count = 0
		}
		fmt.Fprintf(os.Stdout,  "....... %s .......\n", count)
		rpio.Close()
	}
	res = make([]*dsModels.CommandValue, 1)
	res[0], _ = dsModels.NewInt32Value(reqs[0].DeviceResourceName, now, int32(count))
	return
}


