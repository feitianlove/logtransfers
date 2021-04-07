module github.com/feitianlove/logtransfers

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.28.0
	github.com/feitianlove/golib v0.0.0-20210407155501-7b1e4950d1b6
	github.com/golang/snappy v0.0.3 // indirect
	github.com/klauspost/compress v1.11.13 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
