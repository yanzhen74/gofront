package routes

import (
	"testing"

	"github.com/kirinlabs/HttpRequest"
)

func Test_Protocal_Process_state_Post(t *testing.T) {
	var jsonStr = `
	{
		"msgType": "ProtocalReport",
		"PID": 1234,
		"MID": "HXC",
		"BID": "00112233",
		"ID": 1,
		"MainOrBackup": 1,
		"ProcessName": "LINK_CTCC - TL1A1_POAC",
		"Report": {
			"Report_type": "CP_REPORT_ON_TIMER",
			"Command_type": "bb",
			"Command_result": "bb",
			"Recv_status_revert": false,
			"Recv_status": true,
			"First": "yyyyMMddHHmmss",
			"Last": "yyyyMMddHHmmss",
			"Recv_count": 100,
			"Send_no": 100
		}
	}
	`
	HttpRequest.JSON().Post("http://localhost:8080/protocal/process_state", jsonStr)
}
