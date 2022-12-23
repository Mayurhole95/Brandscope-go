package utils

import (
	"encoding/json"
	"fmt"
)

func ReturnError(err error) {
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}
}

func SuccessMessage(status bool, message string, filepath string) (success Success) {

	success.Success = status
	success.Message = message
	success.Filepath = filepath
	return success

}

type Success struct {
	Success  bool   `json:"Success"`
	Message  string `json:"Message"`
	Filepath string `json:"Filepath"`
}

func Marshal(success Success) (successmessage string) {
	status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
	statusstring, err := json.Marshal(status)
	ReturnError(err)
	successmessage = string(statusstring)
	return successmessage
}
