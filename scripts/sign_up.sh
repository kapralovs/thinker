curl -X POST localhost:8080/auth/sign_up \
-H 'Content-Type: application/json' \
-d '{"username":"user1","password":"pass","name":"Jack"}'