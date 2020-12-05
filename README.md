# haumea
CLI for easily switching servers in AWS target groups.

# Installation
This package can be installed with the following command:
```
go get github.com/shuheitakada/haumea
```

# Usage
## Configuring the AWS CLI
This package uses the AWS SDK, so you need to configure the AWS CLI. See [Quick configuration with aws configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html).

## config.yml
Create `~/.haumea.yml`. An example is below.
```
example_1:
  web:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/web-targets/73e2d6bc24d8a067
    targets: [i-80c8dd94, i-71b7ce85]
  admin:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/admin-targets/3bb63f11dfb0faf9
    targets: [i-abo28s30]
  api:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/api-targets/19xk27cb2ka81hs8
    targets: [i-12lz03n2, i-9s020xj3, i-g10x9cn2]
example_2:
  web:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/web-targets/73e2d6bc24d8a067
    targets: [i-ceddcd4d, i-10xl28vk]
  admin:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/admin-targets/3bb63f11dfb0faf9
    targets: [i-pkeb180c]
  api:
    target_group_arn: arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-new-targets/19xk27cb2ka81hs8
    targets: [i-3lzu00z8, i-qjz020z1, i-u92lz02n]
```

## register
To register targets in the target group, run the following command.
```
haumea register example_1 web
haumea register example_1 admin
haumea register example_1 api
```

The above example is the same as the command below.
```
haumea register example_1
```

## deregister
To deregister targets in the target group, run the following command.
```
haumea deregister example_2 web
haumea deregister example_2 admin
haumea deregister example_2 api
```

The above example is the same as the command below.
```
haumea deregister example_2
```

## healthcheck
You can check health of the specified targets
```
haumea healthcheck example_1 web
```

or all of your targets.
```
haumea healthcheck example_1
```
# License
haumea is released under the Apache 2.0 license. See [LICENSE](https://github.com/shuheitakada/haumea/blob/main/LICENSE)
