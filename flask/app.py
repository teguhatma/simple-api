from flask import Flask


app = Flask(__name__)


@app.route("/api/hello-world", methods=["GET"])
def index():
    return {"msg": "Hello, World!"}


if __name__ == "__main__":
    app.run(port="5000")
