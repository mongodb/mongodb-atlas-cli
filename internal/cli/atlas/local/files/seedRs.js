try {
   rs.status();
} catch {
   rs.initiate(
      {
         "_id": "rs0",
         "version": 1,
         "configsvr": false,
         "members": [
            { "_id": 0, "host": "mongod1.internal:27017", "horizons": { "external": "localhost:37017" } },
         ]
      }
   );
}
