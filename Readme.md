## [Security](http://security.stackexchange.com/)

  - Golang
    - [Go.crypto Package](http://godoc.org/code.google.com/p/go.crypto)
    - [Practical Cryptography with Go - book](http://book.gokyle.org/)
    - [Cryptobox](http://cryptobox.tyrfingr.is/)
    - [Talks from GoKyle](http://talks.gokyle.org/denver.gophers/2013/)
      - [Symmetric Cryptography With Go](http://denvergophers.com/2013-03/symmetric.slide)
      - [Public Key Cryptography, 2013.06](http://talks.gokyle.org/denver.gophers/2013/pkc.slide)
      - [Cryptobox: A Crypto API 2013 August 22](http://talks.gokyle.org/denver.gophers/2013/cryptobox.article)
    - [Lessons learned and misconceptions regarding encryption and cryptology](http://security.stackexchange.com/questions/2202/lessons-learned-and-misconceptions-regarding-encryption-and-cryptology)
      - Your best case is to use a high-level well-vetted scheme: for communication security, use TLS (or SSL); for data at rest, use GPG (or PGP). If you can't do that, use a high-level crypto library, like cryptlib, GPGME, Keyczar, or NaCL, instead of a low-level one, like OpenSSL, CryptoAPI, JCE, etc.. Thanks to Nate Lawson for this suggestion.

<!-- PROJECTS_LIST_START -->
    *** GENERATED BY https://github.com/mindreframer/techwatcher (ruby _sh/pull security-stuff) *** 

    ClarityServices/symmetric-encryption:
      Symmetric Encryption for Ruby Projects using OpenSSL
       53 commits, last change: 2013-09-25 08:49:39, 65 stars, 14 forks

    gokyle/cryptobox:
      NaCL-like toolkit using FIPS ciphers.
       38 commits, last change: 2013-08-26 22:17:55, 27 stars, 1 forks

    gokyle/gokyle.talks:
      Nothing to see here. Move along.
       75 commits, last change: 2013-08-29 14:06:17, 0 stars, 0 forks

    jedisct1/libsodium:
      P(ortable|ackageable) NaCl-based crypto library
       537 commits, last change: 2013-10-07 19:58:06, 504 stars, 37 forks
<!-- PROJECTS_LIST_END -->
