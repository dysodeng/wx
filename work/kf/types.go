package kf

import kernelError "github.com/dysodeng/wx/kernel/error"

// UpdateAccountRequest 修改客服账号请求
type UpdateAccountRequest struct {
	OpenKfid string `json:"open_kfid"`
	Name     string `json:"name,omitempty"`
	MediaId  string `json:"media_id,omitempty"`
}

// AccountInfo 客服账号信息
type AccountInfo struct {
	OpenKfid        string `json:"open_kfid"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	ManagePrivilege bool   `json:"manage_privilege"`
}

// AccountListResult 获取客服账号列表结果
type AccountListResult struct {
	AccountList []AccountInfo `json:"account_list"`
}

// ========== 接待人员相关类型 ==========

// ServicerResult 添加/删除接待人员结果项
type ServicerResult struct {
	Userid  string `json:"userid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ServicerInfo 接待人员信息
type ServicerInfo struct {
	Userid       string `json:"userid"`
	Status       int    `json:"status"`
	StopType     int    `json:"stop_type,omitempty"`
	DepartmentId int    `json:"department_id,omitempty"`
}

// ========== 内部响应类型 ==========

// addAccountResponse 添加客服账号响应
type addAccountResponse struct {
	kernelError.ApiError
	OpenKfid string `json:"open_kfid"`
}

// accountListResponse 获取客服账号列表响应
type accountListResponse struct {
	kernelError.ApiError
	AccountListResult
}

// addContactWayResponse 获取客服账号链接响应
type addContactWayResponse struct {
	kernelError.ApiError
	URL string `json:"url"`
}

// servicerAddDelResponse 添加/删除接待人员响应
type servicerAddDelResponse struct {
	kernelError.ApiError
	ResultList []ServicerResult `json:"result_list"`
}

// servicerListResponse 获取接待人员列表响应
type servicerListResponse struct {
	kernelError.ApiError
	ServicerList []ServicerInfo `json:"servicer_list"`
}
