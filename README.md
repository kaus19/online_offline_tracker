# online_offline_tracker

The purpose of this application is returning the online status of user(s).

## Tech Stack
- Utilizes [redis database](https://redis.io/docs/about/)
- Utilizes [chi go package](https://pkg.go.dev/github.com/go-chi/chi)

## Features
1. Shows all online users
2. Users can be added
3. User can change their status to online
4. User's status will be offline(data removed from redis) after 30sec
5. Connection pooling and timeout

## Endpoints
- GET request at endpoint  http://localhost:8000/account/status?username={username}
    - Response: 
        - if user is present: return their online status
        - if user is not present: add user and set status to online

- GET request at endpoint http://localhost:8000/account/all
    - Response:
        - Returns list of all users and their status

## Future Implementation
- If the user is offline, show "was online X mins ago"
- Dashboards to show redis metrics