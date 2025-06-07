from flask import Flask, request, jsonify
from werkzeug.security import generate_password_hash, check_password_hash

app = Flask(__name__)

users = {}
servers = []

@app.route('/')
def hello_world():
    return 'Hello, World! This is the backend.'

@app.route('/register', methods=['POST'])
def register():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    if not username or not password:
        return jsonify({'message': 'Username and password are required'}), 400

    if username in users:
        return jsonify({'message': 'User already exists'}), 400

    users[username] = generate_password_hash(password)
    return jsonify({'message': 'User registered successfully'}), 201

@app.route('/login', methods=['POST'])
def login():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    if not username or not password:
        return jsonify({'message': 'Username and password are required'}), 400

    user_password_hash = users.get(username)
    if not user_password_hash or not check_password_hash(user_password_hash, password):
        return jsonify({'message': 'Invalid username or password'}), 401

    return jsonify({'message': 'Login successful'}), 200

@app.route('/add_server', methods=['POST'])
def add_server():
    data = request.get_json()
    name = data.get('name')
    ip = data.get('ip')

    if not name or not ip:
        return jsonify({'message': 'Server name and IP are required'}), 400

    servers.append({'name': name, 'ip': ip})
    return jsonify({'message': 'Server added successfully'}), 201

@app.route('/list_servers', methods=['GET'])
def list_servers():
    return jsonify(servers), 200

@app.route('/server_status/<server_ip>', methods=['GET'])
def server_status(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if server:
        return jsonify({
            "ip": server_ip,
            "cpu_usage": "15%",
            "memory_usage": "45%",
            "disk_usage": "60%"
        }), 200
    else:
        return jsonify({"error": "Server not found"}), 404

@app.route('/upload/<server_ip>', methods=['POST'])
def upload_file(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    if 'file' not in request.files:
        return jsonify({'message': 'No file part in the request'}), 400

    file = request.files['file']
    if file.filename == '':
        return jsonify({'message': 'No selected file'}), 400

    return jsonify({"message": f"File '{file.filename}' successfully uploaded to {server_ip} (simulated)"}), 200

@app.route('/download/<server_ip>/<path:filename>', methods=['GET'])
def download_file(server_ip, filename):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({"message": f"File '{filename}' successfully downloaded from {server_ip} (simulated)"}), 200

@app.route('/list_files/<server_ip>/<path:directory_path>', methods=['GET'])
def list_files(server_ip, directory_path):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "path": directory_path,
        "files": [
            {"name": "file1.txt", "type": "file", "size": "1KB"},
            {"name": "report.docx", "type": "file", "size": "120KB"}
        ],
        "directories": [
            {"name": "subdir1", "type": "directory"},
            {"name": "another_dir", "type": "directory"}
        ]
    }), 200

@app.route('/list_databases/<server_ip>', methods=['GET'])
def list_databases(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "server_ip": server_ip,
        "databases": [
            {"name": "mysql_db1", "type": "MySQL", "size": "150MB"},
            {"name": "postgres_db1", "type": "PostgreSQL", "size": "200MB"}
        ]
    }), 200

@app.route('/list_tables/<server_ip>/<database_name>', methods=['GET'])
def list_tables(server_ip, database_name):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "server_ip": server_ip,
        "database_name": database_name,
        "tables": ["users", "products", "orders", "inventory_items"]
    }), 200

@app.route('/execute_query/<server_ip>/<database_name>', methods=['POST'])
def execute_query(server_ip, database_name):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    data = request.get_json()
    query = data.get('query')

    if not query:
        return jsonify({'message': 'Query is required'}), 400

    if query.lower().startswith("select"):
        return jsonify({
            "server_ip": server_ip,
            "database_name": database_name,
            "query": query,
            "columns": ["id", "name", "email"],
            "rows": [
                [1, "Alice Wonderland", "alice@example.com"],
                [2, "Bob The Builder", "bob@example.com"]
            ],
            "message": "SELECT query executed successfully (simulated)."
        }), 200
    else:
        return jsonify({
            "server_ip": server_ip,
            "database_name": database_name,
            "query": query,
            "rows_affected": 3, # Mocked
            "message": "Query executed successfully (simulated)."
        }), 200

@app.route('/execute_command/<server_ip>', methods=['POST'])
def execute_command(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    data = request.get_json()
    command = data.get('command')

    if not command:
        return jsonify({'message': 'Command is required'}), 400

    output = ""
    error = ""
    exit_code = 0

    if command == "ls -l /tmp":
        output = "total 0\ndrwxrwxrwt 2 root root 40 Oct 26 10:00 some_dir\n-rw-r--r-- 1 user user 0 Oct 26 10:00 some_file.txt"
    elif command == "whoami":
        output = "mock_user"
    elif command == "pwd":
        output = "/home/mock_user"
    else:
        error = f"command not found: {command.split()[0] if command else ''}"
        exit_code = 127

    return jsonify({
        "server_ip": server_ip,
        "command": command,
        "output": output,
        "error": error,
        "exit_code": exit_code
    }), 200

@app.route('/docker/<server_ip>/images', methods=['GET'])
def docker_images(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "server_ip": server_ip,
        "images": [
            {"id": "sha256:f707a09c95418991968b54a55c513ac29da628f063e095858841914b89a89f60", "repository": "ubuntu", "tag": "latest", "size": "72.8MB"},
            {"id": "sha256:0d17b0b84fd67817692764c1e8260502682900931afe0008b93f2969eba74747", "repository": "nginx", "tag": "stable", "size": "133MB"},
            {"id": "sha256:e9bbb59c3673f9f0f018a968934265108d86305db899946f02a5b895f056349a", "repository": "alpine", "tag": "latest", "size": "5.53MB"}
        ]
    }), 200

@app.route('/docker/<server_ip>/containers', methods=['GET'])
def docker_containers(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "server_ip": server_ip,
        "containers": [
            {"id": "c1a2b3c4d5e6", "name": "web_server_1", "image": "nginx:stable", "status": "Up 2 hours", "ports": "0.0.0.0:80->80/tcp"},
            {"id": "d4e5f6g7h8i9", "name": "app_db_1", "image": "postgres:13", "status": "Exited (0) 5 minutes ago", "ports": ""},
            {"id": "j0k1l2m3n4p5", "name": "my_redis", "image": "redis:alpine", "status": "Up 1 day", "ports": "0.0.0.0:6379->6379/tcp"}
        ]
    }), 200

@app.route('/docker/<server_ip>/pull_image', methods=['POST'])
def docker_pull_image(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    data = request.get_json()
    image_name = data.get('image_name')

    if not image_name:
        return jsonify({'message': 'Image name is required'}), 400

    return jsonify({"message": f"Image '{image_name}' is being pulled on {server_ip} (simulated)."}), 200

@app.route('/docker/<server_ip>/container_logs/<container_id>', methods=['GET'])
def docker_container_logs(server_ip, container_id):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    return jsonify({
        "server_ip": server_ip,
        "container_id": container_id,
        "logs": "[INFO] Starting up Nginx server...\n[INFO] Listening on port 80.\n[ACCESS] 192.168.1.10 - GET /index.html\n[ACCESS] 192.168.1.12 - GET /styles.css"
    }), 200

@app.route('/install_software/<server_ip>', methods=['POST'])
def install_software(server_ip):
    server = next((s for s in servers if s['ip'] == server_ip), None)
    if not server:
        return jsonify({"error": "Server not found"}), 404

    data = request.get_json()
    package_name = data.get('package_name')
    version = data.get('version', 'latest') # Default to 'latest'

    if not package_name:
        return jsonify({'message': 'Package name is required'}), 400

    return jsonify({
        "server_ip": server_ip,
        "package_name": package_name,
        "version": version,
        "message": f"Installation of {package_name} (version: {version}) initiated on {server_ip} (simulated). Check status using task ID 'mock_task_123'.",
        "status": "pending",
        "task_id": "mock_task_123"
    }), 200

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
