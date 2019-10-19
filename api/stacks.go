package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Stack struct {
	StackID   int
	StackName string
}

func GetStacks(environment string) []Stack {

	resp := APICall(environment, "/api/v1/stacks", "")

	if resp.Code == 200 {
		//fmt.Println(resp.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(resp.Body), &result)

		var retStacks []Stack

		// The object stored in the "birds" key is also stored as
		// a map[string]interface{} type, and its type is asserted from
		// the interface{} type
		stacks := result["stacks"].(map[string]interface{})

		for key, value := range stacks {
			// Each value is an interface{} type, that is type asserted as a string

			id, _ := strconv.Atoi(key)
			newStack := Stack{
				StackID:   id,
				StackName: value.(string),
			}
			retStacks = append(retStacks, newStack)
		}

		return retStacks
	} else {
		fmt.Println(resp.Body)
	}

	return []Stack{}
}
