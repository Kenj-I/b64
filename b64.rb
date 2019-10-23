# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://www.rubydoc.info/github/Homebrew/brew/master/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class b64 < Formula
  desc "Encoding base64"
  url "https://github.com/Kenj-I/b64/releases/download/v1.0.2/b64_1.0.2_macOS_64bit.tar.gz"
  sha256 "b77865067680ca3c7007a69f9bfe1f7bd1ae4d5a767e61a6d0c26fc9221123b2"
  # depends_on "cmake" => :build

  def install
    bin.install "b64"
  end
end