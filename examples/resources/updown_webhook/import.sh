# DEPRECATED : In favor of `updown_recipient` (https://updown.io/api#GET-/api/webhooks)
#
# You can find the corresponding IDs by looking at the following API endpoint:
# curl -s https://updown.io/api/webhooks\?api-key\=<your_api_key>
# [{"id":"123456789abcdef","url":"https://example.com"}]

terraform import updown_webhook.my_webhook 123456789abcdef
