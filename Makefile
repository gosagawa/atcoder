
go:
	echo "go" > .mode

cpp:
	echo "cpp" > .mode

run:
	./shell/run.sh

# run with input
runwi:
	./shell/runwi.sh

base:
	./shell/base.sh

m:
	./shell/movefile.sh

mc:
	./shell/movefilecontest.sh

entr:
	./shell/entr.sh

d:
	./shell/download.sh

du:
	./shell/downloadurl.sh

t:
	./shell/test.sh

s:
	./shell/submit.sh

su:
	./shell/submiturl.sh

load:
	./shell/load.sh

login:
	oj login https://atcoder.jp

