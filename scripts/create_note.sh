curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMSIsInBhc3N3b3JkIjoicGFzcyJ9fQ.n9JlUV4LJK3-8ID3SGKNKu4eCNmodEl5Qd2aqAX2eOw' \
-H 'Content-Type: application/json' \
-d '
{
    "user_id":1,
    "title":"Some test title",
    "text":"HELLo, world!!!",
    "tags":["hello_world"]
}'
