curl -X POST http://localhost:8080/lists \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My Favorite NBA Players",
    "category": "Basketball Players",
    "autoValidate": true
}'
