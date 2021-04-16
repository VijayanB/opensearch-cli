# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class OpensearchCli < Formula
  desc "Software to create fast and easy drum rolls."
  homepage "https://github.com/VijayanB/opensearch-cli"
  version "opensearch-cli_2.0.2"
  license "Apache V2"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://d3g5vo6xdbdb9a.cloudfront.net/downloads/elasticsearch-clients/opendistro-cli/opendistro-odfe-cli-1.1.0-macos-x64.zip"
    sha256 "bb480d8841ba79db102ca14cdd5d7db52bd0bd892aeee5c3095b79652497b02b"
  end
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/VijayanB/es-cli/releases/download/v2.0.2/opensearch-cli_Darwin_arm64.tar.gz"
    sha256 "c4dba62b39dca57e28c99282061040033a5ec8e59aa04e8fba4c0813cc5581ff"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://d3g5vo6xdbdb9a.cloudfront.net/downloads/elasticsearch-clients/opendistro-cli/opendistro-odfe-cli-1.1.0-linux-x64.zip"
    sha256 "be89570338c657219ae829d31b2a531c145d8e483303208f93e1c43c42289cb9"
  end
  if OS.linux? && Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
    url "https://d3g5vo6xdbdb9a.cloudfront.net/downloads/elasticsearch-clients/opendistro-cli/opendistro-odfe-cli-1.1.0-linux-arm64.zip"
    sha256 "82a8ea1241be1e9d3515183aa1af3dc8b15ef77daaedf3e88c6989a2a7c7156a"
  end

  def install
    bin.install "odfe-cli"
  end

  test do
    system "#{bin}/odfe-cli --version"
  end
end
