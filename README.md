# Trunk example
## Getting started
## Feature flags
### Open Feature
This project uses [OpenFeature](https://openfeature.dev) for feature flagging.

### Flipt
#### Pros
- Supports multiple storage backends
- Supports GRPC => Good performance
- No feature limit for open source version
#### Cons
- No change request for open source version
- No access control for open source version

### Flagsmith
#### Pros
- Support environment for feature (no need to copy feature between environments)
- Support comparing and searching for features by conditions
- Rich segmentation trait options
- Support object value for feature
#### Cons
- No audit logs for open source version
- No change request for open source version
- No access control for open source version
- Very slow boot up

### Growthbook
#### Pros
- Support environment for feature (no need to copy feature between environments)
- Support comparing and searching for features by conditions
- Rich segmentation trait options
- Support object value for feature
- Most comprehensive
- Have audit logs and webhooks
#### Cons
- No change request for open source version
- No access control for open source version
- Big image size
- Need mongo and redis
