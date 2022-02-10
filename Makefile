
run:
	go run main.go i

# run with input
runwi:
	go run main.go

base:
	cp -i _template/main.go ./main.go

m:
	./shell/movefile.sh

entr:
	find ./ -maxdepth 1  -name main.go -or -name input | entr -c -p go run main.go i

d:
	./shell/download.sh

du:
	./shell/downloadurl.sh

t:
	oj t -c "go run main.go"

login:
	oj login https://atcoder.jp

s:
	./shell/submit.sh
