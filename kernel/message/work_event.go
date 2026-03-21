package message

import "encoding/xml"

// WorkContactEvent 企业微信通讯录变更事件消息
type WorkContactEvent struct {
	ChangeType string

	// 成员变更事件字段
	UserID         string
	NewUserID      string
	Name           string
	Department     string
	MainDepartment string
	IsLeaderInDept string
	DirectLeader   string
	Position       string
	Mobile         string
	Gender         string
	Email          string
	BizMail        string
	Status         string
	Avatar         string
	Alias          string
	Telephone      string
	Address        string
	ExtAttr        *WorkContactExtAttr

	// 部门变更事件字段
	Id       string
	ParentId string
	Order    string

	// 标签变更事件字段
	TagId         string
	AddUserItems  string
	DelUserItems  string
	AddPartyItems string
	DelPartyItems string
}

// WorkContactExtAttr 成员扩展属性
type WorkContactExtAttr struct {
	Item []WorkContactExtAttrItem
}

// WorkContactExtAttrItem 扩展属性项
type WorkContactExtAttrItem struct {
	Name string
	Type string
	Text struct {
		Value string
	}
	Web struct {
		Title string
		Url   string
	}
}

// WorkBatchJobResultEvent 企业微信异步任务完成事件消息
type WorkBatchJobResultEvent struct {
	BatchJob struct {
		JobId   string
		JobType string
		ErrCode int
		ErrMsg  string
	}
}

// WorkExternalContactEvent 企业微信外部联系人变更事件消息
type WorkExternalContactEvent struct {
	ChangeType     string
	UserID         string
	ExternalUserID string
	State          string
	WelcomeCode    string
	Source         string
	FailReason     string
}

// WorkExternalChatEvent 企业微信客户群变更事件消息
type WorkExternalChatEvent struct {
	ChangeType   string
	ChatId       string
	UpdateDetail string
	JoinScene    int
	QuitScene    int
	MemChangeCnt int
}

// WorkExternalTagEvent 企业微信企业客户标签变更事件消息
type WorkExternalTagEvent struct {
	ChangeType string
	Id         string
	TagType    string
}

// WorkTemplateCardEvent 企业微信模板卡片事件消息
type WorkTemplateCardEvent struct {
	EventKey      string
	TaskId        string
	CardType      string
	ResponseCode  string
	SelectedItems struct {
		SelectedItem []struct {
			QuestionKey string
			OptionIds   struct {
				OptionId []string
			}
		}
	}
}

// WorkLivingStatusChangeEvent 企业微信直播事件消息
type WorkLivingStatusChangeEvent struct {
	LivingId string
	Status   int
}

// WorkApprovalEvent 企业微信审批事件消息
type WorkApprovalEvent struct {
	ApprovalInfo struct {
		SpNo       string
		SpName     string
		SpStatus   int
		TemplateId string
		ApplyTime  int64
		Applyer    struct {
			UserId string
			Party  string
		}
		SpRecord []struct {
			SpStatus     int
			ApproverAttr int
			Details      []struct {
				Approver struct {
					UserId string
				}
				Speech   string
				SpStatus int
				SpTime   int64
			}
		}
		Notifyer []struct {
			UserId string
		}
		StatuChangeEvent int
	}
}

// parseWorkContactEvent 从 RawBody 解析企业微信通讯录变更事件
func parseWorkContactEvent(rawBody []byte) *WorkContactEvent {
	if len(rawBody) == 0 {
		return &WorkContactEvent{}
	}
	var op struct {
		XMLName        xml.Name `xml:"xml"`
		ChangeType     string
		UserID         string
		NewUserID      string
		Name           string
		Department     string
		MainDepartment string
		IsLeaderInDept string
		DirectLeader   string
		Position       string
		Mobile         string
		Gender         string
		Email          string
		BizMail        string
		Status         string
		Avatar         string
		Alias          string
		Telephone      string
		Address        string
		ExtAttr        *WorkContactExtAttr
		Id             string
		ParentId       string
		Order          string
		TagId          string
		AddUserItems   string
		DelUserItems   string
		AddPartyItems  string
		DelPartyItems  string
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkContactEvent{}
	}
	return &WorkContactEvent{
		ChangeType:     op.ChangeType,
		UserID:         op.UserID,
		NewUserID:      op.NewUserID,
		Name:           op.Name,
		Department:     op.Department,
		MainDepartment: op.MainDepartment,
		IsLeaderInDept: op.IsLeaderInDept,
		DirectLeader:   op.DirectLeader,
		Position:       op.Position,
		Mobile:         op.Mobile,
		Gender:         op.Gender,
		Email:          op.Email,
		BizMail:        op.BizMail,
		Status:         op.Status,
		Avatar:         op.Avatar,
		Alias:          op.Alias,
		Telephone:      op.Telephone,
		Address:        op.Address,
		ExtAttr:        op.ExtAttr,
		Id:             op.Id,
		ParentId:       op.ParentId,
		Order:          op.Order,
		TagId:          op.TagId,
		AddUserItems:   op.AddUserItems,
		DelUserItems:   op.DelUserItems,
		AddPartyItems:  op.AddPartyItems,
		DelPartyItems:  op.DelPartyItems,
	}
}

// parseWorkBatchJobResultEvent 从 RawBody 解析企业微信异步任务完成事件
func parseWorkBatchJobResultEvent(rawBody []byte) *WorkBatchJobResultEvent {
	if len(rawBody) == 0 {
		return &WorkBatchJobResultEvent{}
	}
	var op struct {
		XMLName  xml.Name `xml:"xml"`
		BatchJob struct {
			JobId   string
			JobType string
			ErrCode int
			ErrMsg  string
		}
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkBatchJobResultEvent{}
	}
	return &WorkBatchJobResultEvent{
		BatchJob: op.BatchJob,
	}
}

// parseWorkExternalContactEvent 从 RawBody 解析企业微信外部联系人变更事件
func parseWorkExternalContactEvent(rawBody []byte) *WorkExternalContactEvent {
	if len(rawBody) == 0 {
		return &WorkExternalContactEvent{}
	}
	var op struct {
		XMLName        xml.Name `xml:"xml"`
		ChangeType     string
		UserID         string
		ExternalUserID string
		State          string
		WelcomeCode    string
		Source         string
		FailReason     string
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkExternalContactEvent{}
	}
	return &WorkExternalContactEvent{
		ChangeType:     op.ChangeType,
		UserID:         op.UserID,
		ExternalUserID: op.ExternalUserID,
		State:          op.State,
		WelcomeCode:    op.WelcomeCode,
		Source:         op.Source,
		FailReason:     op.FailReason,
	}
}

// parseWorkExternalChatEvent 从 RawBody 解析企业微信客户群变更事件
func parseWorkExternalChatEvent(rawBody []byte) *WorkExternalChatEvent {
	if len(rawBody) == 0 {
		return &WorkExternalChatEvent{}
	}
	var op struct {
		XMLName      xml.Name `xml:"xml"`
		ChangeType   string
		ChatId       string
		UpdateDetail string
		JoinScene    int
		QuitScene    int
		MemChangeCnt int
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkExternalChatEvent{}
	}
	return &WorkExternalChatEvent{
		ChangeType:   op.ChangeType,
		ChatId:       op.ChatId,
		UpdateDetail: op.UpdateDetail,
		JoinScene:    op.JoinScene,
		QuitScene:    op.QuitScene,
		MemChangeCnt: op.MemChangeCnt,
	}
}

// parseWorkExternalTagEvent 从 RawBody 解析企业微信企业客户标签变更事件
func parseWorkExternalTagEvent(rawBody []byte) *WorkExternalTagEvent {
	if len(rawBody) == 0 {
		return &WorkExternalTagEvent{}
	}
	var op struct {
		XMLName    xml.Name `xml:"xml"`
		ChangeType string
		Id         string
		TagType    string
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkExternalTagEvent{}
	}
	return &WorkExternalTagEvent{
		ChangeType: op.ChangeType,
		Id:         op.Id,
		TagType:    op.TagType,
	}
}

// parseWorkTemplateCardEvent 从 RawBody 解析企业微信模板卡片事件
func parseWorkTemplateCardEvent(rawBody []byte) *WorkTemplateCardEvent {
	if len(rawBody) == 0 {
		return &WorkTemplateCardEvent{}
	}
	var op WorkTemplateCardEvent
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkTemplateCardEvent{}
	}
	return &op
}

// parseWorkLivingStatusChangeEvent 从 RawBody 解析企业微信直播事件
func parseWorkLivingStatusChangeEvent(rawBody []byte) *WorkLivingStatusChangeEvent {
	if len(rawBody) == 0 {
		return &WorkLivingStatusChangeEvent{}
	}
	var op struct {
		XMLName  xml.Name `xml:"xml"`
		LivingId string
		Status   int
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkLivingStatusChangeEvent{}
	}
	return &WorkLivingStatusChangeEvent{
		LivingId: op.LivingId,
		Status:   op.Status,
	}
}

// parseWorkApprovalEvent 从 RawBody 解析企业微信审批事件
func parseWorkApprovalEvent(rawBody []byte) *WorkApprovalEvent {
	if len(rawBody) == 0 {
		return &WorkApprovalEvent{}
	}
	var op WorkApprovalEvent
	if xml.Unmarshal(rawBody, &op) != nil {
		return &WorkApprovalEvent{}
	}
	return &op
}
