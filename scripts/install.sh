#!/usr/bin/env sh
# This script is used in the README and https://www.diginfra.khulnasoft.com/docs/#quick-start
set -e

# check_sha is separated into a defined function so that we can
# capture the exit code effectively with `set -e` enabled
check_sha() {
  (
    cd /tmp/
    shasum -sc "$1"
  )

  return $?
}

os=$(uname | tr '[:upper:]' '[:lower:]')
arch=$(uname -m | tr '[:upper:]' '[:lower:]' | sed -e s/x86_64/amd64/)
if [ "$arch" = "aarch64" ]; then
  arch="arm64"
fi

version=${DIGINFRA_VERSION:-latest}
url="https://diginfra.khulnasoft.com/downloads/${version}"
tar="diginfra-$os-$arch.tar.gz"
echo "Downloading version ${version} of diginfra-$os-$arch..."
curl -sL "$url/$tar" -o "/tmp/$tar"
echo

code=$(curl -s -L -o /dev/null -w "%{http_code}" "$url/$tar.sha256")
if [ "$code" = "404" ]; then
    echo "Skipping checksum validation as the sha for the release could not be found, no action needed."
else
  if [ -x "$(command -v shasum)" ]; then
    echo "Validating checksum for diginfra-$os-$arch..."
    curl -sL "$url/$tar.sha256" -o "/tmp/$tar.sha256"

    if ! check_sha "$tar.sha256"; then
      echo
      read -r -p "Installation checksum failed. This could be a security issue. Would you like to continue? (y/n) " answer
      if [ "$answer" != "y" ]; then
        echo
        echo "Exiting, please email hello@diginfra.khulnasoft.com for help."
        exit 1
      fi
    fi

    rm "/tmp/$tar.sha256"
  else
    echo "Skipping checksum validation as the shasum command could not be found, no action needed."
  fi
fi
echo

tar xzf "/tmp/$tar" -C /tmp
rm "/tmp/$tar"

echo "Moving /tmp/diginfra-$os-$arch to /usr/local/bin/diginfra (you might be asked for your password due to sudo)"
if [ -x "$(command -v sudo)" ]; then
  sudo mv "/tmp/diginfra-$os-$arch" "/usr/local/bin/diginfra"
else
  mv "/tmp/diginfra-$os-$arch" "/usr/local/bin/diginfra"
fi
echo
echo "Completed installing $(diginfra --version)"
