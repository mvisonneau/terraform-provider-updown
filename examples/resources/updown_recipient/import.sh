# You can find the corresponding IDs by looking at the following API endpoint:
# curl -s https://updown.io/api/recipients\?api-key\=<your_api_key>
# [{"id":"email:123456789","type":"email","name":"foo@bar.baz","immutable":false}]

terraform import updown_recipient.my_recipient email:123456789
