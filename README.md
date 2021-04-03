# Mail microservice
The mail micro service is used to send out emails, created using GRPC framework

## GRPC methods
Below is a list of GRPC methods in the mail service with their respective input and output.

1. SendTestMail -
   - Input - Name, Email, Message, Phone, Title, Preview.
   - output - The method will sent out a test welcome mail to the input email along with other information. 

## Useful make commands

###### Run the service 
- mage build:run         

###### Compile the service 
- mage build:build

###### Transpile protobuffer definitions 
- mage build:protoc

###### Test the service 
- mage build:test
###### List make targets
- mage                    
```
This microservice has been created using go modules;
```

## Things to configure before using the mail microservice
1. Setup an SMTP email account.
2. Fill out the missing SMTP information in mail/mail.go file:
    - Host Name
    - Port Number
    - Account Email
    - User Name
    - Account Password