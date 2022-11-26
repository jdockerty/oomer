# oomer

Simple application to force `OOMKilled` containers in Kubernetes. The main reason I created
this is because `busybox` and other larger images do not container the shell built-in `exit` command,
this is easy to replicate with a tiny Go program.

This is used in conjunction with [`kubectl-oomd`](https://github.com/jdockerty/kubectl-oomd) for testing purposes.
