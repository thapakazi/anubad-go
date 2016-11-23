export JSON_DATA_TO_IMPORT_FROM=./test/anubad_sample.json
export MONGO_URL=localhost:27017
export DBPATH=./test/db
export DBNAME=anubad
export COLNAME=sabdakosh

all:
	@go run main.go handlers.go sabda.go mgo_wrapper.go

start_mongo_first:
	mkdir -p ${DBPATH}
	mongod --dbpath ${DBPATH} &

import: start_mongo_first
	@mongoimport --db ${DBNAME} --collection ${COLNAME} ${JSON_DATA_TO_IMPORT_FROM}

clean: start_mongo_first
	@mongo -eval "db.${COLNAME}.drop()" ${DBNAME}
	# rm -rf ${DBPATH}

single:
	go run sabda.go
