up:
	cd ./sql/schema/ && GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=../../conduit.db goose up

down:
	cd ./sql/schema/ && GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=../../conduit.db goose down

status: 
	cd ./sql/schema/ && GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=../../foo.db goose status

reset: 
	cd ./sql/schema/ && GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=../../foo.db goose reset
