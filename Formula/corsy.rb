# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Corsy < Formula
  desc "Corsy is a proxy injector for development to avoid CORS issues while building SPA applications.."
  homepage "https://github.com/sivsivsree/corsy#corsy"
  version "0.0.3"
  license "MIT"
  bottle :unneeded

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/sivsivsree/corsy/releases/download/v0.0.3/corsy_0.0.3_darwin_amd64.tar.gz", :using => CurlDownloadStrategy.
      sha256 "fe70ad7c8cd50b2f6e62aeda7e5910a5b1391508966ab6f91143414e24d9fca1"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/sivsivsree/corsy/releases/download/v0.0.3/corsy_0.0.3_linux_amd64.tar.gz", :using => CurlDownloadStrategy.
      sha256 "af598d5e3ec5b878f194e8261ebe32a2ac9fdea0d909f5ad87a49772b3c256b6"
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/sivsivsree/corsy/releases/download/v0.0.3/corsy_0.0.3_linux_arm64.tar.gz", :using => CurlDownloadStrategy.
      sha256 "6106c530b67fb7eeb7a99a419f99b82cea1963dd91d6ac5cd36a81ccbf221eab"
    end
  end

  def install
    bin.install "corsy"
  end

  def caveats; <<~EOS
    CORS injector proxy for development
  EOS
  end
end