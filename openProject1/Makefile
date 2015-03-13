NAME = openProject1

SERV = install openProject1/gameServer/

CLIENT = install openProject1/gameClient/

COMM = build openProject1/graphic_server_communication/

EVENT = build openProject1/eventManager

GO = go

all:
	$(GO) $(COMM)
	$(GO) $(EVENT)
	$(GO) $(CLIENT)
	$(GO) $(SERV)

serv: all
	gameServer

client: all
	gameClient

push:
	git add -u && git commit -m "update" && git push origin master

pull:
	git add -u ; git commit -m "update" ; git pull origin master
	$(GO) $(COMM)
	$(GO) $(EVENT)
	$(GO) $(CLIENT)
	$(GO) $(SERV)
