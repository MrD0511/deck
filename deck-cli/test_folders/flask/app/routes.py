from flask import render_template
from flask import current_app as app

@app.route('/')
def hello_world():
    return render_template('index.html')
