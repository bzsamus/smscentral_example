var request     = require('request'),
    url         = 'https://my.smscentral.com.au/wrapper/sms',
    username    = 'username',
    password    = 'password',
    action      = 'send',
    originator  = 'shared',
    recipient   = '0434390783',
    messagetext = 'Hello from SMSCentral!';


request.post(url, {form:{
            'USERNAME'  : username,
            'PASSWORD'  : password,
            'ACTION'    : action,
            'ORIGINATOR': originator,
            'RECIPIENT' : recipient,
            'MESSAGE_TEXT' : messagetext
        }},function(e, r, body){
            console.log(body);
        });
