curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMSIsInBhc3N3b3JkIjoicGFzczEifX0.zZJAspJ2Nh87XsfTmL9nIw4WEKrHvzFWD5yMbnOf6rI' \
-H 'Content-Type: application/json' \
-d '
{
    "title":"Another test title",
    "text":"Bye bye, cruel world!!!",
    "tags":["hello_world","test","fuck"]
}'
