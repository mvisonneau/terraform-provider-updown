# DEPRECATED : In favor of `updown_recipient` (https://updown.io/api#GET-/api/webhooks)
#
resource "updown_webhook" "mywebhook" {
  url = "https://my-nice-webhook.com"
}
