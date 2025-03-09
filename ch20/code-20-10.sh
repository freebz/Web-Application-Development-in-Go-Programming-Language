# 코드 20.10 비밀키와 공개키 쌍 생성

openssl genrsa 4096 > secret.pem
openssl rsa -pubout < secret.pem > public.pem