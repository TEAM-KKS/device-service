package driver

import (
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"math/rand"
	"time"
)

type RandomBool struct {
	DefaultDeviceProfile
}

func (r *RandomBool) get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	res = make([]*dsModels.CommandValue, 1)
	now := time.Now().UnixNano()
	res[0], _ = dsModels.NewBoolValue(reqs[0].DeviceResourceName, now,  bool(rand.Intn(100) % 2 == 0))
	return
}
