package cli

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/profile"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/helper/migration"
	"os"
	"strconv"
	"strings"
)

func checkConfigVersion() error {
	return migration.CheckConfigVersion()
}

type FlagHandler interface {
	IsSet(name string) bool
	Set(name, value string) error
	String(name string) string
	StringSlice(name string) []string
}

func applyConfigFromFile(c FlagHandler) error {
	profileKey := "profile"
	profileName := c.String(profileKey)
	if profileName == "" {
		n, err := profile.GetCurrentName()
		if err != nil {
			return fmt.Errorf("Failed to load current profile: %s", err)
		}
		profileName = n
	}

	// load config file
	v, err := profile.LoadConfigFile(profileName)
	if err != nil {
		return err
	}

	if !c.IsSet("token") && v.AccessToken != "" {
		c.Set("token", v.AccessToken)
		command.GlobalOption.AccessToken = v.AccessToken
	}
	if !c.IsSet("secret") && v.AccessTokenSecret != "" {
		c.Set("secret", v.AccessTokenSecret)
		command.GlobalOption.AccessTokenSecret = v.AccessTokenSecret
	}
	if !c.IsSet("zone") && v.Zone != "" {
		c.Set("zone", v.Zone)
		command.GlobalOption.Zone = v.Zone
	}

	// for string-slice
	zones := []string{}
	if c.IsSet("zones") {
		zones = c.StringSlice("zones")
	} else if len(v.Zones) > 0 {
		zones = v.Zones
	} else {
		if z, ok := os.LookupEnv("USACLOUD_ZONES"); ok {
			zones = strings.Split(z, ",")
		}
	}
	command.GlobalOption.Zones = zones

	if !c.IsSet("api-root-url") && v.APIRootURL != "" {
		c.Set("api-root-url", v.APIRootURL)
		command.GlobalOption.APIRootURL = v.APIRootURL
	}

	if len(command.GlobalOption.Zones) > 0 {
		define.AllowZones = command.GlobalOption.Zones
	}
	if command.GlobalOption.APIRootURL != "" {
		api.SakuraCloudAPIRoot = command.GlobalOption.APIRootURL
	}

	return nil
}

func toSakuraID(id string) (int64, bool) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}

func hasTags(target interface{}, tags []string) bool {
	type tagHandler interface {
		HasTag(target string) bool
	}

	tagHolder, ok := target.(tagHandler)
	if !ok {
		return false
	}

	// 完全一致 + AND条件
	res := true
	for _, p := range tags {
		if !tagHolder.HasTag(p) {
			res = false
			break
		}
	}
	return res
}
