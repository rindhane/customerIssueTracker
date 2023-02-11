from datetime import datetime
from flask import Flask, render_template, request, redirect, url_for, send_from_directory
import os 
import json
app = Flask(__name__)
from azureEmail import sendEmail, get_cred_store

def build_cred_store():
   path=os.path.join(app.root_path,"secrets.toml")
   print(path)
   return get_cred_store(path)

CRED_STORE=build_cred_store()

@app.route('/')
def index():
   print('Request for index page received')
   return render_template('index.html')

@app.route('/favicon.ico')
def favicon():
    return send_from_directory(os.path.join(app.root_path, 'static'),
                               'favicon.ico', mimetype='image/vnd.microsoft.icon')


@app.route('/sendOTP', methods=['POST'])
def otpSendBinding():
   #name = request.form.get('name')
   print(request.data)
   if request.data:
      data_dict=json.loads(request.data)
      print('Request for otp received with email=%s & otp=%s' % 
                                 (data_dict["email"],data_dict["otp"])
      )
      sendEmail(dict(message=data_dict["otp"],email=data_dict["email"]), CRED_STORE)
      return json.dumps(dict(status="ok",remark="request received"))
   print('invalid request -- redirecting')
   return redirect(url_for('index'))  

if __name__ == '__main__':
   app.run()