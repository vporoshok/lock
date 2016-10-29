# Lock with battery

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/vporoshok/lock/master/LICENSE)
[![GoDoc](http://godoc.org/github.com/vporoshok/lock?status.png)](http://godoc.org/github.com/vporoshok/lock)
[![Build Status](https://travis-ci.org/vporoshok/lock.svg?branch=master)](https://travis-ci.org/vporoshok/lock)
[![Coverage Status](https://coveralls.io/repos/github/vporoshok/lock/badge.svg?branch=master)](https://coveralls.io/github/vporoshok/lock?branch=master)

> Yet another TryLock implementation

In one of my work projects I had the need to know the state of the mutex. To my surprise, in a standard implementation of the mutex was no way to do it. This library is built on top of `sync.Mutex` and provide more complex interface `Lock`.
  
**Warning!** This library slower standard realization approx 2 times
 
**Thank you!** The idea of the implementation was stolen at [https://github.com/joshlf/sync](https://github.com/joshlf/sync)

## Requirements

Any supported version of Go.

## Recent Changes

Implement `Lock` with mutex protected bool. Locking use infinite loop with checks.

## Installation

Use the `go` tool:
```sh
$ go get github.com/vporoshok/lock
```

## Support

Please [open an issue](https://github.com/vporoshok/lock/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a fork, add commits, and [open a pull request](https://github.com/vporoshok/lock/compare/).

## Copyright

Copyright (C) 2016 vporoshok@github.com
See [LICENSE](https://github.com/vporoshok/lock/tree/master/LICENSE)
file for details.
