package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductLicenseListCompleteArgs(ctx command.Context, params *params.ListProductLicenseParam, cur, prev, commandName string) {

}

func ProductLicenseReadCompleteArgs(ctx command.Context, params *params.ReadProductLicenseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductLicenseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LicenseInfo {
		fmt.Println(res.LicenseInfo[i].ID)

	}

}
