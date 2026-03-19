# 企业微信SDK基础模块实现计划

## Context

当前 `work/` 包已有基础架构（config、option、access_token、AccountInterface 实现），以及 `mini_program` 和 `server` 子模块。需要按照用户规划的三大分类（base/connect/office）组织企业微信API，本次先完成 `base` 下的3个子模块：账号ID、通讯录管理（成员管理）、身份验证（网页授权登录）。

## 目标目录结构

```
work/
├── work.go              (修改: 新增 Base() 工厂方法)
├── base/                (新建: 基础模块入口)
│   ├── base.go          (Base 结构体, 暴露 AccountId()/Contact()/OAuth())
│   ├── account_id/      (账号ID)
│   │   ├── account_id.go    (userid↔openid互换)
│   │   └── convert.go       (自建应用与第三方应用对接、tmp_external_userid转换)
│   ├── contact/         (通讯录管理-成员管理)
│   │   ├── user.go          (User: 创建/读取/更新/删除成员)
│   │   └── types.go         (请求/响应数据结构)
│   └── oauth/           (身份验证-网页授权登录)
│       ├── oauth.go         (构造授权链接、重定向、获取用户身份)
│       └── types.go         (响应数据结构)
```

## 架构模式（严格遵循现有模式）

所有子模块遵循与 `official/user/`、`official/oauth/` 相同的模式：
- 接收 `contracts.AccountInterface` 作为依赖
- 使用 `work/http` 包发起 HTTP 请求（baseUri = `https://qyapi.weixin.qq.com/`）
- 使用 `kernelError.ApiError` 解析错误响应
- 通过 `account.AccessToken()` 获取 access_token

关键引用文件:
- `work/http/http.go` — 企业微信HTTP客户端
- `kernel/contracts/account.go` — AccountInterface 接口
- `kernel/error/error.go` — ApiError / kernelError
- `official/user/user.go` — 成员管理参考模式
- `official/oauth/oauth.go` — OAuth参考模式

## 实现详情

### 1. `work/base/base.go` — 基础模块入口

```go
package base

import (
    "github.com/dysodeng/wx/kernel/contracts"
    "github.com/dysodeng/wx/work/base/account_id"
    "github.com/dysodeng/wx/work/base/contact"
    "github.com/dysodeng/wx/work/base/oauth"
)

type Base struct {
    account contracts.AccountInterface
}

func New(account contracts.AccountInterface) *Base {
    return &Base{account: account}
}

func (b *Base) AccountId() *account_id.AccountId {
    return account_id.New(b.account)
}

func (b *Base) Contact() *contact.User {
    return contact.NewUser(b.account)
}

func (b *Base) OAuth() *oauth.OAuth {
    return oauth.New(b.account)
}
```

### 2. `work/base/account_id/` — 账号ID模块

**account_id.go** — userid与openid互换:
- `ConvertToOpenid(userid string) (string, error)` — POST `/cgi-bin/user/convert_to_openid`
- `ConvertToUserid(openid string) (string, error)` — POST `/cgi-bin/user/convert_to_userid`

**convert.go** — 自建应用与第三方应用对接 + tmp_external_userid转换:
- `UseridToOpenuserid(useridList []string) (*OpenUseridResult, error)` — POST `/cgi-bin/batch/userid_to_openuserid`
- `OpenuseridToUserid(openUseridList []string, sourceAgentId int) (*UseridResult, error)` — POST `/cgi-bin/batch/openuserid_to_userid`
- `ConvertTmpExternalUserid(tmpExternalUseridList []string, businessType int, userType int) (*TmpExternalUseridResult, error)` — POST `/cgi-bin/idconvert/convert_tmp_external_userid`

### 3. `work/base/contact/` — 通讯录管理（成员管理）

**user.go** — 成员CRUD:
- `Create(user CreateUserRequest) error` — POST `/cgi-bin/user/create`
- `Get(userid string) (*UserInfo, error)` — GET `/cgi-bin/user/get`
- `Update(user UpdateUserRequest) error` — POST `/cgi-bin/user/update`
- `Delete(userid string) error` — GET `/cgi-bin/user/delete`
- `BatchDelete(useridList []string) error` — POST `/cgi-bin/user/batchdelete`
- `ListId(cursor string, limit int) (*UserIdList, error)` — POST `/cgi-bin/user/list_id`

**types.go** — 数据结构:
- `CreateUserRequest` — 创建成员请求体 (userid, name, department, mobile, email 等)
- `UpdateUserRequest` — 更新成员请求体
- `UserInfo` — 成员详细信息
- `UserIdList` — 成员ID列表响应

### 4. `work/base/oauth/` — 身份验证（网页授权登录）

**oauth.go** — OAuth网页授权:
- `WithScope(scope string) *OAuth` — 设置授权作用域 (snsapi_base/snsapi_userinfo/snsapi_privateinfo)
- `WithRedirectUrl(redirectUrl string) *OAuth` — 设置回调URL
- `WithState(state string) *OAuth` — 设置state参数
- `WithAgentId(agentId string) *OAuth` — 设置agentid (snsapi_privateinfo时必填)
- `AuthUrl() string` — 构造授权链接
- `Redirect(writer, request)` — 302重定向到授权页
- `UserFromCode(code string) (*UserIdentity, error)` — GET `/cgi-bin/auth/getuserinfo` 获取用户身份
- `GetUserDetail(userTicket string) (*UserDetail, error)` — POST `/cgi-bin/auth/getuserdetail` 获取敏感信息

授权链接格式:
```
https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&agentid=AGENTID&state=STATE#wechat_redirect
```

**types.go** — 响应结构:
- `UserIdentity` — {UserId, OpenId, DeviceId, UserTicket, ExpiresIn}
- `UserDetail` — {UserId, Gender, Avatar, QrCode, Mobile, Email, BizMail, Address}

### 5. 修改 `work/work.go` — 新增 Base() 入口

```go
import "github.com/dysodeng/wx/work/base"

func (w *Work) Base() *base.Base {
    return base.New(w)
}
```

## 实施顺序

1. 创建 `work/base/base.go`
2. 创建 `work/base/account_id/account_id.go` + `convert.go`
3. 创建 `work/base/contact/types.go` + `user.go`
4. 创建 `work/base/oauth/types.go` + `oauth.go`
5. 修改 `work/work.go` 添加 `Base()` 方法
6. `go build ./...` 验证编译通过

## 验证方式

```bash
cd /Users/dysodeng/project/go/wx && go build ./...
```

## API文档参考

- 账号ID概述: https://developer.work.weixin.qq.com/document/path/98728
- 自建应用与第三方应用对接: https://developer.work.weixin.qq.com/document/path/95884
- tmp_external_userid转换: https://developer.work.weixin.qq.com/document/path/98729
- 通讯录管理概述: https://developer.work.weixin.qq.com/document/path/90193
- 获取成员ID列表: https://developer.work.weixin.qq.com/document/path/100067
- 创建成员: https://developer.work.weixin.qq.com/document/path/90195
- 读取成员: https://developer.work.weixin.qq.com/document/path/90196
- 更新成员: https://developer.work.weixin.qq.com/document/path/90197
- 删除成员: https://developer.work.weixin.qq.com/document/path/90198
- 构造网页授权链接: https://developer.work.weixin.qq.com/document/path/91022
- 获取访问用户身份: https://developer.work.weixin.qq.com/document/path/91023
- 获取访问用户敏感信息: https://developer.work.weixin.qq.com/document/path/95833
