curl -i -H "Content-Type: application/json" http://localhost:7878/json -d '{"name":"chenchao","pass":"123","age":"22"}'
curl -id "name=chenchao&pass=123" localhost:7878/body?name=ding
