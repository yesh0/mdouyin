namespace go douyin.rpc

include "common.thrift"

service FeedService {
    DouyinFeedResponse Feed (1: DouyinFeedRequest Req) (api.get="/douyin/feed/")
    DouyinPublishActionResponse Publish (1: DouyinPublishActionRequest Req) (api.post="/douyin/publish/action/")
    DouyinPublishListResponse List (1: DouyinPublishListRequest Req) (api.get="/douyin/publish/list/")
    DouyinRelationActionResponse Relation (1: DouyinRelationActionRequest Req) (api.post="/douyin/relation/action/")
    DouyinRelationFollowListResponse Following (1: DouyinRelationFollowListRequest Req) (api.get="/douyin/relation/follow/list/")
    DouyinRelationFollowerListResponse Follower (1: DouyinRelationFollowerListRequest Req) (api.get="/douyin/relation/follower/list/")
    DouyinRelationFriendListResponse Friend (1: DouyinRelationFriendListRequest Req) (api.get="/douyin/relation/friend/list/")

    VideoBatchInfoResponse VideoInfo (1: VideoBatchInfoRequest Req)
    FriendCheckResponse IsFriend(1: FriendCheckRequest Req)
}

struct FriendCheckRequest {
    1: i64 UserId
    2: i64 RequestUserId
}
struct FriendCheckResponse {
    1: i8 IsFriend
}
struct VideoBatchInfoRequest {
    1: list<i64> VideoIds // 视频id
    2: i64 RequestUserId  // 用户id
}
struct VideoBatchInfoResponse {
    1: list<common.Video> Videos
}
struct DouyinFeedRequest {
    1: optional i64 LatestTime (api.body="latest_time", api.query="latest_time", api.form="latest_time") // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    // 2: optional string Token (api.body="token", api.query="token", api.form="token") // 可选参数，登录用户设置
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinFeedResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<common.Video> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 视频列表
    4: optional i64 NextTime (api.body="next_time", api.query="next_time", api.form="next_time") // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
struct DouyinPublishActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    // 2: binary Data (api.body="data", api.query="data", api.form="data") // 视频数据
    // 3: string Title (api.body="title", api.query="title", api.form="title") // 视频标题
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: common.Video Video (api.body="token", api.query="token", api.form="token") // 视频信息
}
struct DouyinPublishActionResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
}
struct DouyinPublishListRequest {
    1: i64 UserId (api.body="user_id", api.query="user_id", api.form="user_id") // 用户id
    // 2: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinPublishListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<common.Video> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 用户发布的视频列表
}
struct DouyinRelationActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 ToUserId (api.body="to_user_id", api.query="to_user_id", api.form="to_user_id") // 对方用户id
    3: i32 ActionType (api.body="action_type", api.query="action_type", api.form="action_type") // 1-关注，2-取消关注
}
struct DouyinRelationActionResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
}
struct DouyinRelationFollowListRequest {
    1: i64 UserId (api.body="user_id", api.query="user_id", api.form="user_id") // 用户id
    // 2: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinRelationFollowListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<i64> UserList (api.body="user_list", api.query="user_list", api.form="user_list") // 用户信息列表
}
struct DouyinRelationFollowerListRequest {
    1: i64 UserId (api.body="user_id", api.query="user_id", api.form="user_id") // 用户id
    // 2: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinRelationFollowerListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<i64> UserList (api.body="user_list", api.query="user_list", api.form="user_list") // 用户列表
}
struct DouyinRelationFriendListRequest {
    1: i64 UserId (api.body="user_id", api.query="user_id", api.form="user_id") // 用户id
    // 2: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinRelationFriendListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<i64> UserList (api.body="user_list", api.query="user_list", api.form="user_list") // 用户列表
}
