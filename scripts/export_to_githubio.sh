
#!/bin/sh

scripts=$(dirname $0)

githubio=../spudtrooper.github.io
mkdir -p ${githubio}/opentable
cp -R html/* ${githubio}/opentable

pushd ${githubio} && \
  ./scripts/commit.sh && \
  popd