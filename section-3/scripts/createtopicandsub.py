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

# Creates the new topic and sub on that topic
topic.create()
print('Topic {} created.'.format(topic.name))
assert topic.exists()

sub.create()
print('Subscription {} created.'.format(sub.name))
assert sub.exists()