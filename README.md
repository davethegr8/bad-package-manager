# bad package manger

it's a package manager that just checks out what you tell it to. it's less awful (maybe) than git submodules and much more terrible than actual package managers.

> you probably shouldn't use this  
> \- _me_

but, if you really want to, here's how

### installing

I guess do this?

```
$ go get -u github.com/davethegr8/bad-package-manager/cmd/bpm
```

you'll then have a `bpm` executable in your `$GOPATH/bin`

### using

make a file called `dependencies.json` that looks like this:

```
{
  "require": {
    "{dest}": "{repo}"
  }
}
```

`{dest}` can be any relative folder. `{repo}` is a git repo, and you can specify something that kinda looks like a commit with `#{commitish}`. example: `https://github.com/davethegr8/bad-package-manager#master`

then, run `bpm` and it will check out into your `{dest}` folder
