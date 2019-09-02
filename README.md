prettyprogress
====

a tiny library for printing nice-looking console progress

[![GoDoc](https://godoc.org/github.com/julz/prettyprogress?status.png)](https://godoc.org/github.com/julz/prettyprogress)

# Quick Start

See [example/simple/main.go](example/simple/main.go) for the UI-only structs or
[/example/multistep/main.go](example/multistep/main.go) for the fancy
thread-safe interface.

# Features:

 - Nice unicode icons
 - Multiple parallel steps
 - Unicode progress bars with smooth-scrolling via box-drawing characters (inspired by [this blog post](https://mike42.me/blog/2018-06-make-better-cli-progress-bars-with-unicode-block-characters))
 - Simple, composable API: just prints itself, use e.g. [uilive](https://github.com/gosuri/uilive) for animation
 - Basic [UI-only package](https://godoc.org/github.com/julz/prettyprogress/ui) if you just want to print a nice progress bar or grab a pretty unicode icon

# Example:

Output from [/example/multistep/main.go](example/multistep/main.go):

![example](doc/demo.gif)
