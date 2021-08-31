# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [0ver](https://0ver.org) (more or less).

## [Unreleased]

### Added

- New `updown_webhook` resource to manage webhooks configuration

### Changed

- Switched libs to github.com/hashicorp/terraform-plugin-sdk instead of formerly github.com/hashicorp/terraform
- Upgraded to go 1.17
- Bumped dependencies to their latest versions

## [0.2.1] - 2020-09-11

### Added

- gosec tests
- binary releases for windows, darwin & linux - (386, amd64 & arm64)

### Changed

- Upgraded to go 1.15
- Upgraded to terraform 0.13.2

## [0.2.0] - 2019-05-27

### Added

- new `data` resource : **updown_nodes** returns the list of testing nodes ipv4 and ipv6 addresses
- new `resource` : **updown_check** creates a check
- Makefile
- CI
- License

[Unreleased]: https://github.com/mvisonneau/terraform-provider-updown/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/mvisonneau/terraform-provider-updown/tree/v0.2.1
[0.2.0]: https://github.com/mvisonneau/terraform-provider-updown/tree/0.2.0
