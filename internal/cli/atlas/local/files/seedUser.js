db.createUser({
   "user": "mongoUser",
   "pwd": "hunter1",
   "roles": [
      { "role": "root", "db": "admin" }
   ]
});
