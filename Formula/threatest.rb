# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Threatest < Formula
  desc ""
  homepage "https://github.com/DataDog/threatest"
  version "1.1.1"
  license "Apache-2.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/DataDog/threatest/releases/download/v1.1.1/threatest_1.1.1_Darwin_arm64.tar.gz"
      sha256 "892fe2b378849e7c602c6b6acf87db78b85c9df7a21c03830193358f71b2d84b"

      def install
        bin.install "threatest"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/threatest/releases/download/v1.1.1/threatest_1.1.1_Darwin_x86_64.tar.gz"
      sha256 "def1517c1e33fb7753c7cd527a29ab15a4521b78c449b236d4f6ef1eb7b62e92"

      def install
        bin.install "threatest"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/DataDog/threatest/releases/download/v1.1.1/threatest_1.1.1_Linux_arm64.tar.gz"
      sha256 "dc71e3255e8aae3ca1cd897ed14fc55e4d858781e1a837a196bd68b29fbb4e92"

      def install
        bin.install "threatest"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/threatest/releases/download/v1.1.1/threatest_1.1.1_Linux_x86_64.tar.gz"
      sha256 "08fd698f0441d9e834d23693d5b7cb050e2ea6c4c355e6657aac8eb13f85dca2"

      def install
        bin.install "threatest"
      end
    end
  end
end
