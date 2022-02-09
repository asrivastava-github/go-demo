from flask import Flask, json, request

users = [{"ID": 1, "FirstName": "Avi", "LastName": "Srivastava"}, {"ID": 2, "FirstName": "Avinash", "LastName": "Srivastava"}, {"ID": 3, "FirstName": "Avi", "LastName": "Sriva"}]

api = Flask(__name__)

@api.route('/users', methods=['GET', 'POST', 'PATCH'])
@api.route('/users/<id>', methods=['GET', 'DELETE'])
def get_users(id=None):
    if request.method == 'GET' and id:
        for user in users:
            if user['ID'] == int(id):
                return json.dumps(users[users.index(user)])

    if request.method in ['POST', 'PATCH']:
        reqData = request.form
        if request.data:
            reqData = json.loads(request.data)
        id = max([int(user["ID"]) for user in users]) + 1
        print(request.data)
        print(id)
        firstname = reqData['firstname']
        lastname = reqData['lastname']
        userData = {
            "ID": id,
            "FirstName": firstname,
            "LastName": lastname
        }
        print(userData)
        users.append(userData)
        return json.dumps(users)

    if request.method == 'DELETE' and id:
        for user in users:
            if user['ID'] == int(id):
                del_index = users.index(user)
                del users[del_index]
                return json.dumps(users)
    return json.dumps(users)


if __name__ == '__main__':
    api.run()