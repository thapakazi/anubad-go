export JSON_DATA_TO_IMPORT_FROM=./test/anubad_sample.json
export DBPATH=./test/db
export DBNAME=anubad
export COLNAME=sabda

all:
	@go run main.go

start_mongo_first:
	mkdir -p ${DBPATH}
	mongod --dbpath ${DBPATH} &

import: start_mongo_first
	@mongoimport --db ${DBNAME} --collection ${COLNAME} ${JSON_DATA_TO_IMPORT_FROM}

clean: start_mongo_first
	@mongo -eval "db.${COLNAME}.drop()" ${DBNAME}
	# rm -rf ${DBPATH}
