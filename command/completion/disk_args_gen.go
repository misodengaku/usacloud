package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DiskListCompleteArgs(ctx command.Context, params *params.ListDiskParam, cur, prev, commandName string) {

}

func DiskCreateCompleteArgs(ctx command.Context, params *params.CreateDiskParam, cur, prev, commandName string) {

}

func DiskReadCompleteArgs(ctx command.Context, params *params.ReadDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskUpdateCompleteArgs(ctx command.Context, params *params.UpdateDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskDeleteCompleteArgs(ctx command.Context, params *params.DeleteDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskEditCompleteArgs(ctx command.Context, params *params.EditDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskReinstallFromArchiveCompleteArgs(ctx command.Context, params *params.ReinstallFromArchiveDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskReinstallFromDiskCompleteArgs(ctx command.Context, params *params.ReinstallFromDiskDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskReinstallToBlankCompleteArgs(ctx command.Context, params *params.ReinstallToBlankDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskServerConnectCompleteArgs(ctx command.Context, params *params.ServerConnectDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskServerDisconnectCompleteArgs(ctx command.Context, params *params.ServerDisconnectDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskMonitorCompleteArgs(ctx command.Context, params *params.MonitorDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func DiskWaitForCopyCompleteArgs(ctx command.Context, params *params.WaitForCopyDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Disks {
		fmt.Println(res.Disks[i].ID)
		var target interface{} = &res.Disks[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
