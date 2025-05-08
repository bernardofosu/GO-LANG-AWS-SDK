To check your AWS CLI configuration and change the region, follow these steps:
✅ Check AWS CLI Configuration

Use this command to view your current config:

aws configure list

This will show:

    Access Key

    Secret Key (masked)

    Region

    Output format

🌍 Check Current Region

aws configure get region

✏️ Change Region

To change the default region:

aws configure

It will prompt you again:

AWS Access Key ID [****]:
AWS Secret Access Key [****]:
Default region name [us-east-1]: <type new region here>
Default output format [json]:

Alternatively, change just the region with:

aws configure set region us-west-2

🔁 Replace us-west-2 with your desired AWS region code.

Would you like a list of all valid AWS region codes?
You said:
just path
ChatGPT said:

The AWS CLI configuration file paths are:

    Config file (region, output format, etc.):

~/.aws/config

Credentials file (Access & Secret keys):

    ~/.aws/credentials

On Windows, they are usually located at:

C:\Users\<YourUsername>\.aws\config  
C:\Users\<YourUsername>\.aws\credentials