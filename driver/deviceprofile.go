package driver

import (
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

type DeviceProfile interface {
	get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error)
	set(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error
}

type DefaultDeviceProfile struct {
	DeviceProfile
}

func (d *DefaultDeviceProfile) get(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	err = fmt.Errorf("User calling unimplemented method GET for ", deviceName)
	return nil, err
}

func (d *DefaultDeviceProfile) set(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	err := fmt.Errorf("User calling unimplemented method SET for ", deviceName)
	return err
}
