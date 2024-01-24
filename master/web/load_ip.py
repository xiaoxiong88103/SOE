import json
import os

from flask import jsonify, request


def create_or_load_ip_file():
    ip_file_path = './data/ip.json'
    if not os.path.exists(ip_file_path):
        # 如果文件不存在，创建一个空的 ip.json 文件
        with open(ip_file_path, 'w') as ip_file:
            json.dump({"ip": []}, ip_file)

def load_ip_data():
    ip_file_path = './data/ip.json'
    with open(ip_file_path, 'r') as ip_file:
        return json.load(ip_file)

def save_ip_data(ip_data):
    ip_file_path = './data/ip.json'
    with open(ip_file_path, 'w') as ip_file:
        json.dump(ip_data, ip_file)


def add_ip(ip_to_add):
    ip_data = load_ip_data()
    ip_list = ip_data["ip"]

    if ip_to_add not in ip_list:
        ip_list.append(ip_to_add)
        save_ip_data(ip_data)
        return jsonify({"status": "success", "message": "IP added successfully", "code": 200}), 200
    else:
        return jsonify({"status": "error", "message": "IP already exists", "code": 400}), 400

def del_ip(ip_to_delete):
    ip_data = load_ip_data()
    ip_list = ip_data["ip"]

    if ip_to_delete in ip_list:
        ip_list.remove(ip_to_delete)
        save_ip_data(ip_data)
        return jsonify({"status": "success", "message": "IP deleted successfully", "code": 200}), 200
    else:
        return jsonify({"status": "error", "message": "IP does not exist", "code": 400}), 400

def show_ip():
    ip_data = load_ip_data()
    ip_list = ip_data["ip"]

    return jsonify({"status": "success", "data": ip_list, "code": 200}), 200
