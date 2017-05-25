package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/vnc"
)

func ServerVnc(ctx Context, params *VncServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVnc is failed: %s", e)
	}

	if !p.IsUp() && params.WaitForBoot {

		err := internal.ExecWithProgress(
			fmt.Sprintf("Still booting[ID:%d]...", params.Id),
			fmt.Sprintf("Connect to server[ID:%d]", params.Id),
			GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ServerVnc is failed: %s", err)
		}
	}
	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVnc is failed: %s", e)
	}

	return vnc.OpenVNCClient(vncProxyInfo)
}
