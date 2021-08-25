package oppo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAppDetail(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetAppDetailParams{
		PkgName: "cccc",
	}

	appDetail, _, err := client.GetAppDetail(params)
	if err != nil {
		t.Error(err)
	}

	content, err := json.Marshal(&appDetail)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(content))

}

func TestPublishVersion(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	appInfo, _, err := client.GetAppDetail(&GetAppDetailParams{PkgName: "ccccc"})
	if err != nil {
		t.Error(err)
	}

	var apkUrls []ApkUrl
	apkUrl := ApkUrl{
		Url: "https://storedl.nearme.com.cn/apk/202004/21/xxxx.apk",
		Md5: "258261f5b0f4843106968eaa4ac5240f",
	}
	apkUrls = append(apkUrls, apkUrl)

	params := &PublishVersionParams{
		PkgName:          appInfo.Data.PkgName,
		VersionCode:      appInfo.Data.VersionCode,
		ApkUrl:           apkUrls,
		AppName:          appInfo.Data.AppName,
		SecondCategoryId: appInfo.Data.SecondCategoryID,
		ThirdCategoryId:  appInfo.Data.ThirdCategoryID,
		Summary:          appInfo.Data.Summary,
		DetailDesc:       appInfo.Data.DetailDesc,
		UpdateDesc:       appInfo.Data.UpdateDesc,
		PrivacySourceUrl: appInfo.Data.PrivacySourceURL,
		IconUrl:          appInfo.Data.IconURL,
		PicUrl:           appInfo.Data.PicURL,
		OnlineType:       appInfo.Data.OnlineType,
		TestDesc:         appInfo.Data.TestDesc,
		CopyrightUrl:     appInfo.Data.CopyrightURL,
		BusinessUsername: appInfo.Data.BusinessUsername,
		BusinessEmail:    appInfo.Data.BusinessEmail,
		BusinessMobile:   appInfo.Data.BusinessMobile,
	}

	res, _, err := client.PublishVersion(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Data)

}

func TestGetUploadFileConfig(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	res, _, err := client.GetUploadFileConfig()
	if err != nil {
		t.Error(err)
	}

	content, err := json.Marshal(&res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(content))

}

func TestUploadFile(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	config, _, err := client.GetUploadFileConfig()
	if err != nil {
		t.Error(err)
	}

	params := &UploadFileParams{
		Type:      "apk",
		Sign:      config.Data.Sign,
		FilePath:  "/Users/yinming/Downloads/icon_512.apk",
		UploadURL: config.Data.UploadURL,
	}

	res, _, err := client.UploadFile(params)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res.Data.URL)

}

func TestHmacSha256(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}
	s := `access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAwMzMyNDYsInN1YiI6Ik9QUE8tT09QLU9QRU5BUEkiLCJuYmYiOjE2Mjk4NjA0NDYsImF1ZCI6Im9wcG8tb29wIiwiaWF0IjoxNjI5ODYwNDQ2LCJqdGkiOiJvb3Bfb3BlbmFwaV91c2luZzo5ZDBjNjI1ZjBjOThjMTM4NzlhMDNiOTNlMmIzMjZlOSIsInN0YXR1cyI6MSwiZGF0YSI6eyJjbGllbnRfaWQiOjIwNDQyMDA2MzEzNjk2OTMxMzV9fQ.D76vEkNjX0Z-_YHmw_nIMJbg5lVsyM7IZ9PukpXU8fM&apk_url=[{"cpu_code":0,"md5":"258261f5b0f4843106968eaa4ac5240f","url":"https://storedl.nearme.com.cn/apk/202004/21/215de68f0db7f2818a2aa2f4a7a390cc.apk"}]&app_name=标识云打印&business_email=dev@niimbot.com&business_mobile=18986169493&business_username=Liuxiong&copyright_url=https://cdopic0.heytapimage.com/img/202001/17/ea09700ed4208ed6b7bb8c189874e4e5.jpg,,&detail_desc=标识云打印（LabelCloud）是专业打印标价签的应用软件。支持自定义编辑模板、扫描打印、扫码取模、行业logo、标签分类、蓝牙连接等强大功能。软件支持多种配套硬件提供给用户使用，可以根据用户的需求提供私人订制的标签，以解决多种应用场景的打印需求。功能简单易上手，符合各种用户群体，用户能更快更好的打印出自己需要的标签。&icon_url=https://cdopic0.heytapimage.com/img/202001/17/76d2b8285604ee33f27b6c973ca36ac0.png&online_type=1&pic_url=https://cdopic0.heytapimage.com/img/202004/21/83cf11acca8913a606df9028a98acdd0.png,https://cdopic0.heytapimage.com/img/202004/21/4eae34747bda6007bdce986f81483ecc.png,https://cdopic0.heytapimage.com/img/202004/21/16e961d76999201fc9d04840fa3312a0.png,https://cdopic0.heytapimage.com/img/202004/21/c036d4f9347dd6be1b509a8c07adc434.png&pkg_name=com.gengcon.android.bsydy&privacy_source_url=http://www.niimbot.com&second_category_id=78&summary=这是一款很好用的测试app&test_desc=Username:18986169493,Password:whsf2018&third_category_id=6687&timestamp=1629860446&update_desc=支持日文、韩文、俄文；注销功能&version_code=7`
	sign := HmacSha256(s, client.clientSecret)
	fmt.Println(sign)
}
