# 123云盘-Go-SDK

基于 123 云盘开放平台 OpenAPI 文档封装的 Go SDK。

本项目按官方文档的功能大类进行模块化封装，提供统一的 `Client` 聚合入口，方便在业务项目中直接调用。

## 功能覆盖

- 文件管理：文件列表/详情、创建目录、移动/重命名、回收站、复制/批量复制、下载信息、保险箱
- 上传：V2 分片上传、sha1 秒传、单步上传（multipart）
- 分享管理：分享/付费分享 创建、列表、更新
- 离线下载：创建任务、查询进度
- 用户管理：获取 access_token（clientID/clientSecret）、OAuth2 code/refresh_token 换取 access_token、用户信息、佣金列表
- 直链：启用/禁用、获取直链、刷新缓存、直链日志、直链离线日志、IP 黑名单配置
- 图床：图片列表/详情、删除/移动、云盘复制到图床、图床离线迁移

不包含：
- 视频转码相关接口
- 上传 V1 相关接口（如你需要可再补充）

## 安装

```bash
go get github.com/Qialas/123pan-go-sdk@latest
```

建议通过 tag 固定版本：

```bash
go get github.com/Qialas/123pan-go-sdk@v0.1.0
```

## 快速开始

### 1. 使用 clientID / clientSecret 获取 access_token

```go
package main

import (
	"context"
	"fmt"

	pan123 "github.com/Qialas/123pan-go-sdk"
)

func main() {
	ctx := context.Background()

	c, err := pan123.NewClient()
	if err != nil {
		panic(err)
	}

	tok, err := c.GetAccessToken(ctx, "clientID", "clientSecret")
	if err != nil {
		panic(err)
	}
	c.SetAccessToken(tok.AccessToken)

	me, err := c.User.Info(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(me.UID, me.Nickname)
}
```

### 2. 文件列表（推荐接口 v2）

```go
lst, err := c.File.ListV2(ctx, pan123.ListV2Request{
	ParentFileID: 0,
	Limit:        100,
})
```

### 3. 下载链接

```go
info, err := c.File.DownloadInfo(ctx, 123456789)
_ = info.DownloadURL
```

## 上传说明（V2）

上传相关能力都在 `c.File` 下（文件管理大类），核心接口：

- 创建文件：`CreateFileV2` -> `POST /upload/v2/file/create`
- 上传分片：`UploadSlice` -> `POST {upload_domain}/upload/v2/file/slice`
- 上传完毕：`UploadComplete` -> `POST /upload/v2/file/upload_complete`
- 获取上传域名：`UploadDomains` -> `GET /upload/v2/file/domain`

### 分片上传（手动流程）

```go
create, err := c.File.CreateFileV2(ctx, pan123.CreateFileV2Request{
	ParentFileID: 0,
	Filename:     "a.bin",
	Etag:         "md5hex",
	Size:         123,
})

uploadDomains, err := c.File.UploadDomains(ctx)
uploadBase := uploadDomains[0]

err = c.File.UploadSliceBytes(ctx, uploadBase, create.PreuploadID, 1, "a.bin", []byte("..."))
complete, err := c.File.UploadComplete(ctx, create.PreuploadID)
_ = complete.FileID
```

### sha1 秒传

```go
res, err := c.File.Sha1Reuse(ctx, pan123.Sha1ReuseRequest{
	ParentFileID: 0,
	Filename:     "a.bin",
	Sha1:         "sha1hex",
	Size:         123,
})
_ = res.Reuse
_ = res.FileID
```

### 单步上传（小文件 multipart）

```go
domains, err := c.File.UploadDomains(ctx)
uploadBase := domains[0]

data, err := c.File.SingleUpload(ctx, uploadBase, pan123.SingleUploadRequest{
	ParentFileID: 0,
	Filename:     "a.bin",
	Etag:         "md5hex",
	Size:         123,
}, "a.bin", bytes.NewReader([]byte("...")))
_ = data.FileID
_ = data.Completed
```

## 错误处理

所有接口返回值遵循官方结构：`code/message/data/x-traceID`。

当 HTTP 状态码非 2xx，或业务 `code != 0` 时，会返回 `*pan123.APIError`：

```go
if err != nil {
	if apiErr, ok := err.(*pan123.APIError); ok {
		_ = apiErr.StatusCode
		_ = apiErr.Code
		_ = apiErr.Message
		_ = apiErr.TraceID
	}
}
```

## 模块入口（按大类）

根 Client 聚合字段如下：

- `c.File`：文件管理（含上传 v2、下载、保险箱等）-> [filemgmt](./filemgmt)
- `c.Share`：分享管理 -> [sharemgmt](./sharemgmt)
- `c.Offline`：离线下载 -> [offlinedl](./offlinedl)
- `c.User`：用户管理 -> [usermgmt](./usermgmt)
- `c.DirectLink`：直链（含 IP 黑名单配置）-> [directlinkmgmt](./directlinkmgmt)
- `c.Image`：图床 -> [imagebed](./imagebed)

## 目录结构

- [core](./core)：HTTP 客户端、统一响应/错误、通用请求封装
- [filemgmt](./filemgmt)：文件管理
- [sharemgmt](./sharemgmt)：分享管理
- [offlinedl](./offlinedl)：离线下载
- [usermgmt](./usermgmt)：用户管理
- [directlinkmgmt](./directlinkmgmt)：直链
- [imagebed](./imagebed)：图床
