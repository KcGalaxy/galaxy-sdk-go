package main

import (
	"encoding/json"
	"fmt"
	"github.com/gexue/galaxy-sdk-go/ksc"
	"github.com/gexue/galaxy-sdk-go/ksc/utils"
	"github.com/gexue/galaxy-sdk-go/service/sts"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	domain := "example.com"
	//debug模式的话 打开这个开关
	svc := sts.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL:      true,
		UseInternal: true,
		Domain:      domain,
	})
	var resp *map[string]interface{}
	var err error

	assumeRole := make(map[string]interface{})
	assumeRole["RoleKrn"] = ""
	assumeRole["RoleSessionName"] = "Slb"
	assumeRole["DurationSeconds"] = 900
	assumeRole["Policy"] = ""

	resp, err = svc.AssumeRole(&assumeRole)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
