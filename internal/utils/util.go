package utils

import (
	"fmt"
	"time"
)

// ReturnTopicName Topic格式
func ReturnTopicName(sspSlotId int64, dspId string) string {
	if sspSlotId == 0 {
		return ""
	}
	if dspId == "" {
		return fmt.Sprintf("ssp:%d:%s", sspSlotId, Gen10MinKeyTime(time.Now()))
	}
	// topic 格式 ssp:sspSlotId:dspId
	return fmt.Sprintf("ssp:%d:%s:%s", sspSlotId, dspId, Gen10MinKeyTime(time.Now()))
}
