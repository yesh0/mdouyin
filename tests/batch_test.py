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


def random_name():
    """Generate a random name."""
    return "".join(map(
        lambda x: random.choice(string.ascii_lowercase + string.digits),
        [0] * 8
    ))


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


if __name__ == "__main__":
    s = Server("http://127.0.0.1:8000")
    for i in range(10):
        test_user_info(s)
