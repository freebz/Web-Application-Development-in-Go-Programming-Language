# 코드 19.5 동작 확인

curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 31 May 2022 05:43:37 GMT
Content-Length: 9

{"id":32}%

curl -i -XGET localhost:18000/tasks
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 31 May 2022 05:43:39 GMT
Content-Length: 113

[{"id":31, "title":"Implement a handler","status":"todo"},{"id":32, "title":"Implement a handler", "status":"ststus":"todo"}]%
