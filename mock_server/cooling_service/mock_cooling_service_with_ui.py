from flask import Flask, jsonify, request
from flask_swagger_ui import get_swaggerui_blueprint

app = Flask(__name__)

# === OpenAPI UI ===
SWAGGER_URL = '/docs'
API_URL = '/openapi.yaml'
swaggerui_blueprint = get_swaggerui_blueprint(
    SWAGGER_URL,
    API_URL,
    config={'app_name': "CoolingService Mock API"}
)
app.register_blueprint(swaggerui_blueprint, url_prefix=SWAGGER_URL)

# === Mock endpoints ===
@app.route('/cooling/validate', methods=['POST'])
def validate():
    data = request.json
    return jsonify({
        "is_valid": True,
        "expires_at": "2025-04-18T15:00:00Z",
        "message": f"Cooling period valid for credit_id: {data.get('credit_id')}"
    })

@app.route('/cooling/withdraw', methods=['POST'])
def withdraw():
    data = request.json
    return jsonify({
        "status": "accepted",
        "message": f"Withdrawal accepted for credit_id: {data.get('credit_id')}"
    })

@app.route('/cooling/status', methods=['GET'])
def status():
    return jsonify({
        "credit_id": request.args.get('credit_id'),
        "status": "pending",
        "updated_at": "2025-04-17T10:00:00Z"
    })

@app.route('/cooling/report', methods=['GET'])
def report():
    return jsonify([{
        "credit_id": "CR123456",
        "client_id": "CL123",
        "status": "completed",
        "submitted_at": "2025-04-17T08:00:00Z",
        "resolved_at": "2025-04-17T09:00:00Z"
    }])

# Статический маршрут к YAML (Swagger UI его подгружает)
@app.route('/openapi.yaml', methods=['GET'])
def serve_yaml():
    with open("openapi.yaml", "r", encoding="utf-8") as f:
        return f.read(), 200, {'Content-Type': 'application/yaml'}

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)