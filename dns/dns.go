package main

import (
	"github.com/coredns/coredns/coremain"
	"github.com/coredns/coredns/core/dnsserver"

	_ "github.com/coredns/coredns/plugin/auto"
	_ "github.com/coredns/coredns/plugin/autopath"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/dnssec"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/hosts"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/metadata"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/proxy"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/root"
	_ "github.com/coredns/coredns/plugin/secondary"
	_ "github.com/coredns/coredns/plugin/template"
	_ "github.com/coredns/coredns/plugin/tls"
	_ "github.com/coredns/coredns/plugin/trace"
	_ "github.com/coredns/coredns/plugin/whoami"
	_ "github.com/coredns/fallback"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/hmilkovi/rds"
)

var directives = []string{
	"metadata",
	"tls",
	"reload",
	"nsid",
	"root",
	"bind",
	"trace",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"nspub",
	"loadbalance",
	"cache",
	"dnssec",
	"rewrite",
	"autopath",
	"template",
	"hosts",
	"file",
	"auto",
	"secondary",
	"forward",
	"proxy",
	"whoami",
	"fallback",
	"pdsql_mysql",
	"pdsql_sqlite",
	"pdsql",
}

func init() {
    dnsserver.Directives = directives
}

func main() {
    coremain.Run()
}