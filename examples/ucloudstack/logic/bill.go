package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// Recharge 账户充值
func Recharge(userID int, serialNo string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	chargeReq := ucloudstackClient.NewRechargeRequest()
	chargeReq.Amount = ucloud.Int(100000)
	chargeReq.FromType = ucloud.String("INPOUR_FROM_ALIPAY")
	chargeReq.SerialNo = ucloud.String(serialNo)
	chargeReq.UserID = ucloud.Int(userID)

	chargeResp, err := ucloudstackClient.Recharge(chargeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, chargeResp.Message)
		return
	}

	fmt.Printf("Recharge success")
}
