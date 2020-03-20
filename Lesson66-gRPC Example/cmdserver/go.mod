module cmdserver

go 1.13

replace cmdexecutor => ../cmdexecutor

require (
	cmdexecutor v0.0.0
	google.golang.org/grpc v1.28.0
)
