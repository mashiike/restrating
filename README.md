# restrating
[![CircleCI](https://circleci.com/gh/mashiike/restrating.svg?style=svg)](https://circleci.com/gh/mashiike/restrating)

REST full Rating microservice

now development...

```
[github.com/mashiike/rastrating/cmd/api]$ go run .
[restrating] 04:43:47 HTTP "CreatePlayer" mounted on POST /v1/players
[restrating] 04:43:47 HTTP "ApplyMatch" mounted on POST /v1/matches
[restrating] 04:43:47 HTTP server listening on "localhost:8088"
```

/v1/players
```
$ curl -X POST -H 'Content-Type:application/json' -d '{"name":"goat"}' http://localhost:8088/v1/playe
rs | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    41  100    26  100    15   3991   2302 --:--:-- --:--:-- --:--:--  4333
{
  "rrn": "rrn:player:goat"
}
```

/v1/matches
```
$ curl -X POST -H 'Content-Type:application/json' -d '{"scores":{"rrn:player:sheep":1.0, "rrn:player:goat":0.0}}' http://localhost:8088/v1/matches | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   274  100   216  100    58  38883  10441 --:--:-- --:--:-- --:--:-- 43200
{
  "participants": [
    {
      "rrn": "rrn:player:sheep",
      "rating": {
        "strength": 1664.29,
        "lower": 1080.13,
        "upper": 2248.45
      }
    },
    {
      "rrn": "rrn:player:goat",
      "rating": {
        "strength": 1335.7,
        "lower": 751.5400000000001,
        "upper": 1919.8600000000001
      }
    }
  ]
}
```
