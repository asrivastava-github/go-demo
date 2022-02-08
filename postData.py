import requests

url = 'http://localhost:5000/users'
userData = {
    'firstname': 'Avinash',
    'lastname': 'Srivastava'
}

# x = requests.post(url, data=userData)
y = requests.get('{}/{}'.format(url, 2))

print(y.text)