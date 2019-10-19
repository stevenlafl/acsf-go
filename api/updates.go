package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Update struct {
	UpdateID int
	Added    int
	Status   int
}

type UpdateStatus struct {
	UpdateID   int
	Message    string
	Percentage int
	StartTime  int
	EndTime    int
}

func GetUpdates(environment string) []Update {

	resp := APICall(environment, "/api/v1/update", "")

	if resp.Code == 200 {
		//fmt.Println(resp.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(resp.Body), &result)

		var retUpdates []Update

		// The object stored in the "birds" key is also stored as
		// a map[string]interface{} type, and its type is asserted from
		// the interface{} type

		for updateId, updateBody := range result {
			// Each value is an interface{} type, that is type asserted as a string

			task := updateBody.(map[string]interface{})

			id, _ := strconv.Atoi(updateId)
			// These are returned as strings for some reason by Acquia
			added, _ := strconv.Atoi(task["added"].(string))
			status, _ := strconv.Atoi(task["status"].(string))

			newUpdate := Update{
				UpdateID: id,
				Added:    added,
				Status:   status,
			}
			retUpdates = append(retUpdates, newUpdate)
		}

		return retUpdates
	} else {
		fmt.Println(resp.Body)
	}

	return []Update{}
}

func GetUpdateStatus(environment string, updateId int) UpdateStatus {
	resp := APICall(environment, fmt.Sprintf("/api/v1/update/%d/status", updateId), "")

	if resp.Code == 200 {
		//fmt.Println(resp.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(resp.Body), &result)

		starttime, _ := strconv.Atoi(result["start_time"].(string))

		var endtime int
		if _, val := result["end_time"]; val {
			time, _ := strconv.Atoi(result["end_time"].(string))
			endtime = time
		}

		return UpdateStatus{
			UpdateID:   int(result["id"].(float64)),
			Message:    result["message"].(string),
			Percentage: int(result["percentage"].(float64)),
			StartTime:  starttime,
			EndTime:    endtime,
		}
	} else {
		fmt.Println(resp.Body)
	}

	return UpdateStatus{}
}
