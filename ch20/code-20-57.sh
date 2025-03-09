# 코드 20.57 한 줄 실행

curl -XGET -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "admin_user", "password": "test"}' | jq ".access_toekn" | sed "s/\"//g")" localhost:18000/tasks | jq