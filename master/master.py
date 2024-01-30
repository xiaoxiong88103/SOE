import json
from datetime import timedelta
from flask import Flask, jsonify, request
from influxdb_client import InfluxDBClient
from flask_jwt_extended import JWTManager, jwt_required, create_access_token, get_jwt_identity

from master.web.load_ip import add_ip, del_ip, show_ip

app = Flask(__name__)

def load_config(config):
    # 从配置文件加载配置信息
    with open('config/'+config, 'r') as config_file:
        config = json.load(config_file)
    return config

webconfig = load_config('web.json')
# 配置 JWT
app.config['JWT_SECRET_KEY'] = webconfig.get('token_key')
app.config['JWT_ACCESS_TOKEN_EXPIRES'] = timedelta(minutes=webconfig.get('token_time'))  # 设置令牌有效期为30分钟
jwt = JWTManager(app)


# InfluxDB 连接信息
configjson=load_config('config.json')
bucket = configjson.get('databases')
# 创建 InfluxDB 客户端
client = InfluxDBClient(url=configjson.get('url'), token=configjson.get('token'), org=configjson.get('org'))

@app.route('/query_data', methods=['POST'])
@jwt_required()
def query_data():
    try:
        ip = request.json.get('ip', '')
        if ip == '':
            return jsonify({"code":403}),403

        # 创建查询
        query = f'from(bucket:"{bucket}") |> range(start: -1d) |> filter(fn: (r) => r._measurement == "system_info" and r.ip == "{ip}")'

        # 执行查询
        query_api = client.query_api()
        result = query_api.query(query, org=configjson.get('org'))

        # 解析查询结果
        data = []
        for table in result:
            for record in table.records:
                # 获取字段和值
                fields = record.values
                data.append(fields)

        return jsonify({"status": "success", "data": data,"code":200}),200

    except Exception as e:
        return jsonify({"status": "error", "message": str(e),"code":500}),500


@app.route('/add_ip', methods=['POST'])
@jwt_required()
def add_ip_route():
    data = request.json
    ip_to_add = data.get("ip")
    return add_ip(ip_to_add)

@app.route('/del_ip', methods=['POST'])
@jwt_required()
def del_ip_route():
    data = request.json
    ip_to_del = data.get("ip")
    return del_ip(ip_to_del)


@app.route('/show_ip', methods=['GET'])
@jwt_required()
def show_ip_route():
    return show_ip()



@app.route('/login', methods=['POST'])
def login():
    # 获取请求中的 JSON 数据
    data = request.json

    # 在这里处理登录逻辑，假设用户名和密码正确
    config_username = webconfig.get("username")
    config_password = webconfig.get("password.ini")

    username = data.get("username")
    password = data.get("password.ini")

    # 如果用户名和密码验证成功，生成 token
    if username == config_username and password == config_password:
        access_token = create_access_token(identity=username)
        response_data = {"message": "Login successful", "username": username, "token": access_token,"code":200}
        return jsonify(response_data),200
    else:
        return jsonify({"error": "Invalid credentials"}), 401


@app.route('/update_put_time', methods=['POST'])
@jwt_required()
def update_put_time():
    try:
        data = request.json
        new_put_time = data.get('new_put_time')

        if new_put_time is None:
            return jsonify({"status": "error", "message": "Missing 'new_put_time' in the request body", "code": 400}), 400

        # 读取原始配置
        with open('config/client_config.json', 'r') as config_file:
            config_data = json.load(config_file)

        # 修改put_time值
        config_data['put_time'] = new_put_time

        # 写入新配置
        with open('config/client_config.json', 'w') as config_file:
            json.dump(config_data, config_file, indent=2)

        return jsonify({"status": "success", "message": "put_time updated successfully", "code": 200}), 200

    except Exception as e:
        return jsonify({"status": "error", "message": str(e), "code": 500}), 500



@app.route('/test', methods=['GET'])
def api_test():
    return jsonify({"status": "OK"})



if __name__ == '__main__':
    app.run()
