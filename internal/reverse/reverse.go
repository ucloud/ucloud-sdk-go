// this helper function to reverse the sdk services to json
package main

import (
	"encoding/json"
	"github.com/ucloud/ucloud-sdk-go/services/stepflow"
	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/ucloud/ucloud-sdk-go/services/ipsecvpn"
	"github.com/ucloud/ucloud-sdk-go/services/pathx"
	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/ubill"
	"github.com/ucloud/ucloud-sdk-go/services/udb"
	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/services/udpn"
	"github.com/ucloud/ucloud-sdk-go/services/ufile"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/services/umem"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/uphost"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"
)

type ApisReport struct {
	Context Context
	Apis    []ApiReport
}

type Context struct {
	Lang     string
	IsDryRun bool
}

type ApiReport struct {
	Name          string
	Product       string
	Request       []Attribute
	Response      []Attribute
	RequestModel  []Model
	ResponseModel []Model
}

type Attribute struct {
	Name     string
	Required bool
	Type     string
	ArrType  string
}

type Model struct {
	Name    string
	Product string
	Params  []Attribute
}

func main() {
	data, err := json.Marshal(initialize())
	if err != nil {
		log.Warn("cannot marshal report to json")
	}

	err = ioutil.WriteFile("../../build/sdk.json", data, 0755)
	if err != nil {
		log.Warn("cannot save report to local disk")
	}
}

func initialize() ApisReport {

	var apisReport ApisReport
	apisReport.Context = Context{
		Lang:     "go",
		IsDryRun: false,
	}

	var uhostClient *uhost.UHostClient
	var unetClient *unet.UNetClient
	var ulbClient *ulb.ULBClient
	var vpcClient *vpc.VPCClient
	var uaccountClient *uaccount.UAccountClient
	var pathxClient *pathx.PathXClient
	var udiskClient *udisk.UDiskClient
	var udbClient *udb.UDBClient
	var umemClient *umem.UMemClient
	var udpnClient *udpn.UDPNClient
	var ubillClient *ubill.UBillClient
	var uphostClient *uphost.UPHostClient
	var ipsecvpnClient *ipsecvpn.IPSecVPNClient
	var ufileClient *ufile.UFileClient
	var stepflowClient *stepflow.StepFlowClient
	var ucloudstackClient *ucloudstack.UCloudStackClient

	cfg := ucloud.NewConfig()
	cfg.Region = "..."
	cfg.ProjectId = "..."
	credential := auth.NewCredential()
	credential.PrivateKey = "..."
	credential.PublicKey = "..."

	uhostClient = uhost.NewClient(&cfg, &credential)
	unetClient = unet.NewClient(&cfg, &credential)
	ulbClient = ulb.NewClient(&cfg, &credential)
	vpcClient = vpc.NewClient(&cfg, &credential)
	uaccountClient = uaccount.NewClient(&cfg, &credential)
	pathxClient = pathx.NewClient(&cfg, &credential)
	udiskClient = udisk.NewClient(&cfg, &credential)
	udbClient = udb.NewClient(&cfg, &credential)
	umemClient = umem.NewClient(&cfg, &credential)
	udpnClient = udpn.NewClient(&cfg, &credential)
	ubillClient = ubill.NewClient(&cfg, &credential)
	uphostClient = uphost.NewClient(&cfg, &credential)
	ipsecvpnClient = ipsecvpn.NewClient(&cfg, &credential)
	ufileClient = ufile.NewClient(&cfg, &credential)
	stepflowClient = stepflow.NewClient(&cfg, &credential)
	ucloudstackClient = ucloudstack.NewClient(&cfg, &credential)

	apisReport.inspectProductClient(uhostClient)
	apisReport.inspectProductClient(unetClient)
	apisReport.inspectProductClient(ulbClient)
	apisReport.inspectProductClient(vpcClient)
	apisReport.inspectProductClient(uaccountClient)
	apisReport.inspectProductClient(pathxClient)
	apisReport.inspectProductClient(udiskClient)
	apisReport.inspectProductClient(udbClient)
	apisReport.inspectProductClient(umemClient)
	apisReport.inspectProductClient(udpnClient)
	apisReport.inspectProductClient(ubillClient)
	apisReport.inspectProductClient(uphostClient)
	apisReport.inspectProductClient(ipsecvpnClient)
	apisReport.inspectProductClient(ufileClient)
	apisReport.inspectProductClient(stepflowClient)
	apisReport.inspectProductClient(ucloudstackClient)

	return apisReport
}

var invalidMethod = []string{"InvokeAction", "GenericInvoke", "SetHttpClient", "GetCredential", "GetConfig", "GetCredential", "GetMeta", "SetLogger", "GetLogger", "SetupRequest", "InvokeActionWithPatcher", "AddHttpRequestHandler", "AddRequestHandler", "AddHttpResponseHandler", "AddResponseHandler", "WaitUntilUHostInstanceState"}

type metaClient interface {
	GetMeta() ucloud.ClientMeta
}

func (apisReport *ApisReport) inspectProductClient(client interface{}) {

	rv := reflect.TypeOf(client)
	for i := 0; i < rv.NumMethod(); i++ {
		method := rv.Method(i)
		name := method.Name
		if !strings.HasPrefix(name, "New") && !isStringIn(name, invalidMethod) {
			apisReport.inspectProductApi(method.Type, name, client.(metaClient).GetMeta().Product)
		}
	}
}

func (apisReport *ApisReport) inspectProductApi(api reflect.Type, apiName, productName string) {
	var apiReport ApiReport

	apiReport.Product = productName
	apiReport.Name = apiName

	// Request
	input := api.In(1)
	for input.Kind() == reflect.Ptr || input.Kind() == reflect.Slice {
		input = input.Elem()
	}
	// terminate polling
	if input.Kind().String() != "struct" {
		return
	}
	for i := 0; i < input.NumField(); i++ {
		f := input.Field(i)
		name := f.Name
		if name == "CommonBase" {
			continue
		}
		var apiType, arrType string
		var required bool
		if len(string(f.Tag)) != 0 {
			s := strings.Split(string(f.Tag), ":")
			required, _ = strconv.ParseBool(strings.Trim(s[1], "\""))
		}

		if strings.HasPrefix(f.Type.String(), "[]") {
			apiType = "array"
			arrType = strings.TrimPrefix(f.Type.String(), "[]")
			if strings.Contains(arrType, ".") {
				n := strings.Split(arrType, ".")
				arrType = n[1]
			}
		} else if strings.HasPrefix(f.Type.String(), "*") {
			apiType = strings.TrimPrefix(f.Type.String(), "*")
			arrType = ""
		}
		apiReport.Request = append(apiReport.Request, Attribute{
			Name:     name,
			Required: required,
			Type:     apiType,
			ArrType:  arrType,
		})

		// RequestModel
		apiReport.encodeModel(f, true)

	}

	// Response
	output := api.Out(0)
	for output.Kind() == reflect.Ptr || output.Kind() == reflect.Slice {
		output = output.Elem()
	}
	// terminate polling
	if output.Kind().String() != "struct" {
		return
	}

	for i := 0; i < output.NumField(); i++ {
		f := output.Field(i)
		name := f.Name
		if name == "CommonBase" {
			continue
		}
		var apiType, arrType string
		if strings.HasPrefix(f.Type.String(), "[]") {
			apiType = "array"
			arrType = strings.TrimPrefix(f.Type.String(), "[]")
			if strings.Contains(arrType, ".") {
				n := strings.Split(arrType, ".")
				arrType = n[1]
			}
		} else {
			apiType = strings.TrimPrefix(f.Type.String(), "*")
			arrType = ""
		}
		apiReport.Response = append(apiReport.Response, Attribute{
			Name:    name,
			Type:    apiType,
			ArrType: arrType,
		})

		// ResponseModel
		apiReport.encodeModel(f, false)
	}

	apisReport.Apis = append(apisReport.Apis, apiReport)
}

func (apiReport *ApiReport) encodeModel(f reflect.StructField, flag bool) {
	var apiType, arrType string
	var required bool
	var params []Attribute

	model := f.Type
	for model.Kind() == reflect.Ptr || model.Kind() == reflect.Slice {
		model = model.Elem()
	}
	// terminate polling
	if model.Kind().String() != "struct" {
		return
	}

	for i := 0; i < model.NumField(); i++ {
		f = model.Field(i)
		name := f.Name

		if len(string(f.Tag)) != 0 {
			s := strings.Split(string(f.Tag), ":")
			required, _ = strconv.ParseBool(strings.Trim(s[1], "\""))
		}
		if strings.HasPrefix(f.Type.String(), "[]") {
			apiType = "array"
			arrType = strings.TrimPrefix(f.Type.String(), "[]")
			if strings.Contains(arrType, ".") {
				n := strings.Split(arrType, ".")
				arrType = n[1]
			}
		} else {
			apiType = strings.TrimPrefix(f.Type.String(), "*")
			arrType = ""
			if strings.Contains(apiType, ".") {
				n := strings.Split(apiType, ".")
				apiType = n[1]
			}
		}
		// polling struct
		apiReport.encodeModel(f, flag)

		params = append(params, Attribute{
			Name:     name,
			Required: required,
			Type:     apiType,
			ArrType:  arrType,
		})
	}

	if flag {
		apiReport.RequestModel = append(apiReport.RequestModel, Model{
			Name:    model.Name(),
			Product: apiReport.Product,
			Params:  params,
		})
	} else {
		apiReport.ResponseModel = append(apiReport.ResponseModel, Model{
			Name:    model.Name(),
			Product: apiReport.Product,
			Params:  params,
		})
	}
}

func isStringIn(val string, available []string) bool {
	for _, choice := range available {
		if val == choice {
			return true
		}
	}

	return false
}
