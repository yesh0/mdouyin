# mdouyin - A Micro Douyin Implmentation

## TODO

- [X] User API:
  - [X] Registration
  - [X] Logging-in
  - [X] Info:
    - [X] Basic Info
    - [X] Assemble relationship data & counters (e.g. follower count)
- [X] Feed API:
  - [X] Assemble counter data (e.g. like/comment counts)
  - [X] Publish: Uploads a video, generates its cover image,
    and pushes the info to followers' inboxes.
  - [X] Feed viewing: Lists one's inbox.
  - [X] Listing: Lists one's submissions.
- [X] A counter service
- [X] Reaction API
- [X] Chat API

## Bugs

- [X] The app does not seem to receive the correct feed:
  a user obviously has an empty feed, but the app stil shows something.
  Probably the app has some kind of internal cache such that it merges all feeds?
- [ ] File uploading is really slow.

## Improvements

- [ ] Uploaded files is either kept fully in memory or stored to the disk temporarily.
  However, if it is stored as a temporary file, we should want move it in place instead
  of copying it.
- [ ] Cache.