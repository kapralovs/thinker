curl --request PUT localhost:8080/note/edit/2 \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMiIsInBhc3N3b3JkIjoicGFzMiJ9fQ.0vBp32Mn-IUsZFnPk46XbDYPrATdd7a4Ify--givvfg' \
-d '
{
    "title":"Edited test note",
    "text":"Some edited text",
    "tags":["edited"]
}'
