.PHONY: generate deploy clean test

clean:
	-rm -rf build model

generate: clean
	smithy build hello.smithy
	mv build/smithy/source/go-codegen model
	echo "package model\nvar goModuleVersion = \"0.0.1\"\n" > model/version.go

deploy: generate
	sam build
	sam deploy

test: generate
	go run client.go $$(aws cloudformation describe-stacks --stack-name smithy-lambda-url-example --query "Stacks[].Outputs[].OutputValue" --output text)
