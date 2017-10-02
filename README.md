## ev
explore the evolution of a function in your browser.
```
usage: ev <funcname>:<file>
```
The command will open the browser showing snapshots of how the function `funcname` from `file` evolved in time throughout various git commits. I created it to better help me understand a codebase while trying to learn more about the implementation of Go's standard library. It can be used with any programming language.

Below is an example screenshot viewing the `IndexAny` function from Go's `bytes` package.

![ev](http://i67.tinypic.com/2eatsfc.png)

See a [demo](https://youtu.be/GqfDZX7xLUQ) of it, or try it out yourself!

### Installation

```
go get gbbr.io/ev/cmd/...
```
---

Note that `ev` uses `git log -L:<re>:<fn>` syntax, meaning that it also comes with its limitations. More specifically, if the file has multiple functions sharing the same name (ie. both method and function) it will only refer to the first occurrence starting from the top of the file.
