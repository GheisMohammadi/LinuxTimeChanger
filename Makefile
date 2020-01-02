TMC = ${GOPATH}/src/timechanger

########################################
### make all
all: build

########################################
### Tools & dependencies
deps:

	dep ensure

########################################
### Formatting, linting, and vetting
fmt:
	@go fmt ./...

########################################
### building
build:
	@go build tmc.go
	#@go build client/client.go

run:
	@go run tmc.go

deploy:

	#copy to node 1
	sshpass -p "olUjsUjks22@ty&ksh$$jsa!" scp /home/gheis/goApps/src/timechanger/tmc  root@167.71.205.13:~/tmc
	sshpass -p "olUjsUjks22@ty&ksh$$jsa!" scp /home/gheis/goApps/src/timechanger/disableautotime.sh  root@167.71.205.13:~/disableautotime.sh 

	#copy to node 2
	sshpass -p "ikjsd@kjds$$kjs#ljslkjs365" scp /home/gheis/goApps/src/timechanger/tmc  root@167.71.205.13:~/tmc
	sshpass -p "ikjsd@kjds$$kjs#ljslkjs365" scp /home/gheis/goApps/src/timechanger/disableautotime.sh  root@167.71.205.13:~/disableautotime.sh 

	#copy to node 3
	sshpass -p "ksjdh@lkjs&lkapo198$$ks" scp /home/gheis/goApps/src/timechanger/tmc  root@167.71.205.13:~/tmc
	sshpass -p "ksjdh@lkjs&lkapo198$$ks" scp /home/gheis/goApps/src/timechanger/disableautotime.sh  root@167.71.205.13:~/disableautotime.sh 

	#copy to node 4
	sshpass -p "lskdo#loytkahhf$$mdhkpa!" scp /home/gheis/goApps/src/timechanger/tmc  root@167.71.205.13:~/tmc
	sshpass -p "lskdo#loytkahhf$$mdhkpa!" scp /home/gheis/goApps/src/timechanger/disableautotime.sh  root@167.71.205.13:~/disableautotime.sh 

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: tools deps
.PHONY: build 
.PHONY: fmt metalinter