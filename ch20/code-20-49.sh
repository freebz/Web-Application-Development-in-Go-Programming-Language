# 코드 20.49 틀린 패스워드로 로그인을 시도한다

$ curl -XPOST localhost:18000/login -d '{"user_name": "admin_user", "password": "test?"}'
{"message":"wrong password: crypto/bcrypt: hashedPassword is not the hash of the given password"}%