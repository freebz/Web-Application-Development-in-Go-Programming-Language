# 코드 20.55 GET /tasks 엔드포인트 호출

$ curl -XGET -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/tasks | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     2  100     2    0     0    194      0 --:--:-- --:--:-- --:--:--   500
[]