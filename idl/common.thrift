namespace go douyin.rpc


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