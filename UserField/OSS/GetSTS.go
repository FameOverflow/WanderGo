package OSS

import (
	"encoding/json"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *sts20150401.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Sts
	config.Endpoint = tea.String("sts.cn-shanghai.aliyuncs.com")
	_result = &sts20150401.Client{}
	_result, _err = sts20150401.NewClient(config)
	return _result, _err
}

func getSTS() (string, error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	stsToken := ""
	client, err := CreateClient(tea.String("LTAI5t5rPhn5UbdoLLQncn8r"), tea.String("KGJn7jyzCUBbX7wIgCFpVhngQhY7JK"))
	if err != nil {
		return stsToken, err
	}

	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleArn:         tea.String("acs:ram::1463793512635270:role/sts"),
		RoleSessionName: tea.String("sts"),
		DurationSeconds: tea.Int64(3600),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		response, err := client.AssumeRoleWithOptions(assumeRoleRequest, runtime)
		if err != nil {
			return err
		}
		// 获取临时身份令牌
		stsToken = tea.StringValue(response.Body.Credentials.SecurityToken)
		fmt.Println("Security Token:", stsToken)
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err = util.AssertAsString(error.Message)
		if err != nil {
			return stsToken, err
		}
	}
	return stsToken, nil
}

func GetSTS(ctx *gin.Context) {
	stsToken, err := getSTS()
	if err != nil {
		ctx.JSON(200, gin.H{
			"err": err,
		})
	}
	ctx.JSON(200, gin.H{
		"sts": stsToken,
	})
}
