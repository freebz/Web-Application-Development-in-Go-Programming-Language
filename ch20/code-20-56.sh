# 코드 20.56 비관리자 권한의 사용자는 GET /admin에 접근할 수 없다

$ curl -XGET -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/admin
{"message":"not admin"}%