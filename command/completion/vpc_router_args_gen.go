package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterListCompleteArgs(ctx command.Context, params *params.ListVPCRouterParam, cur, prev, commandName string) {

}

func VPCRouterCreateCompleteArgs(ctx command.Context, params *params.CreateVPCRouterParam, cur, prev, commandName string) {

}

func VPCRouterReadCompleteArgs(ctx command.Context, params *params.ReadVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterUpdateCompleteArgs(ctx command.Context, params *params.UpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDeleteCompleteArgs(ctx command.Context, params *params.DeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterBootCompleteArgs(ctx command.Context, params *params.BootVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterShutdownCompleteArgs(ctx command.Context, params *params.ShutdownVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterShutdownForceCompleteArgs(ctx command.Context, params *params.ShutdownForceVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterResetCompleteArgs(ctx command.Context, params *params.ResetVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterWaitForBootCompleteArgs(ctx command.Context, params *params.WaitForBootVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterWaitForDownCompleteArgs(ctx command.Context, params *params.WaitForDownVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterInterfaceInfoCompleteArgs(ctx command.Context, params *params.InterfaceInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterInterfaceConnectCompleteArgs(ctx command.Context, params *params.InterfaceConnectVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterInterfaceUpdateCompleteArgs(ctx command.Context, params *params.InterfaceUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterInterfaceDisconnectCompleteArgs(ctx command.Context, params *params.InterfaceDisconnectVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticNatInfoCompleteArgs(ctx command.Context, params *params.StaticNatInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticNatAddCompleteArgs(ctx command.Context, params *params.StaticNatAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticNatUpdateCompleteArgs(ctx command.Context, params *params.StaticNatUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticNatDeleteCompleteArgs(ctx command.Context, params *params.StaticNatDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPortForwardingInfoCompleteArgs(ctx command.Context, params *params.PortForwardingInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPortForwardingAddCompleteArgs(ctx command.Context, params *params.PortForwardingAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPortForwardingUpdateCompleteArgs(ctx command.Context, params *params.PortForwardingUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPortForwardingDeleteCompleteArgs(ctx command.Context, params *params.PortForwardingDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterFirewallInfoCompleteArgs(ctx command.Context, params *params.FirewallInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterFirewallAddCompleteArgs(ctx command.Context, params *params.FirewallAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterFirewallUpdateCompleteArgs(ctx command.Context, params *params.FirewallUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterFirewallDeleteCompleteArgs(ctx command.Context, params *params.FirewallDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpServerInfoCompleteArgs(ctx command.Context, params *params.DhcpServerInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpServerAddCompleteArgs(ctx command.Context, params *params.DhcpServerAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpServerUpdateCompleteArgs(ctx command.Context, params *params.DhcpServerUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpServerDeleteCompleteArgs(ctx command.Context, params *params.DhcpServerDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpStaticMappingInfoCompleteArgs(ctx command.Context, params *params.DhcpStaticMappingInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpStaticMappingAddCompleteArgs(ctx command.Context, params *params.DhcpStaticMappingAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpStaticMappingUpdateCompleteArgs(ctx command.Context, params *params.DhcpStaticMappingUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterDhcpStaticMappingDeleteCompleteArgs(ctx command.Context, params *params.DhcpStaticMappingDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPptpServerInfoCompleteArgs(ctx command.Context, params *params.PptpServerInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterPptpServerUpdateCompleteArgs(ctx command.Context, params *params.PptpServerUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterL2tpServerInfoCompleteArgs(ctx command.Context, params *params.L2tpServerInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterL2tpServerUpdateCompleteArgs(ctx command.Context, params *params.L2tpServerUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterUserInfoCompleteArgs(ctx command.Context, params *params.UserInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterUserAddCompleteArgs(ctx command.Context, params *params.UserAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterUserUpdateCompleteArgs(ctx command.Context, params *params.UserUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterUserDeleteCompleteArgs(ctx command.Context, params *params.UserDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterSiteToSiteVpnInfoCompleteArgs(ctx command.Context, params *params.SiteToSiteVpnInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterSiteToSiteVpnAddCompleteArgs(ctx command.Context, params *params.SiteToSiteVpnAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterSiteToSiteVpnUpdateCompleteArgs(ctx command.Context, params *params.SiteToSiteVpnUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterSiteToSiteVpnDeleteCompleteArgs(ctx command.Context, params *params.SiteToSiteVpnDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticRouteInfoCompleteArgs(ctx command.Context, params *params.StaticRouteInfoVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticRouteAddCompleteArgs(ctx command.Context, params *params.StaticRouteAddVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticRouteUpdateCompleteArgs(ctx command.Context, params *params.StaticRouteUpdateVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterStaticRouteDeleteCompleteArgs(ctx command.Context, params *params.StaticRouteDeleteVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterMonitorCompleteArgs(ctx command.Context, params *params.MonitorVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func VPCRouterLogsCompleteArgs(ctx command.Context, params *params.LogsVPCRouterParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetVPCRouterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.VPCRouters {
		fmt.Println(res.VPCRouters[i].ID)
		var target interface{} = &res.VPCRouters[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
