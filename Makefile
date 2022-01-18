
run:
	go run main.go i

# run with input
runwi:
	go run main.go

base:
	cp -i _template/main.go ./main.go

mv:
	./shell/movefile.sh

entr:
	find ./ -maxdepth 1  -name main.go -or -name input | entr -c -p go run main.go i

dl:
	./shell/download.sh

t:
	oj t -c "go run main.go"

login:
	oj login https://atcoder.jp

s:
	oj t -c "go run main.go" ; ./shell/submit.sh
