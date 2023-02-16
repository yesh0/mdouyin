#
# Autogenerated by Thrift Compiler (0.17.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

class Video(object):
    """
    Attributes:
     - Id
     - Author
     - PlayUrl
     - CoverUrl
     - FavoriteCount
     - CommentCount
     - IsFavorite
     - Title

    """


    def __init__(self, Id=None, Author=None, PlayUrl=None, CoverUrl=None, FavoriteCount=None, CommentCount=None, IsFavorite=None, Title=None,):
        self.Id = Id
        self.Author = Author
        self.PlayUrl = PlayUrl
        self.CoverUrl = CoverUrl
        self.FavoriteCount = FavoriteCount
        self.CommentCount = CommentCount
        self.IsFavorite = IsFavorite
        self.Title = Title


class User(object):
    """
    Attributes:
     - Id
     - Name
     - FollowCount
     - FollowerCount
     - IsFollow
     - Avatar
     - Message
     - MsgType

    """


    def __init__(self, Id=None, Name=None, FollowCount=None, FollowerCount=None, IsFollow=None, Avatar=None, Message=None, MsgType=None,):
        self.Id = Id
        self.Name = Name
        self.FollowCount = FollowCount
        self.FollowerCount = FollowerCount
        self.IsFollow = IsFollow
        self.Avatar = Avatar
        self.Message = Message
        self.MsgType = MsgType


class Comment(object):
    """
    Attributes:
     - Id
     - User
     - Content
     - CreateDate

    """


    def __init__(self, Id=None, User=None, Content=None, CreateDate=None,):
        self.Id = Id
        self.User = User
        self.Content = Content
        self.CreateDate = CreateDate


class DouyinFeedRequest(object):
    """
    Attributes:
     - LatestTime
     - Token

    """


    def __init__(self, LatestTime=None, Token=None,):
        self.LatestTime = LatestTime
        self.Token = Token

class DouyinFeedResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - VideoList
     - NextTime

    """


    def __init__(self, StatusCode=None, StatusMsg=None, VideoList=None, NextTime=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.VideoList = VideoList
        self.NextTime = NextTime

class DouyinUserRegisterRequest(object):
    """
    Attributes:
     - Username
     - Password

    """


    def __init__(self, Username=None, Password=None,):
        self.Username = Username
        self.Password = Password

class DouyinUserRegisterResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - UserId
     - Token

    """


    def __init__(self, StatusCode=None, StatusMsg=None, UserId=None, Token=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.UserId = UserId
        self.Token = Token

class DouyinUserLoginRequest(object):
    """
    Attributes:
     - Username
     - Password

    """


    def __init__(self, Username=None, Password=None,):
        self.Username = Username
        self.Password = Password

class DouyinUserLoginResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - UserId
     - Token

    """


    def __init__(self, StatusCode=None, StatusMsg=None, UserId=None, Token=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.UserId = UserId
        self.Token = Token

class DouyinUserRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinUserResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - User

    """


    def __init__(self, StatusCode=None, StatusMsg=None, User=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.User = User

class DouyinPublishActionRequest(object):
    """
    Attributes:
     - Token
     - Data
     - Title

    """


    def __init__(self, Token=None, Data=None, Title=None,):
        self.Token = Token
        self.Data = Data
        self.Title = Title

class DouyinPublishActionResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg

    """


    def __init__(self, StatusCode=None, StatusMsg=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg

class DouyinPublishListRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinPublishListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - VideoList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, VideoList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.VideoList = VideoList

class DouyinFavoriteActionRequest(object):
    """
    Attributes:
     - Token
     - VideoId
     - ActionType

    """


    def __init__(self, Token=None, VideoId=None, ActionType=None,):
        self.Token = Token
        self.VideoId = VideoId
        self.ActionType = ActionType

class DouyinFavoriteActionResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg

    """


    def __init__(self, StatusCode=None, StatusMsg=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg

class DouyinFavoriteListRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinFavoriteListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - VideoList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, VideoList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.VideoList = VideoList

class DouyinCommentActionRequest(object):
    """
    Attributes:
     - Token
     - VideoId
     - ActionType
     - CommentText
     - CommentId

    """


    def __init__(self, Token=None, VideoId=None, ActionType=None, CommentText=None, CommentId=None,):
        self.Token = Token
        self.VideoId = VideoId
        self.ActionType = ActionType
        self.CommentText = CommentText
        self.CommentId = CommentId

class DouyinCommentActionResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - Comment

    """


    def __init__(self, StatusCode=None, StatusMsg=None, Comment=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.Comment = Comment

class DouyinCommentListRequest(object):
    """
    Attributes:
     - Token
     - VideoId

    """


    def __init__(self, Token=None, VideoId=None,):
        self.Token = Token
        self.VideoId = VideoId

class DouyinCommentListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - CommentList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, CommentList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.CommentList = CommentList

class DouyinRelationActionRequest(object):
    """
    Attributes:
     - Token
     - ToUserId
     - ActionType

    """


    def __init__(self, Token=None, ToUserId=None, ActionType=None,):
        self.Token = Token
        self.ToUserId = ToUserId
        self.ActionType = ActionType

class DouyinRelationActionResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg

    """


    def __init__(self, StatusCode=None, StatusMsg=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg

class DouyinRelationFollowListRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinRelationFollowListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - UserList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, UserList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.UserList = UserList

class DouyinRelationFollowerListRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinRelationFollowerListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - UserList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, UserList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.UserList = UserList

class DouyinRelationFriendListRequest(object):
    """
    Attributes:
     - UserId
     - Token

    """


    def __init__(self, UserId=None, Token=None,):
        self.UserId = UserId
        self.Token = Token

class DouyinRelationFriendListResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - UserList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, UserList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.UserList = UserList

class DouyinMessageChatRequest(object):
    """
    Attributes:
     - Token
     - ToUserId

    """


    def __init__(self, Token=None, ToUserId=None,):
        self.Token = Token
        self.ToUserId = ToUserId

class DouyinMessageChatResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg
     - MessageList

    """


    def __init__(self, StatusCode=None, StatusMsg=None, MessageList=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
        self.MessageList = MessageList

class Message(object):
    """
    Attributes:
     - Id
     - ToUserId
     - FromUserId
     - Content
     - CreateTime

    """


    def __init__(self, Id=None, ToUserId=None, FromUserId=None, Content=None, CreateTime=None,):
        self.Id = Id
        self.ToUserId = ToUserId
        self.FromUserId = FromUserId
        self.Content = Content
        self.CreateTime = CreateTime

class DouyinMessageActionRequest(object):
    """
    Attributes:
     - Token
     - ToUserId
     - ActionType
     - Content

    """


    def __init__(self, Token=None, ToUserId=None, ActionType=None, Content=None,):
        self.Token = Token
        self.ToUserId = ToUserId
        self.ActionType = ActionType
        self.Content = Content

class DouyinMessageActionResponse(object):
    """
    Attributes:
     - StatusCode
     - StatusMsg

    """


    def __init__(self, StatusCode=None, StatusMsg=None,):
        self.StatusCode = StatusCode
        self.StatusMsg = StatusMsg
