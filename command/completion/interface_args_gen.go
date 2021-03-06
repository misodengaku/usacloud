package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfaceListCompleteArgs(ctx command.Context, params *params.ListInterfaceParam, cur, prev, commandName string) {

}

func InterfacePacketFilterConnectCompleteArgs(ctx command.Context, params *params.PacketFilterConnectInterfaceParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
		var target interface{} = &res.Interfaces[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InterfaceCreateCompleteArgs(ctx command.Context, params *params.CreateInterfaceParam, cur, prev, commandName string) {

}

func InterfacePacketFilterDisconnectCompleteArgs(ctx command.Context, params *params.PacketFilterDisconnectInterfaceParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
		var target interface{} = &res.Interfaces[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InterfaceReadCompleteArgs(ctx command.Context, params *params.ReadInterfaceParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
		var target interface{} = &res.Interfaces[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InterfaceUpdateCompleteArgs(ctx command.Context, params *params.UpdateInterfaceParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
		var target interface{} = &res.Interfaces[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InterfaceDeleteCompleteArgs(ctx command.Context, params *params.DeleteInterfaceParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
		var target interface{} = &res.Interfaces[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
