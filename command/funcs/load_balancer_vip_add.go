package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerVipAdd(ctx command.Context, params *params.VipAddLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", e)
	}

	// validate
	if len(p.Settings.LoadBalancer) >= 4 {
		return fmt.Errorf("LoadBalancer already has maximum count of VIP")
	}

	for _, v := range p.Settings.LoadBalancer {
		if v.VirtualIPAddress == params.Vip {
			return fmt.Errorf("VIP(%q) is already used", params.Vip)
		}
	}

	// set params
	var vip = &sacloud.LoadBalancerSetting{
		VirtualIPAddress: params.Vip,
		Port:             fmt.Sprintf("%d", params.Port),
		DelayLoop:        fmt.Sprintf("%d", params.DelayLoop),
		SorryServer:      params.SorryServer,
	}

	p.AddLoadBalancerSetting(vip)
	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", err)
	}

	return nil
}
