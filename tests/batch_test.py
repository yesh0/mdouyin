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


def assert_ok(response, ttype):
    """Assert and return the inner json."""
    assert response.status_code == 200
    json = response.json()
    assert json["status_code"] == 0
    return cast_ttype(json, ttype)


class Server:
    """The mdouyin server instance."""

    def __init__(self, base):
        """Initialize the server address."""
        self.base = base

    def register_user(self, name, password):
        """Register a new user, returning the info."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/user/register/?username={name}"
            + f"&password={password}"
        ), ttypes.DouyinUserRegisterResponse)

    def user_info(self, user):
        """Retrieve user info."""
        return assert_ok(requests.get(
            f"{self.base}/douyin/user/?user_id={user.UserId}"
            + f"&token={user.Token}"
        ), ttypes.DouyinUserResponse)

    def login(self, name, password):
        """Login."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/user/login/?username={name}"
            + f"&password={password}"
        ), ttypes.DouyinUserLoginResponse)

    def follow(self, me, user):
        """Follow a user."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/relation/action/?token={me.Token}"
            + f"&to_user_id={user.UserId}&action_type=1"
        ), ttypes.DouyinRelationActionResponse)

    def unfollow(self, me, user):
        """Unfollow a user."""
        return assert_ok(requests.post(
            f"{self.base}/douyin/relation/action/?token={me.Token}"
            + f"&to_user_id={user.UserId}&action_type=2"
        ), ttypes.DouyinRelationActionResponse)


def random_name():
    """Generate a random name."""
    return "".join(map(
        lambda x: random.choice(string.ascii_lowercase + string.digits),
        [0] * 8
    ))


test_counts = {}
def log_test(func):
    def wrapper(*args, **kwargs):
        name = func.__name__
        if name not in test_counts:
            test_counts[name] = 0
        print(f">>> {name}({test_counts[name]}): {func.__doc__}")
        result = func(*args, **kwargs)
        test_counts[name] = test_counts[name] + 1
        return result
    return wrapper


@log_test
def test_user_info(s):
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
def test_relation(s):
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
    for i, me in enumerate(users):
        for j, user in enumerate(users[i+1:]):
            s.unfollow(user, me)
            u = cast_ttype(
                s.user_info(me).User,
                ttypes.User
            )
            assert u.FollowerCount == 10 - i - j - 2
            assert u.FollowCount == 0


if __name__ == "__main__":
    s = Server("http://127.0.0.1:8000")
    for i in range(10):
        test_user_info(s)
    test_relation(s)
