package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductDiskListCompleteArgs(ctx command.Context, params *params.ListProductDiskParam, cur, prev, commandName string) {

}

func ProductDiskReadCompleteArgs(ctx command.Context, params *params.ReadProductDiskParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductDiskAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.DiskPlans {
		fmt.Println(res.DiskPlans[i].ID)

	}

}
