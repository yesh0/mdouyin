package main

import (
	"common/kitex_gen/douyin/rpc"
	"common/utils"
	"context"
	"message/internal/cql"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// Chat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Chat(ctx context.Context, req *rpc.DouyinMessageChatRequest) (resp *rpc.DouyinMessageChatResponse, err error) {
	resp = rpc.NewDouyinMessageChatResponse()
	resp.MessageList = cql.ListMessages(req.RequestUserId, req.ToUserId, 300)
	return
}

// Message implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Message(ctx context.Context, req *rpc.DouyinMessageActionRequest) (resp *rpc.DouyinMessageActionResponse, err error) {
	resp = rpc.NewDouyinMessageActionResponse()
	switch req.ActionType {
	case 1: // Send message
		resp.StatusCode = int32(cql.Send(req.RequestUserId, req.ToUserId, req.Content))
	default:
		resp.StatusCode = int32(utils.ErrorWrongParameter)
	}
	return
}

func (s *MessageServiceImpl) LatestMessages(ctx context.Context, req *rpc.LatestMessageRequest) (resp *rpc.LatestMessageResponse, err error) {
	resp = rpc.NewLatestMessageResponse()
	resp.Messages = cql.LatestMessages(ctx, req.RequestUserId, req.Friends)
	return
}
