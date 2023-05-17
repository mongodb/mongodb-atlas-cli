if (db.system.users.find({"user":"mongoUser"}).count() == 0) {
   db.createUser({
      "user": "mongoUser",
      "pwd": "hunter1",
      "roles": [
         { "role": "root", "db": "admin" }
      ]
   });
}
