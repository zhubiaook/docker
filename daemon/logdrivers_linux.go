package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/zhubiaook/docker/daemon/logger/awslogs"
	_ "github.com/zhubiaook/docker/daemon/logger/fluentd"
	_ "github.com/zhubiaook/docker/daemon/logger/gcplogs"
	_ "github.com/zhubiaook/docker/daemon/logger/gelf"
	_ "github.com/zhubiaook/docker/daemon/logger/journald"
	_ "github.com/zhubiaook/docker/daemon/logger/jsonfilelog"
	_ "github.com/zhubiaook/docker/daemon/logger/local"
	_ "github.com/zhubiaook/docker/daemon/logger/logentries"
	_ "github.com/zhubiaook/docker/daemon/logger/loggerutils/cache"
	_ "github.com/zhubiaook/docker/daemon/logger/splunk"
	_ "github.com/zhubiaook/docker/daemon/logger/syslog"
)
