
run:
	go run main.go i

# run with input
runwi:
	go run main.go

base:
	cp -i _template/main.go ./main.go

move:
	./movefile.sh

