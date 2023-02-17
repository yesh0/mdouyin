#!/usr/bin/python3

"""Tests for mdouyin."""

import random
import requests
import string
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


def random_name():
    """Generate a random name."""
    return "".join(map(
        lambda x: random.choice(string.ascii_lowercase + string.digits),
        [0] * 8
    ))


def log_test(func):
    count = 0
    def wrapper(*args, **kwargs):
        nonlocal count
        name = func.__name__
        print(f">>> {name}({count}): {func.__doc__}")
        result = func(*args, **kwargs)
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


if __name__ == "__main__":
    s = Server("http://127.0.0.1:8000")
    for i in range(10):
        test_user_info(s)
    for i in range(3):
        test_relation(s)
    for i in range(3):
        test_follow_list(s)
