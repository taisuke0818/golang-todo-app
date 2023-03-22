mongoimport \
    -u ${MONGO_INITDB_ROOT_USERNAME} \
    -p ${MONGO_INITDB_ROOT_PASSWORD} \
    --db ${MONGO_INITDB_DATABASE} \
    --collection ${MONGO_INITDB_COLLECTION} \
    --file /docker-entrypoint-initdb.d/${MONGO_INITDB_COLLECTION}.json \
    --jsonArray
