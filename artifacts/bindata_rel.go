// +build rel

//go:generate go-bindata -o bindata_generated.go -pkg=artifacts -prefix=data/ data/riak_mesos_executor.tar.gz data/riak.conf data/advanced.config data/riak-bin.tar.gz data/riak-basho-patches.tar.gz

package artifacts
