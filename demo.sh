#! /bin/bash
# To run this demo, in a terminal:
# $ go run cmd/interpreter/main.go
# In another terminal:
# $ ./demo.sh
readonly sayHiFn=$(cat <<EOF
{
    "text": "func sayHi() {\\n\\tfmt.Println(\"Hello World, from iGo\")\\n}\"}"
}
EOF
)

readonly sayHiExpression=$(cat <<EOF
{
    "text": "sayHi()"
}
EOF
)

main() {
  curl -XPOST "http://localhost:9999/interpret" -d "$sayHiFn"
  curl -XPOST "http://localhost:9999/interpret" -d "$sayHiExpression"
}

main