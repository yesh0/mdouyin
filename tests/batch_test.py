#!/usr/bin/python3

"""Tests for mdouyin."""

import colorama
import random
import requests
import string
import sys
import time
import ttypes


def CamelCase(s):
    """Turn a string into CamelCase."""
    return "".join(map(str.capitalize, s.split("_")))


def cast_ttype(json, ttype):
    """Convert json to the type."""
    return ttype(
        **{CamelCase(k): v for k, v in json.items()}
    )


def cast_ttype_array(array, ttype):
    """Convert an array of the type."""
    return list(map(lambda j: cast_ttype(j, ttype), array))


def assert_ok(response: requests.Response, ttype):
    """Assert and return the inner json."""
    assert response.status_code == 200
    json = response.json()
    assert json["status_code"] == 0
    return cast_ttype(json, ttype)


class Server:
    """The mdouyin server instance."""

    def __init__(self, base: str):
        """Initialize the server address."""
        self.base = base

    def register_user(self, name: str, password: str) -> ttypes.DouyinUserRegisterResponse:
        """Register a new user, returning the info."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/user/register/?username={name}&password={password}"
        ), ttypes.DouyinUserRegisterResponse)

    def user_info(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinUserResponse:
        """Retrieve user info."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/user/?user_id={user.UserId}&token={user.Token}"
        ), ttypes.DouyinUserResponse)

    def login(self, name: str, password: str) -> ttypes.DouyinUserLoginResponse:
        """Login."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/user/login/?username={name}&password={password}"
        ), ttypes.DouyinUserLoginResponse)

    def follow(self, me: ttypes.DouyinUserLoginResponse, user: ttypes.DouyinUserLoginResponse):
        """Follow a user."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/relation/action/?token={me.Token}&to_user_id={user.UserId}&action_type=1"
        ), ttypes.DouyinRelationActionResponse)

    def unfollow(self, me: ttypes.DouyinUserLoginResponse, user: ttypes.DouyinUserLoginResponse):
        """Unfollow a user."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/relation/action/?token={me.Token}&to_user_id={user.UserId}&action_type=2"
        ), ttypes.DouyinRelationActionResponse)

    def list_follow(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinRelationFollowListResponse:
        """List following users."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/relation/follow/list/?token={user.Token}&user_id={user.UserId}"
        ), ttypes.DouyinRelationFollowListResponse)

    def list_follower(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinRelationFollowerListResponse:
        """List following users."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/relation/follower/list/?token={user.Token}&user_id={user.UserId}"
        ), ttypes.DouyinRelationFollowerListResponse)

    def publish(self, user: ttypes.DouyinUserLoginResponse, file: str, title: str):
        """Publish a video."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/publish/action/",
            data={
                "token": user.Token,
                "title": title,
            },
            files={
                "data": open(file, "rb"),
            },
        ), ttypes.DouyinPublishActionResponse)

    def list_videos(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinPublishListResponse:
        """List videos published by the user."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/publish/list/?token={user.Token}&user_id={user.UserId}"
        ), ttypes.DouyinPublishListResponse)

    def feed(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinFeedResponse:
        """Fetch the feed."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/feed/?token={user.Token}"
        ), ttypes.DouyinFeedResponse)

    def like(self, user: ttypes.DouyinUserLoginResponse, video: ttypes.Video):
        """Favorite a video."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/favorite/action/?token={user.Token}&video_id={video.Id}&action_type=1"
        ), ttypes.DouyinFavoriteActionResponse)

    def list_likes(self, user: ttypes.DouyinUserLoginResponse) -> ttypes.DouyinFavoriteListResponse:
        """List favorite videos."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/favorite/list/?token={user.Token}&user_id={user.UserId}"
        ), ttypes.DouyinFavoriteListResponse)

    def unlike(self, user: ttypes.DouyinUserLoginResponse, video: ttypes.Video):
        """Undo favoriting a video."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/favorite/action/?token={user.Token}&video_id={video.Id}&action_type=2"
        ), ttypes.DouyinFavoriteActionResponse)


def random_name():
    """Generate a random name."""
    return "".join(map(
        lambda x: random.choice(string.ascii_lowercase + string.digits),
        [0] * 8
    ))


indent = 0
def log_test(func):
    count = 0
    colorama.just_fix_windows_console()
    def wrapper(*args, **kwargs):
        global indent
        nonlocal count
        name = func.__name__
        doc = func.__doc__
        prefix = ">>>>" * indent
        info = "" if indent == 0 else "nested "
        print(f"{prefix} Running {info}{colorama.Fore.YELLOW}{name}({count}):{colorama.Style.RESET_ALL}")
        print(
            f"{prefix} {colorama.Fore.WHITE}{colorama.Style.BRIGHT}  {doc}{colorama.Style.RESET_ALL}",
            end=None,
        )
        indent += 1
        result = func(*args, **kwargs)
        indent -= 1
        print(f"\t{colorama.Fore.GREEN}{colorama.Style.BRIGHT}ok{colorama.Style.RESET_ALL}")
        count = count + 1
        return result
    return wrapper


@log_test
def test_user_info(s: Server):
    """Test registering and retrieving info."""
    name = random_name()
    password = "password"
    user = s.register_user(name, password)
    u = cast_ttype(s.user_info(user).User, ttypes.User)
    assert u.Id == user.UserId
    assert u.Name == name
    assert not u.IsFollow
    assert u.Avatar
    v = s.login(name, password)
    assert v.UserId
    assert v.UserId == user.UserId


@log_test
def test_relation(s: Server, test=None):
    """Test following, unfollowing and counters."""
    users = []
    password = "password"
    for i in range(10):
        name = random_name()
        users.append(s.register_user(name, password))
    for i, me in enumerate(users):
        for j, user in enumerate(users[i+1:]):
            s.follow(user, me)
            u = cast_ttype(
                s.user_info(me).User,
                ttypes.User
            )
            assert u.FollowerCount == j + 1
            assert u.FollowCount == i
    if test != None:
        test(users)
    for i, me in enumerate(users):
        for j, user in enumerate(users[i+1:]):
            s.unfollow(user, me)
            u = cast_ttype(
                s.user_info(me).User,
                ttypes.User
            )
            assert u.FollowerCount == 10 - i - j - 2
            assert u.FollowCount == 0


@log_test
def test_follow_list(s: Server):
    """Test follow, follower listing."""
    def assert_list_match(users, list):
        assert len(users) == len(list)
        ids = []
        for user in users:
            ids.append(user.Id)
        for user in list:
            assert user.UserId in ids
    def tester(users):
        for i, user in enumerate(users):
            list = s.list_follower(user)
            assert_list_match(cast_ttype_array(list.UserList, ttypes.User), users[i+1:])
            list = s.list_follow(user)
            assert_list_match(cast_ttype_array(list.UserList, ttypes.User), users[:i])
    test_relation(s, test=tester)


def assert_contains(videos, title) -> ttypes.Video:
    for video in videos:
        if video.Title == title:
            return video
    assert video.Title == title


@log_test
def test_video_publish(s: Server, test=None):
    """Test video publishing, listing and feed."""
    def tester(users):
        titles = []
        for user in users:
            title = "CC Ink Stamp Animation " + random_name()
            s.publish(user, "./cc_ink_stamp_animation_cc0.mp4", title)
            titles.append(title)
        # The server needs some time to generate the cover images.
        time.sleep(2)
        published = []
        for user, title in zip(users, titles):
            res = s.list_videos(user)
            videos = cast_ttype_array(res.VideoList, ttypes.Video)
            assert len(videos) == 1
            v: ttypes.Video = videos[0]
            published.append(v)
            assert v.Author["id"] == user.UserId
            assert v.PlayUrl
            assert v.CoverUrl
            assert v.Title == title
        for i, user in enumerate(users):
            feed = s.feed(user)
            assert feed.NextTime
            videos = cast_ttype_array(feed.VideoList, ttypes.Video)
            following = titles[:i]
            assert len(videos) >= len(following)
            for title in following:
                assert_contains(videos, title)
        if test != None:
            test(users, published)
    test_relation(s, test=tester)


@log_test
def test_video_reaction(s: Server):
    """Test video reaction listing and counters."""
    def tester(users, published):
        for i, user in enumerate(users):
            for following in published[:i]:
                s.like(user, following)
        for i, user in enumerate(users):
            own = s.list_videos(user)
            assert len(own.VideoList) == 1
            v = cast_ttype(own.VideoList[0], ttypes.Video)
            assert v.CommentCount == 0
            assert v.FavoriteCount == 10 - 1 - i
            feed = s.feed(users[-1])
            likes = s.list_likes(user)
            for videos in [
                cast_ttype_array(feed.VideoList, ttypes.Video),
                cast_ttype_array(likes.VideoList, ttypes.Video),
            ]:
                for j, following in enumerate(published[:i]):
                    v = assert_contains(videos, following.Title)
                    assert v.CommentCount == 0
                    assert v.FavoriteCount == 10 - 1 - j
                    assert v.IsFavorite

        for i, user in enumerate(users):
            for following in published[:i]:
                s.unlike(user, following)
        for i, user in enumerate(users):
            own = s.list_videos(user)
            assert len(own.VideoList) == 1
            v = cast_ttype(own.VideoList[0], ttypes.Video)
            assert v.CommentCount == 0
            assert v.FavoriteCount == 0
            likes = s.list_likes(user)
            assert len(likes.VideoList) == 0
    test_video_publish(s, test=tester)


if __name__ == "__main__":
    args = sys.argv[1:]
    available = []
    def wants(s: str):
        available.append(s)
        return len(args) == 0 or s in args
    s = Server("http://127.0.0.1:8000")
    if wants("user"):
        for i in range(3):
            test_user_info(s)
    if wants("follow"):
        for i in range(3):
            # test_relation(s): Tested by test_follow_list
            test_follow_list(s)
    if wants("publish"):
        test_video_publish(s)
    if wants("reaction"):
        test_video_reaction(s)

    if len(args) != 0 and args[0] in ["-h", "--help", "help"]:
        print("Available tests:", available)
