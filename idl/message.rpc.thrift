namespace go douyin.rpc

include "common.thrift"

service MessageService {
    DouyinMessageChatResponse Chat (1: DouyinMessageChatRequest Req) (api.get="/douyin/message/chat/")
    DouyinMessageActionResponse Message (1: DouyinMessageActionRequest Req) (api.post="/douyin/message/action/")
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
