## ev
explore the evolution of a function in your browser.
```
usage: ev <funcname>:<file>
```
Will open the browser showing a friendly UI of the git-log history of the function `funcname` from `file`. Below is an example screenshot viewing the `IndexAny` function from the Go standard library `bytes` package.

![ev](http://i67.tinypic.com/2eatsfc.png)

See a demo of it in action on `NewBufferString` from the `bytes` package [on YouTube](https://youtu.be/GqfDZX7xLUQ), or try it out yourself!

### Installation

```
go get gbbr.io/ev
go install gbbr.io/ev/cmd/...
```
---

Note that `ev` uses `git log -L:<re>:<fn>` syntax to search so it may not always match exactly.
