---
subcategory: "Route 53 Recovery Control Config"
layout: "aws"
page_title: "AWS: aws_route53recoverycontrolconfig_control_panel"
description: |-
  Provides an AWS Route 53 Recovery Control Config Control Panel
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_route53recoverycontrolconfig_control_panel

Provides an AWS Route 53 Recovery Control Config Control Panel.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53RecoverycontrolconfigControlPanel } from "./.gen/providers/aws/route53-recoverycontrolconfig-control-panel";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Route53RecoverycontrolconfigControlPanel(this, "example", {
      clusterArn:
        "arn:aws:route53-recovery-control::123456789012:cluster/8d47920e-d789-437d-803a-2dcc4b204393",
      name: "balmorhea",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `clusterArn` - (Required) ARN of the cluster in which this control panel will reside.
* `name` - (Required) Name describing the control panel.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the control panel.
* `defaultControlPanel` - Whether a control panel is default.
* `routingControlCount` - Number routing controls in a control panel.
* `status` - Status of control panel: `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Route53 Recovery Control Config Control Panel using the control panel arn. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53RecoverycontrolconfigControlPanel } from "./.gen/providers/aws/route53-recoverycontrolconfig-control-panel";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Route53RecoverycontrolconfigControlPanel.generateConfigForImport(
      this,
      "mypanel",
      "arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8"
    );
  }
}

```

Using `terraform import`, import Route53 Recovery Control Config Control Panel using the control panel arn. For example:

```console
% terraform import aws_route53recoverycontrolconfig_control_panel.mypanel arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8
```

<!-- cache-key: cdktf-0.20.8 input-2684118596a4f1a5ee55c039284d0ef47e9b181e6c6212004319ac88f885b5d3 -->