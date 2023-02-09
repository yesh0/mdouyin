// Code generated by Kitex v0.4.4. DO NOT EDIT.

package reactionservice

import (
	rpc "common/kitex_gen/douyin/rpc"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return reactionServiceServiceInfo
}

var reactionServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ReactionService"
	handlerType := (*rpc.ReactionService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Favorite":      kitex.NewMethodInfo(favoriteHandler, newReactionServiceFavoriteArgs, newReactionServiceFavoriteResult, false),
		"ListFavorites": kitex.NewMethodInfo(listFavoritesHandler, newReactionServiceListFavoritesArgs, newReactionServiceListFavoritesResult, false),
		"Comment":       kitex.NewMethodInfo(commentHandler, newReactionServiceCommentArgs, newReactionServiceCommentResult, false),
		"ListComments":  kitex.NewMethodInfo(listCommentsHandler, newReactionServiceListCommentsArgs, newReactionServiceListCommentsResult, false),
		"TestFavorites": kitex.NewMethodInfo(testFavoritesHandler, newReactionServiceTestFavoritesArgs, newReactionServiceTestFavoritesResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "rpc",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.ReactionServiceFavoriteArgs)
	realResult := result.(*rpc.ReactionServiceFavoriteResult)
	success, err := handler.(rpc.ReactionService).Favorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReactionServiceFavoriteArgs() interface{} {
	return rpc.NewReactionServiceFavoriteArgs()
}

func newReactionServiceFavoriteResult() interface{} {
	return rpc.NewReactionServiceFavoriteResult()
}

func listFavoritesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.ReactionServiceListFavoritesArgs)
	realResult := result.(*rpc.ReactionServiceListFavoritesResult)
	success, err := handler.(rpc.ReactionService).ListFavorites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReactionServiceListFavoritesArgs() interface{} {
	return rpc.NewReactionServiceListFavoritesArgs()
}

func newReactionServiceListFavoritesResult() interface{} {
	return rpc.NewReactionServiceListFavoritesResult()
}

func commentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.ReactionServiceCommentArgs)
	realResult := result.(*rpc.ReactionServiceCommentResult)
	success, err := handler.(rpc.ReactionService).Comment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReactionServiceCommentArgs() interface{} {
	return rpc.NewReactionServiceCommentArgs()
}

func newReactionServiceCommentResult() interface{} {
	return rpc.NewReactionServiceCommentResult()
}

func listCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.ReactionServiceListCommentsArgs)
	realResult := result.(*rpc.ReactionServiceListCommentsResult)
	success, err := handler.(rpc.ReactionService).ListComments(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReactionServiceListCommentsArgs() interface{} {
	return rpc.NewReactionServiceListCommentsArgs()
}

func newReactionServiceListCommentsResult() interface{} {
	return rpc.NewReactionServiceListCommentsResult()
}

func testFavoritesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.ReactionServiceTestFavoritesArgs)
	realResult := result.(*rpc.ReactionServiceTestFavoritesResult)
	success, err := handler.(rpc.ReactionService).TestFavorites(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReactionServiceTestFavoritesArgs() interface{} {
	return rpc.NewReactionServiceTestFavoritesArgs()
}

func newReactionServiceTestFavoritesResult() interface{} {
	return rpc.NewReactionServiceTestFavoritesResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Favorite(ctx context.Context, req *rpc.DouyinFavoriteActionRequest) (r *rpc.DouyinFavoriteActionResponse, err error) {
	var _args rpc.ReactionServiceFavoriteArgs
	_args.Req = req
	var _result rpc.ReactionServiceFavoriteResult
	if err = p.c.Call(ctx, "Favorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFavorites(ctx context.Context, req *rpc.DouyinFavoriteListRequest) (r *rpc.DouyinFavoriteListResponse, err error) {
	var _args rpc.ReactionServiceListFavoritesArgs
	_args.Req = req
	var _result rpc.ReactionServiceListFavoritesResult
	if err = p.c.Call(ctx, "ListFavorites", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Comment(ctx context.Context, req *rpc.DouyinCommentActionRequest) (r *rpc.DouyinCommentActionResponse, err error) {
	var _args rpc.ReactionServiceCommentArgs
	_args.Req = req
	var _result rpc.ReactionServiceCommentResult
	if err = p.c.Call(ctx, "Comment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListComments(ctx context.Context, req *rpc.DouyinCommentListRequest) (r *rpc.DouyinCommentListResponse, err error) {
	var _args rpc.ReactionServiceListCommentsArgs
	_args.Req = req
	var _result rpc.ReactionServiceListCommentsResult
	if err = p.c.Call(ctx, "ListComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TestFavorites(ctx context.Context, req *rpc.FavoriteTestRequest) (r *rpc.FavoriteTestResponse, err error) {
	var _args rpc.ReactionServiceTestFavoritesArgs
	_args.Req = req
	var _result rpc.ReactionServiceTestFavoritesResult
	if err = p.c.Call(ctx, "TestFavorites", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
