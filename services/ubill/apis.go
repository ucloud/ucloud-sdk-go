// Code is generated by ucloud-model, DO NOT EDIT IT.




package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UBill API Schema
















// DescribeOrderDetailInfoRequest is request schema for DescribeOrderDetailInfo action
type DescribeOrderDetailInfoRequest struct {
	request.CommonBase




    //  
    AzGroups []string `required:"false"`

    //  
    BeginTime *int `required:"true"`

    //  
    ChargeTypes []string `required:"false"`

    //  
    EndTime *int `required:"true"`

    //  
    Invoiceds []string `required:"false"`

    //  
    Limit *int `required:"false"`

    //  
    Offset *int `required:"false"`

    //  
    OrderStates []string `required:"false"`

    //  
    OrderTypes []string `required:"false"`

    //  
    QueryAll *string `required:"false"`

    //  
    Regions []string `required:"false"`

    //  
    ResourceIds []string `required:"false"`

    //  
    ResourceTypes []string `required:"false"`

    //  
    Tags []string `required:"false"`

    //  
    TradeNos []string `required:"false"`

}


// DescribeOrderDetailInfoResponse is response schema for DescribeOrderDetailInfo action
type DescribeOrderDetailInfoResponse struct {
	response.CommonBase


    //  
    OrderInfos []OrderInfo 

}


// NewDescribeOrderDetailInfoRequest will create request of DescribeOrderDetailInfo action.
func (c *UBillClient) NewDescribeOrderDetailInfoRequest() *DescribeOrderDetailInfoRequest {
    req := &DescribeOrderDetailInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}


/*
API: DescribeOrderDetailInfo


*/
func (c *UBillClient) DescribeOrderDetailInfo(req *DescribeOrderDetailInfoRequest) (*DescribeOrderDetailInfoResponse, error) {
	var err error
	var res DescribeOrderDetailInfoResponse

	reqCopier := *req
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
	err = c.Client.InvokeAction("DescribeOrderDetailInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}








// GetBillDataFileUrlRequest is request schema for GetBillDataFileUrl action
type GetBillDataFileUrlRequest struct {
	request.CommonBase




    //  
    BillPeriod *int `required:"true"`

    //  
    BillType *int `required:"true"`

    //  
    PaidType *int `required:"false"`

}


// GetBillDataFileUrlResponse is response schema for GetBillDataFileUrl action
type GetBillDataFileUrlResponse struct {
	response.CommonBase


    //  
    FileUrl string 

    //  
    IsValid string 

}


// NewGetBillDataFileUrlRequest will create request of GetBillDataFileUrl action.
func (c *UBillClient) NewGetBillDataFileUrlRequest() *GetBillDataFileUrlRequest {
    req := &GetBillDataFileUrlRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}


/*
API: GetBillDataFileUrl


*/
func (c *UBillClient) GetBillDataFileUrl(req *GetBillDataFileUrlRequest) (*GetBillDataFileUrlResponse, error) {
	var err error
	var res GetBillDataFileUrlResponse

	reqCopier := *req
    
    
    
    
    
    
    
	err = c.Client.InvokeAction("GetBillDataFileUrl", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}


