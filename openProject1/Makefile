NAME = openProject1

SERV = install openProject1/gameServer/

CLIENT = install openProject1/gameClient/

COMM = build openProject1/graphic_server_communication/

EVENT = build openProject1/eventManager

GO = go

all:
	$(GO) $(SERV)
	$(GO) $(CLIENT)
	$(GO) $(COMM)
	$(GO) $(EVENT)

serv: all
	gameServer

client: all
	gameClient