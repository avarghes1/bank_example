# Bank Example

The following is a sample bank api.

[Sample bank api](http://ec2-54-172-55-155.compute-1.amazonaws.com)

The application uses basic auth for all api except user creation.

# API

1. PUT /api/v1/users
    - create a user (requires user json in the body of the request)

1. POST /api/v1/users
    - update user (basic auth and requires user json in body of the request)

1. GET /api/v1/users
    - fetch a user (uses basic auth to check auth)
```
curl 'http://ec2-54-172-55-155.compute-1.amazonaws.com/api/v1/users' --compressed -H 'Authorization: Basic dGVzdEB0ZXN0LmNvbTp0ZXN0aW5nUGFzc3dvcmQ=' -H 'Connection: keep-alive' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache'
```
1. GET /api/v1/users/balance
    - fetch users balance (assumed BANK merchants are deposits)
```
curl 'http://ec2-54-172-55-155.compute-1.amazonaws.com/api/v1/users/balance' --compressed -H 'Authorization: Basic dGVzdEB0ZXN0LmNvbTp0ZXN0aW5nUGFzc3dvcmQ=' -H 'Connection: keep-alive' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache'
```

1. POST /api/v1/users/authorize
    - authorize a transaction (requires transaction json in body of the request)
```
curl 'http://ec2-54-172-55-155.compute-1.amazonaws.com/api/v1/users/authorize' --compressed -H 'Content-Type: application/json' -H 'Authorization: Basic dGVzdEB0ZXN0LmNvbTp0ZXN0aW5nUGFzc3dvcmQ=' -H 'Connection: keep-alive' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' --data-raw '{"amountInCents":-70}'
```
