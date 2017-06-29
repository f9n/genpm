# genpm
General package manager, it covers pacman,apt-get...
Golang version must be over 1.8

# Installation

```sh
  $ go get github.com/pleycpl/genpm
  $ cd ~/go/src/github.com/pleycpl/genpm
  $ go install
  $ export PATH=$PATH:~/go/bin
  $ genpm version
```

# Usage

```sh
  $ genpm version
  $ genpm install python
  $ genpm remove python
  $ genpm search python
  $ genpm info python
  $ genpm upgrade
```
