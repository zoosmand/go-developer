# Golang modules

~~~ bash
go mod init lesson_10/zoosman
~~~

~~~ bash
cat << EOF >> go.mod
require github.com/zoosmand/usecons/v3 v3.0.4
EOF
~~~

~~~ bash
go mod tidy
~~~
