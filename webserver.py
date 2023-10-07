from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

@app.route('/newfile')
def newfile():
    return 'This is a new file page.'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=4000)
