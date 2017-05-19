package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DatabaseReset(ctx Context, params *ResetDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseReset is failed: %s", e)
	}

	// set params
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still resetting[ID:%d]...", params.Id),
		fmt.Sprintf("Reset database[ID:%d]", params.Id),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.RebootForce(params.Id)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilDatabaseRunning(params.Id, client.DefaultTimeoutDuration, 30)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DatabaseReset is failed: %s", err)
	}

	return nil

}
