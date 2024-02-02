from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/', methods=['GET'])
def index():
    return jsonify(message=f"Hello from API server!")

@app.route('/livez', methods=['GET'])
def live():
    return jsonify(message=f"Api is Live!")

@app.route('/readyz', methods=['GET'])
def ready():
    return jsonify(message=f"Api is Ready")

if __name__ == '__main__':
    app.run()
