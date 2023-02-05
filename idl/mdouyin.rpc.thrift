namespace go douyin.rpc


service ReactionService {
    DouyinFavoriteActionResponse Favorite (1: DouyinFavoriteActionRequest Req) (api.post="/douyin/favorite/action/")
    DouyinFavoriteListResponse ListFavorites (1: DouyinFavoriteListRequest Req) (api.get="/douyin/favorite/list/")
    DouyinCommentActionResponse Comment (1: DouyinCommentActionRequest Req) (api.post="/douyin/comment/action/")
    DouyinCommentListResponse ListComments (1: DouyinCommentListRequest Req) (api.get="/douyin/comment/list/")
}

service MessageService {
    DouyinMessageChatResponse Chat (1: DouyinMessageChatRequest Req) (api.get="/douyin/message/chat/")
    DouyinMessageActionResponse Message (1: DouyinMessageActionRequest Req) (api.post="/douyin/message/action/")
}

service FeedService {
    DouyinFeedResponse Feed (1: DouyinFeedRequest Req) (api.get="/douyin/feed/")
    DouyinPublishActionResponse Publish (1: DouyinPublishActionRequest Req) (api.post="/douyin/publish/action/")
    DouyinPublishListResponse List (1: DouyinPublishListRequest Req) (api.get="/douyin/publish/list/")
    DouyinRelationActionResponse Relation (1: DouyinRelationActionRequest Req) (api.post="/douyin/relation/action/")
    DouyinRelationFollowListResponse Following (1: DouyinRelationFollowListRequest Req) (api.get="/douyin/relation/follow/list/")
    DouyinRelationFollowerListResponse Follower (1: DouyinRelationFollowerListRequest Req) (api.get="/douyin/relation/follower/list/")
    DouyinRelationFriendListResponse Friend (1: DouyinRelationFriendListRequest Req) (api.get="/douyin/relation/friend/list/")
}


struct Video {
    1: i64 Id (api.body="id", api.query="id", api.form="id") // 视频唯一标识
    2: User Author (api.body="author", api.query="author", api.form="author") // 视频作者信息
    3: string PlayUrl (api.body="play_url", api.query="play_url", api.form="play_url") // 视频播放地址
    4: string CoverUrl (api.body="cover_url", api.query="cover_url", api.form="cover_url") // 视频封面地址
    5: i64 FavoriteCount (api.body="favorite_count", api.query="favorite_count", api.form="favorite_count") // 视频的点赞总数
    6: i64 CommentCount (api.body="comment_count", api.query="comment_count", api.form="comment_count") // 视频的评论总数
    7: bool IsFavorite (api.body="is_favorite", api.query="is_favorite", api.form="is_favorite") // true-已点赞，false-未点赞
    8: string Title (api.body="title", api.query="title", api.form="title") // 视频标题
}
struct User {
    1: i64 Id (api.body="id", api.query="id", api.form="id") // 用户id
    2: string Name (api.body="name", api.query="name", api.form="name") // 用户名称
    3: optional i64 FollowCount (api.body="follow_count", api.query="follow_count", api.form="follow_count") // 关注总数
    4: optional i64 FollowerCount (api.body="follower_count", api.query="follower_count", api.form="follower_count") // 粉丝总数
    5: bool IsFollow (api.body="is_follow", api.query="is_follow", api.form="is_follow") // true-已关注，false-未关注
    6: string Avatar (api.body="avatar", api.query="avatar", api.form="avatar") // 用户头像Url
    7: optional string Message (api.body="message", api.query="message", api.form="message") // 和该好友的最新聊天消息
    8: optional i64 MsgType (api.body="msg_type", api.query="msg_type", api.form="msg_type") // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}
struct Comment {
    1: i64 Id (api.body="id", api.query="id", api.form="id") // 视频评论id
    2: User User (api.body="user", api.query="user", api.form="user") // 评论用户信息
    3: string Content (api.body="content", api.query="content", api.form="content") // 评论内容
    4: string CreateDate (api.body="create_date", api.query="create_date", api.form="create_date") // 评论发布日期，格式 mm-dd
}
struct DouyinFeedRequest {
    1: optional i64 LatestTime (api.body="latest_time", api.query="latest_time", api.form="latest_time") // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    // 2: optional string Token (api.body="token", api.query="token", api.form="token") // 可选参数，登录用户设置
    2: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
}
struct DouyinFeedResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<Video> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 视频列表
    4: optional i64 NextTime (api.body="next_time", api.query="next_time", api.form="next_time") // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
struct DouyinPublishActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    // 2: binary Data (api.body="data", api.query="data", api.form="data") // 视频数据
    // 3: string Title (api.body="title", api.query="title", api.form="title") // 视频标题
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: Video Video (api.body="token", api.query="token", api.form="token") // 视频信息
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
    3: list<Video> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 用户发布的视频列表
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
    3: list<Video> VideoList (api.body="video_list", api.query="video_list", api.form="video_list") // 用户点赞视频列表
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
    3: optional Comment Comment (api.body="comment", api.query="comment", api.form="comment") // 评论成功返回评论内容，不需要重新拉取整个列表
}
struct DouyinCommentListRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 VideoId (api.body="video_id", api.query="video_id", api.form="video_id") // 视频id
}
struct DouyinCommentListResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<Comment> CommentList (api.body="comment_list", api.query="comment_list", api.form="comment_list") // 评论列表
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
struct DouyinMessageChatRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 ToUserId (api.body="to_user_id", api.query="to_user_id", api.form="to_user_id") // 对方用户id
}
struct DouyinMessageChatResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
    3: list<Message> MessageList (api.body="message_list", api.query="message_list", api.form="message_list") // 消息列表
}
struct Message {
    1: i64 Id (api.body="id", api.query="id", api.form="id") // 消息id
    2: i64 ToUserId (api.body="to_user_id", api.query="to_user_id", api.form="to_user_id") // 该消息接收者的id
    3: i64 FromUserId (api.body="from_user_id", api.query="from_user_id", api.form="from_user_id") // 该消息发送者的id
    4: string Content (api.body="content", api.query="content", api.form="content") // 消息内容
    5: optional string CreateTime (api.body="create_time", api.query="create_time", api.form="create_time") // 消息创建时间
}
struct DouyinMessageActionRequest {
    // 1: string Token (api.body="token", api.query="token", api.form="token") // 用户鉴权token
    1: i64 RequestUserId (api.body="token", api.query="token", api.form="token") // 用户id
    2: i64 ToUserId (api.body="to_user_id", api.query="to_user_id", api.form="to_user_id") // 对方用户id
    3: i32 ActionType (api.body="action_type", api.query="action_type", api.form="action_type") // 1-发送消息
    4: string Content (api.body="content", api.query="content", api.form="content") // 消息内容
}
struct DouyinMessageActionResponse {
    1: i32 StatusCode (api.body="status_code", api.query="status_code", api.form="status_code") // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg (api.body="status_msg", api.query="status_msg", api.form="status_msg") // 返回状态描述
}
