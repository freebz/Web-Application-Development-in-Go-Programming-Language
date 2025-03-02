# 코드 17.12 동작 확인

$ curl -i -XGET localhost:18000/tasks
HTTP/1.1 200 OK
Context-Type: application/json; charset=utf-8
Date: Wed, 19 Feb 2025 14:11:52 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

[]%
$ curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
HTTP/1.1 200 OK
Context-Type: application/json; charset=utf-8
Date: Wed, 19 Feb 2025 14:14:17 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8

{"id":1}%
$ curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/bad_req.json.golden
HTTP/1.1 400 Bad Request
Context-Type: application/json; charset=utf-8
Date: Wed, 19 Feb 2025 14:16:28 GMT
Content-Length: 90
Content-Type: text/plain; charset=utf-8

{"message":"Key: 'Title' Error:Field validation for 'Title' failed on the 'required' tag"}%
$ curl -i -XGET localhost:18000/tasks
HTTP/1.1 200 OK
Context-Type: application/json; charset=utf-8
Date: Wed, 19 Feb 2025 14:22:28 GMT
Content-Length: 56
Content-Type: text/plain; charset=utf-8

[{"id":1,"title":"Implement a handler","status":"todo"}]%