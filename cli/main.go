package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"buf.build/gen/go/riza/transform/connectrpc/go/v1/transformv1connect"
	transform "buf.build/gen/go/riza/transform/protocolbuffers/go/v1"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	ctx := context.Background()
	client := transformv1connect.NewTransformerClient(
		http.DefaultClient,
		"http://localhost:8089",
		// "https://sqlc-transform-4mb2ctq5aq-uw.a.run.app",
	)
	log.Println("Uploading binary...")

	blob, err := os.ReadFile("handler.wasm")
	if err != nil {
		log.Fatal(err)
	}

	upload, err := client.UploadBinary(
		context.Background(),
		connect.NewRequest(&transform.UploadBinaryRequest{
			Namespace: "transform-go",
			Code:      blob}),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Transforming event...")
	res, err := client.TransformEvent(ctx, connect.NewRequest(
		&transform.TransformEventRequest{
			Namespace: "transform-go",
			CodeId:    upload.Msg.Id,
			Request: &transform.Request{
				Body: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiZGF0YSI6IntcImRhdGFcIjogXCJ7fVwifSIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.SiquQfOZUkIDlgJJDkvxprw1LjK__0MwoOA4heMkuQ8",
				Headers: map[string]string{
					"Content-Length": "1",
				},
			},
			Env: map[string]string{
				"TEST": "foo",
			},
		},
	))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(protojson.Format(res.Msg.Request))
}

// function main() {
//   const client = createPromiseClient(Transformer, transport);
//
//   const handler = process.argv[2];
//
//   console.log("Compiling code...");
//   const compile = await client.compileCode({
//     namespace: "example",
//     code: fs.readFileSync(handler),
//   })
//
//   const request = JSON.parse(fs.readFileSync("request.json"));
//
//   console.log("Transforming event...");
//   const res = await client.transformEvent({
//     namespace: "example",
//     codeId: compile.id,
//     request: request,
//     env: {
//         "TEST": "foo",
//     },
//   });
//
//   console.log(res);
// }
