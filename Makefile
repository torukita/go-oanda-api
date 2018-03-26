.PHONY: swagger requirements

swagger:
	swagger generate client -f ./v20.yaml -a oanda

requirements:
	go get -u github.com/go-swagger/go-swagger
	go get -u github.com/golang/dep/cmd/dep


