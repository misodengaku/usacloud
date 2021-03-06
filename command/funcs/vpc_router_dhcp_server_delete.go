package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterDhcpServerDelete(ctx command.Context, params *params.DhcpServerDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP servers", params.Id)
	}

	cnf := p.Settings.Router.FindDHCPServerAt(params.Index)
	if cnf == nil {
		return fmt.Errorf("DHCP server is not found on eth%d", params.Index)
	}

	p.Settings.Router.RemoveDHCPServerAt(params.Index)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", err)
	}

	return nil
}
