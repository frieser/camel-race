Camel Race
===============

Build
----------
```bash
$ GOPATH=`pwd`
$ export GOPATH 

$ go build src/github.com/frieser/camelrace/camelrace.go
```

Run
---------

```bash
$ GOPATH=`pwd`
$ export GOPATH 

$ go run src/github.com/frieser/camelrace/camelrace.go
```

Also is generated binaries for Linux x64 OS

```bash
./bin/camelrace --config=config.yml
```

Configuration
--------
You can pass a config file with the --config flag

There is an example of config file:

```yaml
distance: 30
players: 10
terrain: "SAND"
round_time: 4
round_divider: 2
movement_per_round: 3
```

Extensibility
------------
* You can add more kind of 'Actors' that implement the Actor interface
* You can add more kind of Players that has the Player underlying type
* You can add more kind of camels that implement the Movable interface
* You can add the attributes that you need

TODO
---------
* The camel attribute system is not finished, it is a preliminary idea with an event system
* The input timeout is not implement because I suppose is not the main target of the test. Instead I ask for the human movement desired
* The 'movement_per_round' config value is ignored
* The player's rolls are sequentially, maybe make it with threads improve the realistic of the game. I try to counteract this with a map that is not visited always in the same order
* The kind of camels is not configurable
* More test
* More doc
* Improve output
* Improve code cleanup

Testing
-----------

It is no all the code tested.

```bash
$ go test github.com/frieser/camelrace/{camel,player,race,event,round,lane}
ok  	github.com/frieser/camelrace/camel	0.006s
ok  	github.com/frieser/camelrace/player	0.001s
?   	github.com/frieser/camelrace/race	[no test files]
?   	github.com/frieser/camelrace/event	[no test files]
?   	github.com/frieser/camelrace/round	[no test files]
?   	github.com/frieser/camelrace/lane	[no test files]

```

Disclaimer
------
The test code is in the 'src/github.com/frieser' directory. The 'src/gopkg.in' is code of a vendor dependency.

Golang encourages short name variables when is no doubt of the name of them, for example in:

```go
func (e Event) Dispatch(){
        e.data
        //...
}
```

I choose name it event Event in favour of readability of the test code.

