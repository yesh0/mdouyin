namespace go douyin.rpc

include "common.thrift"

service ReactionService {
    DouyinFavoriteActionResponse Favorite (1: DouyinFavoriteActionRequest Req) (api.post="/douyin/favorite/action/")
    DouyinFavoriteListResponse ListFavorites (1: DouyinFavoriteListRequest Req) (api.get="/douyin/favorite/list/")
    DouyinCommentActionResponse Comment (1: DouyinCommentActionRequest Req) (api.post="/douyin/comment/action/")
    DouyinCommentListResponse ListComments (1: DouyinCommentListRequest Req) (api.get="/douyin/comment/list/")

    FavoriteTestResponse TestFavorites (1: FavoriteTestRequest Req);
}


struct FavoriteTestRequest {
    1: i64 RequestUserId  // 用户id
    2: list<i64> Videos // 视频id
}
struct FavoriteTestResponse {
    1: i32 StatusCode // 状态码，0-成功，其他值-失败
    2: list<i8> IsFavorites // Favorite 状况
}
struct DouyinFavoriteActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 VideoId (api.body="video_id", api.query="video_id", api.form="video_id") // 视频id
    3: i32 ActionType (api.body="action_type", api.query="action_type", api.form="action_type") // 1-点赞，2-取消点赞
}
struct DouyinFavoriteActionResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
}
struct DouyinFavoriteListRequest {
    1: i64 UserId (api.body="user_id", api.query="user_id", api.form="user_id") // 用户id
    // 2: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinFavoriteListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<i64> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 用户点赞视频列表
}
struct DouyinCommentActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 VideoId (api.body="video_id", api.query="video_id", api.form="video_id") // 视频id
    3: i32 ActionType (api.body="action_type", api.query="action_type", api.form="action_type") // 1-发布评论，2-删除评论
    4: optional string CommentText (api.body="comment_text", api.query="comment_text", api.form="comment_text") // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 CommentId (api.body="comment_id", api.query="comment_id", api.form="comment_id") // 要删除的评论id，在action_type=2的时候使用
}
struct DouyinCommentActionResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: optional common.Comment Comment (api.body="comment", api.query="comment", api.form="comment") // 评论成功返回评论内容，不需要重新拉取整个列表
}
struct DouyinCommentListRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 VideoId (api.body="video_id", api.query="video_id", api.form="video_id") // 视频id
}
struct DouyinCommentListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<common.Comment> CommentList (api.body="comment_list", api.query="comment_list", api.form="comment_list") // 评论列表
}
