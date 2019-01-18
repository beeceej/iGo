#! /bin/bash
# To run this demo, in a terminal:
# $ go run cmd/interpreter/main.go
# In another terminal:
# $ ./demo.sh
readonly sayHiFn=$(cat <<EOF
{
    "raw": "func sayHi() {\\n\\tfmt.Println(\"Hello World, from iGo\")\\n}"
}
EOF
)

readonly sayHiExpression=$(cat <<EOF
{
    "raw": "sayHi()"
}
EOF
)

readonly plainOlGo=$(cat <<EOF
{
  "raw": "fmt.Printf(\"Hello, %s\", \"Brian\")"
}
EOF
)

main() {
  curl -XPOST "localhost:9999/interpret" -d "$sayHiFn"
  curl -XPOST "localhost:9999/interpret" -d "$sayHiExpression"
  curl -XPOST "localhost:9999/interpret" -d "$plainOlGo" 
}

main