post:
    curl \
        -X POST http://localhost:8080/heapsort \
        -H "Content-Type: application/json" \
        -d '{"numbers": [64, 34, 25, 12, 22, 11, 90]}'

test:
    go test ./internal