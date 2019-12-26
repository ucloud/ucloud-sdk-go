package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateUser 创建租户
func CreateUser(email, password string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	createUserReq := ucloudstackClient.NewCreateUserRequest()
	createUserReq.UserEmail = ucloud.String(email)
	createUserReq.PassWord = ucloud.String(password)

	createUserResp, err := ucloudstackClient.CreateUser(createUserReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, createUserResp.Message)
		return
	}

	fmt.Printf("New User's ID: %d", createUserResp.UserID)
}

// DescribeUser 查询用户信息
func DescribeUser(userID int) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	describeUserReq := ucloudstackClient.NewDescribeUserRequest()
	describeUserReq.Limit = ucloud.Int(10)
	describeUserReq.Offset = ucloud.Int(0)
	describeUserReq.UserIDs = []int{userID}
	describeUserResp, err := ucloudstackClient.DescribeUser(describeUserReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, describeUserResp.Message)
		return
	}

	fmt.Printf("User's info: %v", describeUserResp.Infos)
}
