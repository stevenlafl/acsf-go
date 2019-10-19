package api

import (
	"encoding/json"
	"fmt"
)

type VCS struct {
	Available []string
	Current   string
}

func GetVCS(environment string, stackId int) VCS {

	resp := APICall(environment, fmt.Sprintf("/api/v1/vcs?type=sites&stack_id=%d", stackId), "")

	if resp.Code == 200 {
		//fmt.Println(resp.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(resp.Body), &result)

		var retRefs []string

		// The object stored in the "birds" key is also stored as
		// a map[string]interface{} type, and its type is asserted from
		// the interface{} type

		for _, ref := range result["available"].([]interface{}) {
			// Each value is an interface{} type, that is type asserted as a string

			retRefs = append(retRefs, ref.(string))
		}

		return VCS{
			Available: retRefs,
			Current:   result["current"].(string),
		}
	} else {
		fmt.Println(resp.Body)
	}

	return VCS{}
}
