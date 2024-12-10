#!/bin/sh

set -u
set -o pipefail
set -e

echo "[pre-destroy] preparing to execute pre-destroy script"
SESSION_NAME="pre-destroy-"$NUON_INSTALL_ID

echo "[pre-destroy] assuming deprovision role"
aws sts assume-role --duration-seconds 1800 --role-arn $DEPROVISION_ROLE --role-session-name $SESSION_NAME > /tmp/SESSION_NAME
eval $(cat /tmp/$SESSION_NAME | jq -r '.Credentials | "export AWS_ACCESS_KEY_ID=\(.AccessKeyId)\nexport AWS_SECRET_ACCESS_KEY=\(.SecretAccessKey)\nexport AWS_SESSION_TOKEN=\(.SessionToken)\n"')
aws sts get-caller-identity

echo "[pre-destroy] getting cluster credentials with aws eks update-kubeconfig"
aws eks update-kubeconfig --name $NUON_INSTALL_ID --alias $NUON_INSTALL_ID

echo "[pre-destroy] deleting install namespace"
kubectl delete namespace $NUON_INSTALL_ID --ignore-not-found=true

echo "[pre-destroy] cleaning up session"
rm -f /tmp/SESSION_NAME
