const express = require("express");

const app = express();

app.get("/api/hello-world", (req, res) => {
  res.send({ msg: "Hello, World!" });
});

app.listen(5000, console.log("App running on Port: 5000"));
