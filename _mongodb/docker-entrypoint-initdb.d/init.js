const env = process.env
var user = {
    user: env.MONGO_INITDB_ROOT_USERNAME,
    pwd: env.MONGO_INITDB_ROOT_PASSWORD,
    roles: [{
        role: "dbOwner",
        db: env.MONGO_INITDB_DATABASE,
    }]
};

db.createUser(user);
db.createCollection(env.MONGO_INITDB_COLLECTION);