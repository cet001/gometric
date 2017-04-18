all : install

clean :
	@echo ">>> Cleaning and initializing gometric project <<<"
	@go clean
	@gofmt -w .
	@go get github.com/cet001/mathext
	@go get github.com/stretchr/testify

test : clean
	@echo ">>> Running unit tests <<<"
	@go test ./ ./strdist

test-coverage : clean
	@echo ">>> Running unit tests and calculating code coverage <<<"
	@go test ./ ./strdist -cover

install : test
	@echo ">>> Building and installing gometric <<<"
	@go install
	@echo ">>> gometric installed successfully! <<<"
	@echo ""
