curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsInVzZXIiOnsidXNlcm5hbWUiOiJ1c2VyMiIsInBhc3N3b3JkIjoicGFzczIifX0.RZm_0-ILjOv7FP82I_3Hwre2OPR7HDhQRje0WABnp_4' \
-H 'Content-Type: application/json' \
-d '
{
    "user_id":1,
    "title":"Another test title",
    "text":"Bye bye, cruel world!!!",
    "tags":["hello_world","test","fuck"]
}'
