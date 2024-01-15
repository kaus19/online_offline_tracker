# online_offline_tracker

The purpose of this application is: An API in golang returning status of an existing user. 
- Utilizes redis database.
- Utilizes chi go package(small, idiomatic and composable router for building HTTP services).
- With Authorization check(add 'Authorization' token in request header).
- Hit GET request at endpoint  http://localhost:8000/account/status?username={username}

- Response: 
    - if successful: return online/offline status;
    - otherwise: throws an error message

[To-do] 
1. Adding an user via API call
2. Showing list of all users
3. User can change their status to online
4. User's status will be offline after 1 min
5. If the user is offline, also show "was online X mins ago"