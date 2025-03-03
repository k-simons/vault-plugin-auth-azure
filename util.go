package azureauth

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/hashicorp/vault/sdk/helper/useragent"
	"github.com/hashicorp/vault/sdk/logical"
)

// Using the same time parsing logic from https://github.com/coreos/go-oidc
// This code is licensed under the Apache 2.0 license
type jsonTime time.Time

func (j *jsonTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	var unix int64

	if t, err := n.Int64(); err == nil {
		unix = t
	} else {
		f, err := n.Float64()
		if err != nil {
			return err
		}
		unix = int64(f)
	}
	*j = jsonTime(time.Unix(unix, 0))
	return nil
}

// strListContains does a case-insensitive search of the string
// list for the value
func strListContains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if strings.EqualFold(item, needle) {
			return true
		}
	}
	return false
}

// userAgent returns the User-Agent header that's included in HTTP requests
// made by this plugin. Returns an empty string if the pluginEnv is nil.
func userAgent(pluginEnv *logical.PluginEnvironment) string {
	if pluginEnv == nil {
		return ""
	}

	return useragent.PluginString(pluginEnv, "azure-auth")
}
