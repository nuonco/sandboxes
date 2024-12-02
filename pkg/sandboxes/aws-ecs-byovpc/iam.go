package awsecsbyovpc

import perms "github.com/nuonco/sandboxes/pkg/sandboxes/permissions"

// provision role permissions specific to this sandbox
var ProvisionPermissions = append([]string{
	"application-autoscaling:*",
	"ecs:CreateCapacityProvider",
	"ecs:DescribeCapacityProviders",
	"ecs:CreateCluster",
	"ecs:PutClusterCapacityProviders",
	"ecs:DescribeClusters",
	"ecs:TagResource",
	"ecs:ListTagsForResource",
	"ecs:RegisterTaskDefinition",
}, perms.BaseProvisionPermissions...)

// Full provision rol policy for this sandbox
var ProvisionPolicy = perms.Policy{
	Version: "2012-10-17",
	Statement: []perms.Statement{
		{
			Effect:   "Allow",
			Resource: "*",
			Action:   ProvisionPermissions,
		},
	},
}

// deprovision role permissions specific to this sandbox
var DeprovisionPermissions = append([]string{
	"ecs:DeleteCapacityProvider",
	"ecs:DeleteCluster",
	"logs:ListTagsForResource",
}, perms.BaseDeprovisionPermissions...)

// Full deprovision role policy for this sandbox
var DeprovisionPolicy = perms.Policy{
	Version: "2012-10-17",
	Statement: []perms.Statement{
		{
			Effect:   "Allow",
			Resource: "*",
			Action:   DeprovisionPermissions,
		},
	},
}
