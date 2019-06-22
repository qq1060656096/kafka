package protocol

import "fmt"

// RequestHeader 请求头部
type RequestHeader struct {
	// 请求长度
	Size int32
	// 请求类型
	ApiKey int16
	// api版本
	ApiVersion int16
	// 用户定义ID来关联服务器和客户机之间的请求,用户提供的整数值,将与响应一起传回
	CorrelationId int32
	// 客户端id
	ClientId string
}

func (r *RequestHeader) String() string {
	return fmt.Sprintf(
		"correlation id: %d, api key: %d, client: %s, size: %d",
		r.CorrelationID,
		r.APIKey,
		r.ClientID,
		r.Size,
	)
}

func (r *RequestHeader) String() string {
	return fmt.Sprintf("Size: %d, ApiKey: %d, ApiVersion: %d, CorrelationId:%d, ClientIdLength:%d, client_id:%d ",
		r.Size,
		r.ApiKey,
		r.ApiVersion,
		r.CorrelationId,
		r.ClientId,
	)
}