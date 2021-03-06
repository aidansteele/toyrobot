# Toy Robot

[![GoDoc](https://godoc.org/github.com/aidansteele/toyrobot/pkg/toyrobot?status.svg)][godoc]
[![Build Status](https://travis-ci.org/aidansteele/toyrobot.svg?branch=master)][travis]

* The interesting code is in [`pkg/toyrobot`](pkg/toyrobot).

* Documentation is autogenerated from comments and is available at [GoDoc.org][godoc].

* Files ending in `_test.go` have special treatment. They are excluded from the 
  compiled output and are executed by running `go test`.
  
* Functions in `_test.go` files named `Example*()` are _also_ tests. Everything 
  after the `// Output:` line is compared to the output written stdout in that
  function. Non-matching is considered a failure. 

* Provided example inputs + outputs and an "end-to-end" example is available 
  in [`pkg/toyrobot/simulation_test.go`](pkg/toyrobot/simulation_test.go).

* Continuous integration has been set up at [Travis CI][travis].

[godoc]: https://godoc.org/github.com/aidansteele/toyrobot/pkg/toyrobot
[travis]: https://travis-ci.org/aidansteele/toyrobot
