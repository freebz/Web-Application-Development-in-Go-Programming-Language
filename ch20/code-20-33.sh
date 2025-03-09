# 코드 20.33 로그인 성공 시의 응답

$ curl -XPOST localhost:18000/login -d '{"user_name": "freebz", "password": "test"}' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1012  100   971  100    41   5017    211 --:--:-- --:--:-- --:--:--  5411
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE0NDc1OTYsImlhdCI6MTc0MTQ0NTc5NiwiaXNzIjoiZ2l0aHViLmNvbS9mcmVlYnovZ29fdG9kb19hcHAiLCJqdGkiOiIzNTYzOTI2Yy0wMTA0LTQzMjItOGE3NC01MDFlOWIwOWY4MGUiLCJyb2xlIjoiYWRtaW4iLCJzdWIiOiJhY2Nlc3NfdG9rZW4iLCJ1c2VyX25hbWUiOiJqb2huIn0.iARuwWzE0CnsbjZ4-Sw7JrtZ4zIFhWgaN8dtybAGkNaMIaF1rXVtZUmSXVHhw50yNeDwzCSITG4pEcnbhe9EoPwyb0cxa3yL8TC3d-8MKk_4Hgr7mWgYvBxVrxstfZMU4gKDYSllJdZsOq2rLyRUCFzYD1tg-EI_yvFRbf-zJECcgHsz7amlRz6wS2vuJT7h9BTaOdK20NixLKGwoNft1QdHOMfWv0xtTrhskFFFeGCkKksN5XcW_opXRra-BhxTUFIsbHcwGiTy6RzZnCCfgpRDKUCigx7eks5_JwnKju__RO5Vq-eEJP6AvYH6c2V2B1_oWSehjCKIulZlhYWXt2QVTsc-1FA-FLHXM-2Uu-ZhfBsgSrB9IgYUjxKr6HxSQRJxoXLDjP8SOb13P9hV-ajbozPVDXZ8dqiGT9DvsqH_-8IBhQd6mabdBA5egRwI-u2Dm5mqTwLmPYjrSwh2eDFKMjq0GTjdIWgnLanME_ejH0MY8GcFOat-BQoFI6pw3Yrioho9Qz_N0E8CPKZFlgJ_9MyTuFvjV0OaGUoch39vY8giNYfxtLJgBXmacPi0D0rMT5Vr0mtJ10OAcECpTwemzvQRyk-rdMUCTg5Yq7ASzspN_wWV3wmOQOLNSHjBQ8pGXxSCH91f1Jz_tlxki7-cGsoWLOXRSBI3RxIrSDs"
}