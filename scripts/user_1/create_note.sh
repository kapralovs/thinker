curl -X POST localhost:8080/note/create \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjYwMDAwMDAwMDAwLCJ1c2VyIjp7ImlkIjoxLCJuYW1lIjoidXNlcjEiLCJ1c2VybmFtZSI6InVzZXIxIiwicGFzc3dvcmQiOiJwYXNzMSIsImN1cnJlbnRfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbGVIQWlPall3TURBd01EQXdNREF3TENKMWMyVnlJanA3SW1sa0lqb3hMQ0p1WVcxbElqb2lkWE5sY2pFaUxDSjFjMlZ5Ym1GdFpTSTZJblZ6WlhJeElpd2ljR0Z6YzNkdmNtUWlPaUp3WVhOek1TSXNJbU4xY25KbGJuUmZkRzlyWlc0aU9pSmxlVXBvWWtkamFVOXBTa2xWZWtreFRtbEpjMGx1VWpWalEwazJTV3R3V0ZaRFNqa3VaWGxLYkdWSVFXbFBhazB5VFVSQmQwMUVRWGROUkVGM1RVUkJjMGx1Vm5wYVdFbHBUMjV6YVdSWVRteGpiVFZvWWxkVmFVOXBTakZqTWxaNVRWTkpjMGx1UW1oak0wNHpZak5LYTBscWIybGpSMFo2WTNwRmFXWllNQzVWY0RRd2NESkdjbmhHY0VaNWVGSTJNM2QwYUVWTFp6TTNkblkwYlc5MVUwTjJSbEZPT0doc1dESlZJbjE5LlBfd2pjU0dEVjZVY1ZMdGt4Qzh5dVlYdE9RY0NiZ0sxdndWdmVlQV9BdzQifX0.vyxIA9mCw3G9mMLz0m2Z03bxmz7Tt061Cswk1P-S1fA' \
-H 'Content-Type: application/json' \
-d '
{
    "title":"Some test title",
    "text":"HELLo, world!!!",
    "tags":["hello_world"]
}'
