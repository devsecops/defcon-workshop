# Imports the Google Cloud client library
from google.cloud import pubsub

# Instantiates a client
pubsub_client = pubsub.Client()

# The name for the topic and subscription
sub_name = 'my-new-sub'
topic_name = 'my-new-topic'

# Prepares the new topic and subscription
topic = pubsub_client.topic(topic_name)
sub = topic.subscription(sub_name)

# Verifying the subscription exists
assert sub.exists()

# Listening from the subscription
while True:
    results = sub.pull(return_immediately=True)

    if len(results) == 0:
        print "Waiting for a message in the topic"
        time.sleep(3)
    else:
        print "Length of result: " + str(len(results)) + "\n"
        for ack_id, message in results:
            sub.acknowledge([ack_id]) # Ack'ing the message so that it gets removed from the topic
            print "Received Message Data: " + message.data + "\n"
        break