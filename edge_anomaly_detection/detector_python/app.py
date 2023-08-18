from flask import Flask
from flask import render_template
from flask import request
from pickle import dump, load
import os

app = Flask(__name__)
port = os.environ["PORT"]
print(port)


model = load(open("cicids_model", "rb"))


@app.route("/")
@app.route("/hello", methods=["POST"])
def hello():
    return "Hello, World!"


@app.route("/check", methods=["POST"])
def check():
    print(request.data)
    income_data = request.form.get("features", None)
    if income_data != None:
        income_data = income_data.replace("[", "").replace("]", "").split(",")
        income_data = [float(i) for i in income_data]
        return f"{model.predict([income_data])}"
    else:
        return "ppp"
    # return RF.predict([income_data])


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=port)
