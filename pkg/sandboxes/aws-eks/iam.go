package awseks

import perms "github.com/nuonco/sandboxes/pkg/sandboxes/permissions"

// provision role permissions specific to this sandbox
var ProvisionPermissions = append([]string{
	"ec2:DescribeAddressesAttribute",
	"ec2:CreateNetworkAclEntry",
	"ecr:UntagResource",
	"eks:ListAccessEntries",
	"eks:CreateAccessEntry",
	"eks:DescribeAccessEntry",
	"eks:UpdateAccessEntry",
	"eks:AssociateAccessPolicy",
	"eks:DisassociateAccessPolicy",
	"eks:CreateAddon",
	"eks:DescribeAddon",
	"eks:UpdateAddon",
	"eks:DescribeAddonConfiguration",
	"eks:DescribeAddonVersions",
	"eks:ListAddons",
	"eks:ListAssociatedAccessPolicies",
	"eks:CreateCluster",
	"eks:DescribeCluster",
	"eks:CreateNodegroup",
	"eks:DescribeNodegroup",
	"eks:UpdateNodegroupVersion",
	"eks:TagResource",
	"eks:UntagResource",
	"eks:ListTagsForResource",
	"eks:DescribeUpdate",
	"iam:UntagPolicy",
	"iam:UntagRole",
	"kms:UntagResource",
	"logs:UntagResource",
	"logs:ListTagsForResource",
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
	"ec2:DeleteNetworkAclEntry",
	"eks:DeleteAddon",
	"eks:DeleteCluster",
	"eks:DescribeCluster",
	"eks:DeleteNodegroup",
	"eks:DescribeNodegroup",
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