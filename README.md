# online_offline_tracker

The purpose of this application is: An API in golang returning status of an existing user. 
- Uses mock database.
- With Authorization check(add 'Authorization' token in request header).
- Hit GET request at endpoint  http://localhost:8000/account/status?username={username}