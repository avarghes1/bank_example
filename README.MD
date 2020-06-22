# Bank Example

The following is a sample bank api.

[Sample bank api](http://ec2-54-172-55-155.compute-1.amazonaws.com)

The application uses basic auth for all api except user creation.
The database has been seeded with the provided csv file.
The api is on an ec2 instance. I have no monitoring and alerts on it. So if its down, please
dont hesitate to contact me.

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

4. GET /api/v1/users/balance
    - fetch users balance (assumed BANK merchants are deposits)
```
curl 'http://ec2-54-172-55-155.compute-1.amazonaws.com/api/v1/users/balance' --compressed -H 'Authorization: Basic dGVzdEB0ZXN0LmNvbTp0ZXN0aW5nUGFzc3dvcmQ=' -H 'Connection: keep-alive' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache'
```

5. POST /api/v1/users/authorize
    - authorize a transaction (requires transaction json in body of the request)
```
curl 'http://ec2-54-172-55-155.compute-1.amazonaws.com/api/v1/users/authorize' --compressed -H 'Content-Type: application/json' -H 'Authorization: Basic dGVzdEB0ZXN0LmNvbTp0ZXN0aW5nUGFzc3dvcmQ=' -H 'Connection: keep-alive' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' --data-raw '{"amountInCents":-70}'
```

# Dev

1. `docker-compose up`

  - The db will be empty. I have not added scripts to seed the db due to time constraints.
    (I do have a volume I can zip if needed)

# Unit Test

I have added one unit test. Given more time I would have added more tests.
In order to run tests in the test folder run the command,

```
ginkgo -r
```

# Validation

Again I did not complete all validation due to time constraints. But create and update user
has validation.

# Bonus Questions

I did not have time to provide the api for the bonus questions.
However, I checked query performance for merchant summary and it seem to be ok.
Provided more time I would have been able to finish them off.

# Things Missing

1. Logging
  - I added some logs messages but not with appropriate log messages.

2. Metrics
  - I have only one place where I measures time taken. Prometheus would be nice
  to collect metrics

3. Slow logs on DB

4. Monitoring and Alerts
  - The api is on an ec2 instance. I have no monitoring and alerting setup
  - Its missing load balancer
  - The database is running on docker and not rds
