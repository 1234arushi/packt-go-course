go get -u github.com/gin-gonic/gin


hitting requests on shell ->arushisharma@Admins-MacBook-Air mytodo % curl -X POST http://localhost:8080/tasks \ 
  -H "Content-Type: application/json" \
  -d @payload.json
{"id":1,"title":"Buy groceries","description":"Buy tomatoes,bun,burgers and lettuce","done":false}% 