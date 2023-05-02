FROM debian:latest

ARG USER_UID=1000 \
  USER_GID=1000 \
  USER_NAME=vscode

ARG KUBECTL_VERSION=1.25.4 \
  KUBECTL_SHASUM_arm64=a8e9cd3c6ca80b67091fc41bc7fe8e9f246835925c835823a08a20ed9bcea1ba \
  KUBECTL_SHASUM_amd64=e4e569249798a09f37e31b8b33571970fcfbdecdd99b1b81108adc93ca74b522

ARG GO_VERSION=1.20.3 \
  GO_SHASUM_arm64=eb186529f13f901e7a2c4438a05c2cd90d74706aaa0a888469b2a4a617b6ee54 \
  GO_SHASUM_amd64=979694c2c25c735755bf26f4f45e19e64e4811d661dd07b8c010f7a8e18adfca

ARG OPERATOR_SDK_VERSION=v1.28.0 \
  OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download \
  OPERATOR_SDK_SHASUM_arm64=a99eefb61f4d66266ce337f80961f4bdb4b6c8bd5cf0c2d01f44081f1359846e \
  OPERATOR_SDK_SHASUM_amd64=da5b57c59f5c9751c0466569968d35e1ad118308f4f0a9fd38f800130cf85e8b

ADD docker.gpg /tmp/docker.gpg

RUN export DEBIAN_FRONTEND=noninteractive \
  ARCH=$(dpkg --print-architecture) \
  && apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    gnupg \
    make \
    openssh-client \
  && groupadd --gid ${USER_GID} ${USER_NAME} \
  && useradd -s /bin/bash --uid ${USER_UID} --gid ${USER_GID} -m ${USER_NAME} \
  # kubectl
  && curl -sSL -o /tmp/kubectl https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/${ARCH}/kubectl \
  && bash -c "echo \${KUBECTL_SHASUM_${ARCH}}  /tmp/kubectl" | sha256sum -c - \
  && chmod +x /tmp/kubectl \
  && mv /tmp/kubectl /usr/local/bin/kubectl \
  # Go
  && curl -sSL -o /tmp/go.tar.gz https://golang.org/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz \
  && bash -c "echo \${GO_SHASUM_${ARCH}}  /tmp/go.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf /tmp/go.tar.gz \
  && rm /tmp/go.tar.gz \
  # Docker
  && mkdir -p /etc/apt/keyrings/ \
  && gpg --dearmor -o /etc/apt/keyrings/docker.gpg < /tmp/docker.gpg \
  && chmod a+r /etc/apt/keyrings/docker.gpg \
  && echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" > /etc/apt/sources.list.d/docker.list \
  && apt-get update \
  && apt-get install -y --no-install-recommends \
    docker-ce \
    docker-ce-cli \
    docker-buildx-plugin \
    docker-compose-plugin \
  && usermod -a -G docker ${USER_NAME} \
  # Operator SDK
  && curl -fsSL -o /tmp/operator-sdk ${OPERATOR_SDK_DL_URL}/${OPERATOR_SDK_VERSION}/operator-sdk_linux_${ARCH} \
  && bash -c "echo \${OPERATOR_SDK_SHASUM_${ARCH}}  /tmp/operator-sdk" | sha256sum -c - \
  && chmod +x /tmp/operator-sdk \
  && mv /tmp/operator-sdk /usr/local/bin/operator-sdk \
  # Cleanup
  && rm -rf /var/lib/apt/lists/*