Native Go Zookeeper Client Library (Fork)
=========================================

[![GoDoc](https://godoc.org/github.com/multicharts/go-zookeeper/zk?status.svg)](https://godoc.org/github.com/multicharts/go-zookeeper/zk) [![Travis](https://travis-ci.org/multicharts/go-zookeeper.svg?branch=master)](https://travis-ci.org/multicharts/go-zookeeper)

[Zookeeper](https://zookeeper.apache.org/) client library for the Go language.

### Tools

The following tools are available:

* `zkcluster` - start ZooKeeper test cluster

    Usage:

        $ go get github.com/multicharts/go-zookeeper/cmd/zkcluster
        $ $GOPATH/bin/zkcluster -size 3

### Library

#### Tour

Detailed examples and additional code can be found in [tour.go](examples/tour.go). To run the tour:

    $ go run tour.go

#### Connecting

To establish a basic connection to zookeeper:

```go
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second*20)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
```

#### Znode

A znode is the fundamental entity for which operations are performed against. A znode in ZooKeeper can have data associated with it as well as children. It is like having a file-system that allows a file to also be a directory (ZooKeeper was designed to store coordination data: status information, configuration, location information, etc., so the data stored at each node is usually small, in the byte to kilobyte range.) The term znode and node are used interchangeably.

See the apache [zookeeper](https://zookeeper.apache.org/doc/trunk/zookeeperOver.html) project docs for more details.

#### Znode API

The following are fundamental operations on Znodes:
* Create
* Get
* Set
* Delete
* Exists
* Children
* Sync

#### Znode API example

```go
	var path string
	var err error

	if path, err = conn.Create("/myznode", []byte{}, zk.FlagPersistent, zk.WorldACL(zk.PermAll)); err != nil {
		panic(err)
	}

	if yes, _, err := conn.Exists(path); !yes || err != nil {
		panic(err)
	}

	if _, err = conn.Set(path, []byte("hello"), -1); err != nil {
		panic(err)
	}

	if data, _, err := conn.Get(path); err != nil || string(data) != "hello" {
		panic(err)
	} 

	if err = conn.Delete(path, -1); err != nil {
		panic(err)
	}
```

Acknowledgements
----------------

* This is a fork of [samuel/go-zookeeper](https://github.com/samuel/go-zookeeper) with a few bug fixes and improvements.
* Most bugfixes are ported from other forks.
* README and tour is adopted from [talbright/go-zookeeper](https://github.com/talbright/go-zookeeper).

License
-------

3-clause BSD. See LICENSE file.
