curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMSIsInBhc3N3b3JkIjoicGFzcyJ9fQ.ZTIkUeAglpVrQTqZDXqBXLrCKHbEehjOjSyPBaZRADo' \
-H 'Content-Type: application/json' \
-d '
{
    "user_id":1,
    "title":"Another test title",
    "text":"Bye bye, cruel world!!!",
    "tags":["hello_world","test","fuck"]
}'