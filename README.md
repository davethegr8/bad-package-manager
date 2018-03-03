# bad package manger

it's a package manager that just checks out what you tell it to

> you probably shouldn't use this  
>  - _me_

but, if you really want to, here's how

### installing

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
