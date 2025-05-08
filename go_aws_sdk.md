To list all EC2 instance IDs using the AWS CLI, you can use the following command:

aws ec2 describe-instances --query "Reservations[*].Instances[*].InstanceId" --output text

🔍 Explanation:

    describe-instances: Retrieves information about your EC2 instances.

    --query: Filters the output to just the instance IDs.

    --output text: Makes the result cleaner (just IDs in plain text).

✅ Example output:
```sh
i-01f2ea7b820b96c58
i-0e895bfc3270142d7
i-0eb21ace5c041a907
i-0cd358de55c5f2001
```