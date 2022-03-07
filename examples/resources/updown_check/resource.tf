resource "updown_check" "mywebsite" {
  alias        = "https://example.com"
  apdex_t      = 1.0
  enabled      = true
  period       = 30
  published    = true
  url          = "https://test.example.com/healthz"
  string_match = "OK"
  mute_until   = "tomorrow"

  disabled_locations = [
    "mia",
  ]

  custom_headers = {
    "X-GREAT-HEADER" = "yay!"
  }

  recipients = [
    "email:123456789"
  ]
}
