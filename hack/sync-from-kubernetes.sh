#!/bin/bash

dir=$(mktemp -d "${TMPDIR:-/tmp/}$(basename 0).XXXXXXXXXXXX")
echo ${dir}

git remote add upstream-kube git@github.com:kubernetes/kubernetes.git
git fetch upstream-kube

previousKubeSHA=$(cat kubernetes-sha)

git log --format='%H' upstream-kube/master ${previousKubeSHA}..upstream-kube/master -- staging/src/k8s.io/apimachinery > ${dir}/commits

while read commitSha; do
	echo "working on ${commitSha}"
	git show ${commitSha} > ${dir}/diff-${commitSha}.diff

	git am --include=staging/src/k8s.io/apimachinery -p4
done <${dir}/commits