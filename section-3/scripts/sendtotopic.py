# Imports the Google Cloud client library
from google.cloud import pubsub

# Instantiates a client
pubsub_client = pubsub.Client()

# The name for the new topic
topic_name = 'my-new-topic'

# Prepares the new topic
topic = pubsub_client.topic(topic_name)

# Verifying the topic exists
assert topic.exists()

string = "This is a test message!"

# Publishing a string data
print "Data to be published is: " + string + "\n"
response = topic.publish(string)
print "Message ID: " + response + "\n"
