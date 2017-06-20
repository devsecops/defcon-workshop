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

# Asserting the topic and sub exists
assert topic.exists()
assert sub.exists()

# Deleting the topic and sub
topic.delete()
print('Topic {} deleted.'.format(topic.name))

sub.delete()
print('Subscription {} deleted.'.format(sub.name))
