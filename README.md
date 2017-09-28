## ev
explore the evolution of a function in your browser.
```
usage: ev <funcname>:<file>
```
Will open the browser showing the git-log history of the function `funcname` from `file` inside the current working directory. Note that this uses `git log -L:<re>:<fn>` syntax to search so it may not always match exactly.

![ev](http://i66.tinypic.com/jtx9uv.png)

See a demo of it in action on `NewBufferString` from the `bytes` package [on YouTube](https://youtu.be/Xawz4zR2kjc).

### Installation

```
go get gbbr.io/ev
go install gbbr.io/ev/cmd/...
```
