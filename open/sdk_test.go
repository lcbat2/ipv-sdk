package sdk

import (
	"testing"
	"time"

	"github.com/clpublic/ipv-sdk/dto"
)

func getClient() *IpvClient {
	//根据实际需要 以下参数需要修改配置
	//Endpoint := "https://api.ipipv.com"

	Endpoint := "https://sandbox.ipipv.com"         //测试环境
	appKey := "AK20240708150154"                    //your AppKey
	appSecret := `xx1vsogzzvqvrgml7f6sc233p8vdqev2` //your appSecret
	client, err := NewClient(Endpoint, appKey, appSecret, Encrypt_AES)
	if err != nil {
		panic(err)
	}
	return client
}

func TestGetAppInfo(t *testing.T) {
	ps, err := getClient().GetAppInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestGetProductStock(t *testing.T) {
	ps, err := getClient().GetProductStock(dto.AppProductSyncReq{
		ProxyType: []int{103},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestCreateUser(t *testing.T) {
	ps, err := getClient().CreateUser(dto.AppUserReq{AuthName: "", No: "", AppUsername: "ad67a4bb7e1e4486bba8bc77027295b8", Password: ""})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestCreateUser2(t *testing.T) {
	ps, err := getClient().CreateUser(dto.AppUserReq{AuthName: "", No: "", AppUsername: "aaaabbb", Password: "bbbbbb"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestUserAuth(t *testing.T) {
	ps, err := getClient().UserAuth(dto.AppAuthUserReq{AuthName: "aaaa", No: "ccc", Username: "aaaabbb"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestGetOrder(t *testing.T) {
	orderNo := "C20240510134623046203"
	ps, err := getClient().GetOrder(dto.AppGetOrderReq{OrderNo: orderNo, Page: 1, PageSize: 10})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestGetInstance(t *testing.T) {
	instances := []string{"c_gzmtux84eu8ipkq"}
	ps, err := getClient().GetInstance(dto.AppGetInstanceReq{instances})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestGetArea(t *testing.T) {
	ps, err := getClient().GetArea(dto.AppGetAreaReq{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func getOrderNo() string {
	return "SDK_TEST" + time.Now().Format("20060102150405")
}

func TestInstanceOpen(t *testing.T) {
	params := []dto.OpenParam{
		dto.OpenParam{
			ProductNo: "ipideash_598", // tx_166 ip90_1189 aws_light_205  jg_226=随机地区  jg_227=安徽省芜湖市 ipidea_590=纽约 ipidea_d_283=美国
			Count:     20,
			Duration:  12,
		},
		dto.OpenParam{
			ProductNo: "mb_gmhd5exp2", // tx_166 ip90_1189 aws_light_205  jg_226=随机地区  jg_227=安徽省芜湖市 ipidea_590=纽约 ipidea_d_283=美国
			Count:     20,
			Duration:  365,
		},
	}
	o := dto.AppInstanceOpenReq{
		AppOrderNo: getOrderNo(),
		Params:     params,
	}
	ps, err := getClient().InstanceOpen(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestInstanceRenew(t *testing.T) {
	// ipidea => c_gzm9tradpjmqwg4
	o := dto.AppInstanceRenewReq{
		AppOrderNo: getOrderNo(),
		Instances:  []dto.Instance{dto.Instance{InstanceNo: "c_gzqk6ykpa2qb59o"}}, // tx=>c_gzmtux84eu8ipkq  ip90=>c_gzmk5maz55cv536  aws=>c_gzmk7asab4wy6vp jg => c_gzms93ztk8jnh0r
	}
	ps, err := getClient().InstanceRenew(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestInstanceRelease(t *testing.T) {
	o := dto.AppInstanceReleaseReq{
		OrderNo:   getOrderNo(),
		Instances: []string{"c_gzraaycyuhps5pg"}, // c_gzmycfycqejct4m  ip90=>c_gzmk5maz55cv536 aws=>c_gzmk7asab4wy6vp
	}
	ps, err := getClient().InstanceRelease(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestDrawByPwd(t *testing.T) {
	o := dto.AppDrawByPwdReq{
		AppUsername:  "app6_user1",
		AddressCode:  "USA",
		SessTime:     "",
		Num:          1,
		ProxyType:    104,
		MaxFlowLimit: 10,
	}
	ps, err := getClient().DrawByPwd(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestProxyInfo(t *testing.T) {
	o := dto.AppProxyInfoReq{
		Username:  "app6_user",
		ProxyType: 104,
	}
	ps, err := getClient().ProxyInfo(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_ProxyFlowUseLog(t *testing.T) {
	var req dto.AppFlowUseLogReq
	req.AppUsername = "app6_user"
	req.StartTime = "2024-06-01"
	req.EndTime = "2024-08-01"

	ps, err := getClient().ProxyFlowUseLog(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestProductAreaList(t *testing.T) {
	req := dto.AppProductAreaReq{
		ProductNo: "out_dynamic_1",
	}
	ps, err := getClient().ProductAreaList(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_GetCityList(t *testing.T) {
	req := dto.AppCityListReq{
		Codes: nil,
	}
	ps, err := getClient().GetCityList(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_DynamicProxyReturn(t *testing.T) {
	req := dto.AppProxyReturnReq{
		AppUsername: "",
		ProxyType:   104,
		ProductNo:   "out_dynamic_1",
		IpNum:       0,
		FlowNum:     1,
		Remark:      "流量回收",
	}
	ps, err := getClient().DynamicProxyReturn(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_CreateProxyUser(t *testing.T) {
	req := dto.AppProxyUserReq{
		AppUsername:     "",
		Password:        "",
		LimitFlow:       0,
		MainUsername:    "",
		AppMainUsername: "",
		Remark:          "",
		Status:          1,
	}
	ps, err := getClient().CreateProxyUser(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_InstanceOpenAssignIp(t *testing.T) {
	o := dto.AppAssignIpInstanceOpenReq{
		AppOrderNo: getOrderNo(),
		AssignIp:   "",
		ProductNo:  "",
		CycleTimes: 1,
	}
	ps, err := getClient().InstanceOpenAssignIp(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_GetAssignIpInfo(t *testing.T) {
	o := dto.AppGetAssignIpInfoReq{
		Ip: "",
	}
	ps, err := getClient().GetAssignIpInfo(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_GetOrderList(t *testing.T) {
	o := dto.AppGetOrderListReq{
		StartTime: "2024-01-01",
		EndTime:   "",
		Page:      1,
		PageSize:  10,
	}
	ps, err := getClient().GetOrderList(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestIpvClient_GetInstanceList(t *testing.T) {
	o := dto.AppGetInstanceListReq{
		StartTime: "2024-01-01",
		EndTime:   "",
		Page:      1,
		PageSize:  10,
	}
	ps, err := getClient().GetInstanceList(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}
