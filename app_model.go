package oppo

type GetAppDetailParams struct {
	PkgName     string `json:"pkg_name"`
	VersionCode string `json:"version_code,omitempty"`
}

type GetAppDetailRes struct {
	Ret
	Data AppInfo `json:"data"`
}

type AppInfo struct {
	AppID                   string                   `json:"app_id"`
	PkgName                 string                   `json:"pkg_name"`
	Type                    int64                    `json:"type"`
	Sign                    string                   `json:"sign"`
	DevID                   string                   `json:"dev_id"`
	AppKey                  string                   `json:"app_key"`
	UpdateTime              string                   `json:"update_time"`
	AppCreateTime           string                   `json:"app_create_time"`
	AppName                 string                   `json:"app_name"`
	IsFreeze                string                   `json:"is_freeze"`
	FreezeReason            interface{}              `json:"freeze_reason"`
	RefuseReason            string                   `json:"refuse_reason"`
	TagList                 interface{}              `json:"tag_list"`
	IsBusiness              string                   `json:"is_business"`
	GameType                string                   `json:"game_type"`
	SecondCategoryID        string                   `json:"second_category_id"`
	ThirdCategoryID         string                   `json:"third_category_id"`
	CopyrightURL            string                   `json:"copyright_url"`
	SpecialURL              string                   `json:"special_url"`
	SpecialFileURL          string                   `json:"special_file_url"`
	FreezeFile              interface{}              `json:"freeze_file"`
	BusinessUsername        string                   `json:"business_username"`
	BusinessEmail           string                   `json:"business_email"`
	BusinessMobile          string                   `json:"business_mobile"`
	BusinessQq              string                   `json:"business_qq"`
	BusinessPosition        string                   `json:"business_position"`
	BusinessAddress         string                   `json:"business_address"`
	FreezeAdvice            interface{}              `json:"freeze_advice"`
	AppType                 string                   `json:"app_type"`
	AppRealType             string                   `json:"app_real_type"`
	ElectronicCERTURL       string                   `json:"electronic_cert_url"`
	ICPURL                  string                   `json:"icp_url"`
	RelationAppID           string                   `json:"relation_app_id"`
	VersionID               string                   `json:"version_id"`
	VersionCode             string                   `json:"version_code"`
	VersionName             string                   `json:"version_name"`
	ApkURL                  string                   `json:"apk_url"`
	ApkSize                 string                   `json:"apk_size"`
	ApkMd5                  string                   `json:"apk_md5"`
	HeaderMd5               string                   `json:"header_md5"`
	PackagePermission       string                   `json:"package_permission"`
	Resolution              interface{}              `json:"resolution"`
	VersionType             string                   `json:"version_type"`
	CreateTime              string                   `json:"create_time"`
	VerSecondCategoryID     string                   `json:"ver_second_category_id"`
	VerThirdCategoryID      string                   `json:"ver_third_category_id"`
	ApkFullURL              string                   `json:"apk_full_url"`
	OnlineType              string                   `json:"online_type"`
	ScheOnlineTime          string                   `json:"sche_online_time"`
	TestType                string                   `json:"test_type"`
	TestStartTime           string                   `json:"test_start_time"`
	TestEndTime             string                   `json:"test_end_time"`
	PlayerCustomerEmail     interface{}              `json:"player_customer_email"`
	PlayerCustomerPhone     interface{}              `json:"player_customer_phone"`
	PlayerCustomerQq        string                   `json:"player_customer_qq"`
	IsSignature             string                   `json:"is_signature"`
	IsPreDownload           string                   `json:"is_pre_download"`
	IconURL                 string                   `json:"icon_url"`
	IconMd5                 string                   `json:"icon_md5"`
	Summary                 string                   `json:"summary"`
	DetailDesc              string                   `json:"detail_desc"`
	UpdateDesc              string                   `json:"update_desc"`
	AppSubname              string                   `json:"app_subname"`
	TestDesc                string                   `json:"test_desc"`
	VideoURL                string                   `json:"video_url"`
	PicURL                  string                   `json:"pic_url"`
	PackagePermissionDesc   interface{}              `json:"package_permission_desc"`
	VideoPicURL             interface{}              `json:"video_pic_url"`
	CoverURL                interface{}              `json:"cover_url"`
	LandscapePicURL         string                   `json:"landscape_pic_url"`
	PrivacySourceURL        string                   `json:"privacy_source_url"`
	Level                   string                   `json:"level"`
	State                   string                   `json:"state"`
	OnlineTime              string                   `json:"online_time"`
	OfflineTime             string                   `json:"offline_time"`
	BusinessRefuseReason    string                   `json:"business_refuse_reason"`
	OnlineInfoOfflineApply  []OnlineInfoOfflineApply `json:"online_info_offline_apply"`
	Size                    string                   `json:"size"`
	AuditStatusName         string                   `json:"audit_status_name"`
	OfflineInfo             interface{}              `json:"offline_info"`
	TransferState           int64                    `json:"transfer_state"`
	UpdateInfoCheck         int64                    `json:"update_info_check"`
	LevelTag                string                   `json:"level_tag"`
	RefuseAdvice            string                   `json:"refuse_advice"`
	RefuseFile              string                   `json:"refuse_file"`
	LandscapePicURLMaterial []PicURLMaterial         `json:"landscape_pic_url_material"`
	PicURLMaterial          []PicURLMaterial         `json:"pic_url_material"`
	VideoURLMaterial        []interface{}            `json:"video_url_material"`
}

type PicURLMaterial struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Md5    string `json:"md5"`
	Size   string `json:"size"`
}

type OnlineInfoOfflineApply struct {
	AppID       string `json:"app_id"`
	VersionID   string `json:"version_id"`
	OnlineState string `json:"online_state"`
	OnlineTime  string `json:"online_time"`
	OfflineTime string `json:"offline_time"`
}

// OppoUpdateReq Oppo应用更新请求参数
type PublishVersionParams struct {
	PkgName          string   `json:"pkg_name"`
	VersionCode      string   `json:"version_code"`
	ApkUrl           []ApkUrl `json:"apk_url"`
	AppName          string   `json:"app_name"`
	SecondCategoryId string   `json:"second_category_id"`
	ThirdCategoryId  string   `json:"third_category_id"`
	Summary          string   `json:"summary"`
	DetailDesc       string   `json:"detail_desc"`
	UpdateDesc       string   `json:"update_desc"`
	PrivacySourceUrl string   `json:"privacy_source_url"`
	IconUrl          string   `json:"icon_url"`
	PicUrl           string   `json:"pic_url"`
	OnlineType       string   `json:"online_type"`
	TestDesc         string   `json:"test_desc"`
	CopyrightUrl     string   `json:"copyright_url"`
	BusinessUsername string   `json:"business_username"`
	BusinessEmail    string   `json:"business_email"`
	BusinessMobile   string   `json:"business_mobile"`
}

type ApkUrl struct {
	Url     string `json:"url"`
	Md5     string `json:"md5"`
	CpuCode int    `json:"cpu_code"`
}

type PublishVersionRes struct {
	Errno int64       `json:"errno"`
	Data  interface{} `json:"data"`
}

type GetUploadFileConfigRes struct {
	Ret
	Data PreUploadBody `json:"data"`
}

type PreUploadBody struct {
	UploadURL string `json:"upload_url"`
	Sign      string `json:"sign"`
}

type UploadFileParams struct {
	Type      string `json:"type"`
	Sign      string `json:"sign"`
	FilePath  string `json:"file_path"`
	UploadURL string `json:"upload_url"`
}

type UploadFileRes struct {
	Ret
	Data UploadObj `json:"data"`
}

type UploadObj struct {
	URL           string `json:"url"`
	URIPath       string `json:"uri_path"`
	Md5           string `json:"md5"`
	Sign          string `json:"sign"`
	Width         int64  `json:"width"`
	Height        int64  `json:"height"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
}
