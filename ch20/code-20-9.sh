# 코드 20.9 openssl 명령 설치

$ brew install openssl
$ echo 'export PATH="/opt/homebrew/opt/openssl@3/bin:$PATH"' >> ~/.zshrc
$ source ~/.zshrc
$ openssl version
OpenSSL 3.4.1 11 Feb 2025 (Library: OpenSSL 3.4.1 11 Feb 2025)