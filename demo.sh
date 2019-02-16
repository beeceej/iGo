#! /bin/bash
# To run this demo, in a terminal:
# $ go run cmd/interpreter/main.go
# In another terminal:
# $ ./demo.sh
readonly sayHiFn=$(cat <<EOF
{
    "raw": "func sayHi() {\\n\\tfmt.Println(\"Hello World, from iGo\")\\r}"
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

readonly testRemoveByteSlice=$(cat <<EOF
{
  "raw": "func testRemoveByteSlice() {\n\tfrom := 2\n\ttext := []byte{0, 1, 2, 3, 4, 5, 6, 7}\n\tto := 5\n\n\tsize := to - from\n\tcopy(text[from:], text[to:])\n\ttext = text[:len(text)-size]\n\tfmt.Println(text)\n}\n"
}
EOF
)

readonly diagFunc=$(cat <<EOF
{
  "raw": "func byte_slice_insert(text []byte, offset int, what []byte) []byte {\n\tn := len(text) + len(what)\n\tfmt.Println(\"n is\", n)\n\tfmt.Println(\"before growing, text is\", text)\n\ttext = byte_slice_grow(text, n)\n\tfmt.Println(\"after growing, text is\", text)\n\ttext = text[:n] // make size == to capacity?\n\tfmt.Println(\"after text = text[:n] // make size == to capacity? text is\", text)\n\tcopy(text[offset+len(what):], text[offset:])\n\tfmt.Println(\"after copy(text[offset+len(what):], text[offset:]) text is\", text)\n\tcopy(text[offset:], what)\n\tfmt.Println(\"after copy(text[offset:], text) text is\", text)\n\treturn text\n}"
}
EOF
)

readonly byteslicegrow=$(cat <<EOF
{
  "raw": "\n\n\n\nfunc byte_slice_grow(s []byte, desired_cap int) []byte {\n\tif cap(s) < desired_cap {\n\t\tns := make([]byte, len(s), desired_cap)\n\t\tcopy(ns, s)\n\t\treturn ns\n\t}\n\treturn s\n}"
}
EOF
)






main() {
  curl -XPOST "localhost:9999/interpret" -d "$sayHiFn"
  curl -XPOST "localhost:9999/interpret" -d "$sayHiExpression"
  curl -XPOST "localhost:9999/interpret" -d "$plainOlGo" 
  curl -XPOST "localhost:9999/interpret" -d "$testRemoveByteSlice"
  curl -XPOST "localhost:9999/interpret" -d "$diagFunc"
  curl -XPOST "localhost:9999/interpret" -d "$byteslicegrow"
  
}

main