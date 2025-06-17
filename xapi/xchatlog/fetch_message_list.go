package xchatlog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xurl"
)

// Message ...
type Message struct {
	Seq        int64     `json:"seq"`        // 序列
	Time       time.Time `json:"time"`       // 时间
	Talker     string    `json:"talker"`     // 会话ID
	TalkerName string    `json:"talkerName"` // 会话名称
	IsChatRoom bool      `json:"isChatRoom"` // 是否群聊
	Sender     string    `json:"sender"`     // 发送者
	SenderName string    `json:"senderName"` // 发送者名称
	IsSelf     bool      `json:"isSelf"`     // 是否自己
	Type       int       `json:"type"`       // 类型
	SubType    int       `json:"subType"`    // 类型
	Content    string    `json:"content"`    // 消息
	Contents   struct {
		Md5   string   `json:"md5"`
		Refer *Message `json:"refer"` // 引用
	} `json:"contents"` // 消息
}

// FetchMessageListReq ...
type FetchMessageListReq struct {
	Talker string `json:"talker"` // 会话
	Sender string `json:"sender"` // 发送人
	Time   string `json:"time"`   // 时间
	Limit  int    `json:"limit"`  // 数量
	Offset int    `json:"offset"` // 游标
	Format string `json:"format"` // 格式化
}

// FetchMessageList 查询消息列表
func FetchMessageList(ctx context.Context, req *FetchMessageListReq) ([]*Message, error) {
	req.Format = "json"
	response, err := FetchMessageListRaw(ctx, req)
	if err != nil {
		return nil, err
	}
	list := make([]*Message, 0)
	if err := json.Unmarshal([]byte(response), &list); err != nil {
		return nil, err
	}
	for k, v := range list {
		if v.TalkerName == "" {
			list[k].TalkerName = v.Talker
		}
		if v.SenderName == "" {
			list[k].SenderName = v.Sender
		}
	}
	return list, nil
}

// FetchMessageListRaw 查询消息列表
func FetchMessageListRaw(ctx context.Context, req *FetchMessageListReq) (string, error) {
	if req.Talker == "" {
		return "", fmt.Errorf("请指定会话名称")
	}
	url := fmt.Sprintf("%s/api/v1/chatlog", ChatlogAPI)
	query := make(map[string]interface{})
	query["talker"] = req.Talker
	if req.Sender != "" {
		query["sender"] = req.Sender
	}
	if req.Time != "" {
		query["time"] = req.Time
	}
	if req.Limit > 0 {
		query["limit"] = req.Limit
	}
	if req.Offset > 0 {
		query["offset"] = req.Offset
	}
	if req.Format == "" {
		req.Format = "text"
	}
	query["format"] = req.Format
	reqURL := xurl.BuildURL(url, query)
	response, err := xhttp.New().Get(ctx, reqURL, xhttp.RandomHeader())
	if err != nil {
		return "", err
	}
	xlog.Ctx(ctx).Debugf("FetchMessageListRaw url: %s, response: %s", reqURL, response.String())
	return response.String(), nil
}
