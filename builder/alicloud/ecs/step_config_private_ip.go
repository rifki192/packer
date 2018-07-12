package ecs

import (
	"context"
	"fmt"

	"github.com/denverdino/aliyungo/ecs"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type stepConfigAlicloudPrivateIP struct {
	privateIPAddress string
	RegionId         string
}

func (s *stepConfigAlicloudPrivateIP) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	instance := state.Get("instance").(*ecs.InstanceAttributesType)

	if len(instance.VpcAttributes.PrivateIpAddress.IpAddress) < 1 {
		state.Put("error", "No Private IP Available")
		ui.Say(fmt.Sprintf("Error no private ip"))
		return multistep.ActionHalt
	}
	ipaddress := instance.VpcAttributes.PrivateIpAddress.IpAddress[0]

	s.privateIPAddress = ipaddress
	ui.Say(fmt.Sprintf("Allocated private ip address %s.", ipaddress))
	state.Put("ipaddress", ipaddress)
	return multistep.ActionContinue
}

func (s *stepConfigAlicloudPrivateIP) Cleanup(state multistep.StateBag) {

}
