package main

import (
	"common/kitex_gen/douyin/rpc"
	"context"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// Chat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Chat(ctx context.Context, req *rpc.DouyinMessageChatRequest) (resp *rpc.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// Message implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Message(ctx context.Context, req *rpc.DouyinMessageActionRequest) (resp *rpc.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
