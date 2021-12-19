
run:
	go run main.go i

# run with input
runwi:
	go run main.go

base:
	cp -i _template/main.go ./main.go

move:
	./movefile.sh

entr:
	find ./ -maxdepth 1  -name main.go -or -name input | entr -c -p go run main.go i

