# 코드 20.52 태스크 목록을 확인한다

$ curl -XGET -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/tasks | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   115  100   115    0     0  10835      0 --:--:-- --:--:-- --:--:-- 28750
[
  {
    "id": 68,
    "title": "Implement a handler",
    "status": "todo"
  },
  {
    "id": 69,
    "title": "Implement a handler",
    "status": "todo"
  }
]