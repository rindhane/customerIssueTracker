import tomli
import os
import time
from azure.communication.email import EmailClient
from azure.communication.email import EmailContent
from azure.communication.email import EmailMessage
from azure.communication.email import EmailAddress
from azure.communication.email import EmailRecipients

def get_toml_string_from_file(filePath):
    file=open(filePath,"rt")
    stringData=file.read()
    file.close()
    return stringData

class email_container :
    def __init__(self,subject,plain_text, fromAddress, recipient) -> None:
        self.subject=subject
        self.plain_text=plain_text
        self.fromAddress=fromAddress
        self.recipient=recipient
    def getContent(self) :
        return EmailContent(
            subject=self.subject,
            plain_text=self.plain_text
        )
    def getEmailMessage(self):
        recipient_email = EmailAddress(
            email=self.recipient
        )
        recipients=EmailRecipients(
            to=[recipient_email],
        )
        return EmailMessage(
            sender = self.fromAddress,
            content=self.getContent(),
            recipients=recipients
        )


def get_email_Client(cred):
    connection_string = cred["creds"][0]["connection_string"]
    email_client = EmailClient.from_connection_string(connection_string)
    return email_client

def get_cred_store(filePath):
    tomlData=get_toml_string_from_file(filePath)
    cred = tomli.loads(tomlData)
    return cred

def sendEmail (details_dict, cred):
    OTP=details_dict["message"]
    RECIPIENT = details_dict["email"]
    trial_emailMessage = email_container(
        subject="OTP for Login (valid for 5 mins)",
        plain_text=f"Your OTP is {OTP}",
        fromAddress=cred["creds"][0]["mailFrom"],
        recipient=RECIPIENT
    )
    email_client=get_email_Client(cred)
    response = email_client.send(trial_emailMessage.getEmailMessage())
    print(f"response:{response}")
    message_id = response.message_id
    return message_id

def test_sendEmail(message, emailAddr, cred_store):
    try:
        message_id= sendEmail(dict(message=message,email=emailAddr), cred_store)
        email_client=get_email_Client(cred_store)
        counter = 0
        while True:
            counter+=1
            send_status = email_client.get_send_status(message_id)
            if (send_status):
                print(f"Email status for message_id {message_id} is {send_status.status}.")
            if (send_status.status.lower() == "queued" and counter < 12):
                time.sleep(10)  # wait for 10 seconds before checking next time.
                counter +=1
            else:
                if(send_status.status.lower() == "outfordelivery"):
                    print(f"Email delivered for message_id {message_id}.")
                    break
                else:
                    print("Looks like we timed out for checking email send status.")
                    break
    except Exception as ex:
        print('Exception:')
        print(ex)

if __name__ == '__main__':
   filePath='./secrets.toml'
   test_sendEmail("otp:123456", "test@example.com", get_cred_store(filePath))