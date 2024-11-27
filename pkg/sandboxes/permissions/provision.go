package permissions

// permissions for the provision role common to all aws sandboxes
var BaseProvisionPermissions = []string{
	"ec2:AllocateAddress",
	"ec2:AssociateRouteTable",
	"ec2:AttachInternetGateway",
	"ec2:AuthorizeSecurityGroupEgress",
	"ec2:AuthorizeSecurityGroupIngress",
	"ec2:CreateInternetGateway",
	"ec2:CreateLaunchTemplate",
	"ec2:CreateLaunchTemplateVersion",
	"ec2:CreateNatGateway",
	"ec2:CreateRoute",
	"ec2:CreateRouteTable",
	"ec2:CreateSecurityGroup",
	"ec2:CreateSubnet",
	"ec2:CreateTags",
	"ec2:CreateVpc",
	"ec2:DescribeAddresses",
	"ec2:DescribeAvailabilityZones",
	"ec2:DescribeInternetGateways",
	"ec2:DescribeLaunchTemplateVersions",
	"ec2:DescribeLaunchTemplates",
	"ec2:DescribeNatGateways",
	"ec2:DescribeNetworkAcls",
	"ec2:DescribeRouteTables",
	"ec2:DescribeSecurityGroupReferences",
	"ec2:DescribeSecurityGroupRules",
	"ec2:DescribeSecurityGroups",
	"ec2:DescribeSubnets",
	"ec2:DescribeTags",
	"ec2:DescribeVpcAttribute",
	"ec2:DescribeVpcClassicLink",
	"ec2:DescribeVpcClassicLinkDnsSupport",
	"ec2:DescribeVpcs",
	"ec2:ModifyLaunchTemplate",
	"ec2:ModifySubnetAttribute",
	"ec2:ModifyVpcAttribute",
	"ec2:RevokeSecurityGroupEgress",
	"ec2:RunInstances",
	"ecr:CreateRepository",
	"ecr:DescribeRepositories",
	"ecr:ListTagsForResource",
	"ecr:TagResource",
	"iam:AttachRolePolicy",
	"iam:CreateOpenIDConnectProvider",
	"iam:CreatePolicy",
	"iam:CreatePolicyVersion",
	"iam:CreateRole",
	"iam:CreateServiceLinkedRole",
	"iam:GetOpenIDConnectProvider",
	"iam:GetPolicy",
	"iam:GetPolicyVersion",
	"iam:GetRole",
	"iam:GetRolePolicy",
	"iam:ListAttachedRolePolicies",
	"iam:ListRolePolicies",
	"iam:PassRole",
	"iam:PutRolePolicy",
	"iam:TagOpenIDConnectProvider",
	"iam:TagPolicy",
	"iam:TagRole",
	"iam:UpdateAssumeRolePolicy",
	"kms:CreateAlias",
	"kms:CreateGrant",
	"kms:CreateKey",
	"kms:DescribeKey",
	"kms:GetKeyPolicy",
	"kms:GetKeyRotationStatus",
	"kms:ListAliases",
	"kms:ListResourceTags",
	"kms:PutKeyPolicy",
	"kms:TagResource",
	"logs:CreateLogGroup",
	"logs:DescribeLogGroups",
	"logs:ListTagsLogGroup",
	"logs:PutRetentionPolicy",
	"logs:TagLogGroup",
	"logs:TagResource",
	"route53:ChangeResourceRecordSets",
	"route53:ChangeTagsForResource",
	"route53:CreateHostedZone",
	"route53:GetChange",
	"route53:GetHostedZone",
	"route53:ListResourceRecordSets",
	"route53:ListTagsForResource",
	"s3:GetObject",
	"s3:ListBucket",
	"s3:PutObject",
}