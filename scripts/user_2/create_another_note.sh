curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMiIsInBhc3N3b3JkIjoicGFzMiJ9fQ.0vBp32Mn-IUsZFnPk46XbDYPrATdd7a4Ify--givvfg' \
-H 'Content-Type: application/json' \
-d '
{
    "title":"Another test title",
    "text":"Bye bye, cruel world!!!",
    "tags":["hello_world","test","fuck"]
}'
