
run:
	go run main.go i

# run with input
runwi:
	go run main.go

base:
	cp -i _template/main.go ./main.go

m:
	./shell/movefile.sh main.go

mcpp:
	./shell/movefile.sh main.cpp

mc:
	./shell/movefilecontest.sh

entr:
	find ./ -maxdepth 1  -name main.go -or -name input | entr -c -p go run main.go i

entrcpp:
	find ./ -maxdepth 1  -name main.cpp -or -name input | entr -c -p ./shell/runcpp.sh

d:
	./shell/download.sh

du:
	./shell/downloadurl.sh

t:
	oj t -c "go run main.go"

login:
	oj login https://atcoder.jp

s:
	./shell/submit.sh main.go

scpp:
	./shell/submit.sh main.cpp
