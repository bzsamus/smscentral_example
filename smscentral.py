import httplib, urllib

gateway = "my.smscentral.com.au"
requestpath = "/wrapper/sms"
username = "username"
password = "password"
recipient = "04xxxxxxxx"
messagetext = "Hello from SMSCentral"

params = urllib.urlencode({
        "USERNAME" : username,
        "PASSWORD" : password,
        "ACTION" : "send",
        "ORIGINATOR" : "shared",
        "RECIPIENT" : recipient,
        "MESSAGE_TEXT" : messagetext
})

headers = {"Content-type": "application/x-www-form-urlencoded",
           "Accept": "text/plain"}

conn = httplib.HTTPConnection(gateway)
conn.request("POST", requestpath, params, headers)

response = conn.getresponse()
data = response.read()
print data

conn.close()
