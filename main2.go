package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var instanceIDs = []string{
	"i-01f2ea7b820b96c58",
	"i-0e895bfc3270142d7",
	"i-0eb21ace5c041a907",
	"i-0cd358de55c5f2001",
}

const newInstanceType = "t3.medium"
const originalType = "t2.micro"
const tagKey = "instance_mode"
const tagValue = "other_instances"

func main() {

	// Create a context for the AWS SDK operations context.Background() is a simple "empty" context, it doesn't affect the rest of your code, and everything else would remain unchanged. This makes your code future-proof and much easier to manage when you need to add timeouts or cancellation, as it avoids any major restructuring
	ctx := context.Background()

	// // Create a context with a 2-minute timeout. in the future we can just comment out ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// defer cancel() // Ensure we call cancel to avoid context leak

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("🚫 Failed to load AWS config:", err)
	}

	// Create EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Describe Instances
	descOutput, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	})

	if err != nil {
		log.Fatalf("❌ Failed to describe instances: %v", err)
	}

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	// Loop through all instances and perform actions concurrently
	for _, reservation := range descOutput.Reservations {
		for _, instance := range reservation.Instances {
			id := *instance.InstanceId
			currentType := string(instance.InstanceType)
			state := string(instance.State.Name)

			var modeTag string
			for _, tag := range instance.Tags {
				if *tag.Key == tagKey {
					modeTag = *tag.Value
					break
				}
			}

			// Start goroutine to stop and modify instance based on conditions
			wg.Add(1) // Increment the counter for goroutines
			go func(instanceID, currentType, state, modeTag string) {
				defer wg.Done() // Decrement the counter when the goroutine completes

				// Revert instance if tagged as "other_instances" and running
				if modeTag == tagValue && state == "running" {
					fmt.Printf("🛑 Stopping Instance %s to 🔄 Revert Back to Type %s...\n", instanceID, originalType)
					stopAndModify(ctx, ec2Client, instanceID, originalType, true, false)
				}

				// If instance is already t2.micro, upgrade to new type
				if currentType == originalType {
					stopAndModify(ctx, ec2Client, instanceID, newInstanceType, false, true)
					fmt.Printf("⬆️ Upgrading instance %s to %s and ▶️ starting...\n", instanceID, newInstanceType)
				}
			}(id, currentType, state, modeTag)
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

// stopAndModify stops the instance, modifies its type, optionally removes a tag, and optionally restarts it.
func stopAndModify(ctx context.Context, client *ec2.Client, instanceID string, newType string, removeTag bool, restartAfter bool) {
	// Stop the instance
	stopOutput, err := client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		log.Printf("❌ Failed to stop instance %s: %v", instanceID, err)
		return
	}
	log.Printf("🛑 Stopping instance %s, current state: %v", instanceID, stopOutput.StoppingInstances)

	// Wait until the instance is fully stopped
	waitUntilStopped(ctx, client, instanceID)

	// Modify the instance type (ignore output because it's always empty unless there's an error)
	_, err = client.ModifyInstanceAttribute(ctx, &ec2.ModifyInstanceAttributeInput{
		InstanceId:   aws.String(instanceID),
		InstanceType: &ec2types.AttributeValue{Value: aws.String(newType)},
	})
	if err != nil {
		log.Printf("❌ Failed to modify instance type for %s: %v", instanceID, err)
		return
	}
	log.Printf("🔧 Modified instance %s to type %s", instanceID, newType)

	// Tag management
	if removeTag {
		_, err := client.DeleteTags(ctx, &ec2.DeleteTagsInput{
			Resources: []string{instanceID},
			Tags:      []ec2types.Tag{{Key: aws.String(tagKey)}},
		})
		if err != nil {
			log.Printf("❌ Failed to delete tag on %s: %v", instanceID, err)
		} else {
			log.Printf("🏷️  Removed tag '%s' from %s", tagKey, instanceID)
		}
	} else {
		_, err := client.CreateTags(ctx, &ec2.CreateTagsInput{
			Resources: []string{instanceID},
			Tags:      []ec2types.Tag{{Key: aws.String(tagKey), Value: aws.String(tagValue)}},
		})
		if err != nil {
			log.Printf("❌ Failed to add tag on %s: %v", instanceID, err)
		} else {
			log.Printf("🏷️  Added tag '%s=%s' to %s", tagKey, tagValue, instanceID)
		}
	}

	// Restart instance if needed
	if restartAfter {
		startOutput, err := client.StartInstances(ctx, &ec2.StartInstancesInput{
			InstanceIds: []string{instanceID},
		})
		if err != nil {
			log.Printf("❌ Failed to start instance %s: %v", instanceID, err)
		} else {
			log.Printf("▶️  Started instance %s, current state: %v", instanceID, startOutput.StartingInstances)
		}
	}
}

// waitUntilStopped waits until the specified instance is stopped.
func waitUntilStopped(ctx context.Context, client *ec2.Client, instanceID string) {
	fmt.Printf("🕒 Waiting for instance %s to 🛑 stop...\n", instanceID)
	for {
		out, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			InstanceIds: []string{instanceID},
		})
		if err != nil {
			log.Printf("Error while waiting: %v", err)
			break
		}

		state := out.Reservations[0].Instances[0].State.Name
		if state == "stopped" {
			break
		}
		time.Sleep(5 * time.Second)
	}
}
