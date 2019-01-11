#!/bin/bash

# Output command before executing
set -x

# Exit on error
set -e

# Source environment variables of the jenkins slave
# that might interest this worker.
function load_jenkins_vars() {
  if [ -e "jenkins-env" ]; then
    cat jenkins-env \
      | grep -E "(DEVSHIFT_TAG_LEN|JENKINS_URL|GIT_BRANCH|GIT_COMMIT|BUILD_NUMBER|ghprbSourceBranch|ghprbActualCommit|BUILD_URL|ghprbPullId)=" \
      | sed 's/^/export /g' \
      > ~/.jenkins-env
    source ~/.jenkins-env
  fi
}

function install_deps() {
  # We need to disable selinux for now, XXX
  /usr/sbin/setenforce 0 || :

  # Get all the deps in
  yum -y install \
    docker \
    make \
    curl

  service docker start

  echo 'CICO: Dependencies installed'
}

function prepare() {
  # Let's test
  make docker-start
  make docker-deps
  make docker-build
  echo 'CICO: Preparation complete'
}

function cico_setup() {
  load_jenkins_vars;
  install_deps;
  prepare;
}

function run_tests_with_coverage() {
  make docker-test-unit

  # Upload coverage to codecov.io
  # -t <upload_token> copy from https://codecov.io/gh/fabric8-services/fabric8-cluster-client/settings
  bash <(curl -s https://codecov.io/bash) -t 8c44edb8-9d93-4404-b1e3-90691befff68

  echo "CICO: ran tests and uploaded coverage"
}
