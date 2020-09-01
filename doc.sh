alias doc="docker"
alias pose="docker-compose"

up() {
  set -x
  pose up --force-recreate "$@"
  set +x
}

b() {
  set -x
  docker build \
         --build-arg USER_ID=$(id -u) \
         --build-arg GROUP_ID=$(id -g) \
         -t mathapp .
  set +x
}

down() {
  set -x
  pose down "$@"
  set +x
}

build() {
  set -x
  pose build "$@"
  set +x
}

logs() {
  set -x
  pose logs -f "$@"
  set +x
}

vol() {
  set -x
  doc volume inspect $1
  set +x
}
