*This is a demonstration repo for https://github.com/golang/go/issues/38196.*

This repository simulates a private key generator CLI tool.

---

The complicated way of using this tool is to manually clone the git repository first, then install the tool and then run it:

```
$ cd /tmp

$ git clone https://github.com/karalabe/go-issue-38196-keygen
Cloning into 'go-issue-38196-keygen'...
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (5/5), done.
remote: Total 5 (delta 0), reused 5 (delta 0), pack-reused 0
Unpacking objects: 100% (5/5), done.

$ go install
go: downloading github.com/karalabe/go-issue-38196-crypto v1.0.0

$ go-issue-38196-keygen
Generating crypto key: secure key
```

As expected, it prints out a "secure" key because it depends on `v1.0.0` of our demo crpyto package, which is secure: https://github.com/karalabe/go-issue-38196-crypto/blob/bdfb384b3808ff99bee76c437b2ec584337396df/crypto.go#L4

---

Now lets do it the simple way, simply go getting the command:

```
$ cd /tmp
$ go get github.com/karalabe/go-issue-38196-keygen

$ go-issue-38196-keygen
Generating crypto key: insecure key
```

Well, shit.

Go ignores the `go.mod` file embedded in this repository (which pins the crypto package to `v1.0.0`) and pulls in the latest development version (which is "deliberately" broken).

---

The problem is that unless there's an API breakage, the user who used `go get` to pull a Go binary will have **absolutely no warning** that their code was bulit against different dependencies. Worse, they will live in the false belief that they are safe because the dependencies are pinned.

This can lead to buggy behavior at best (if a dependency is simply broken at `master`), all the way to a security vulnerability at worse (if crypto is broken at `master`).

Furthermore, if a malicious entity gains push access to a crypto package - by breaking the security on master - they can force this behavior on existing dependents of that package, all whilst Go modules touts that it checksums and pins and whatnots everything.