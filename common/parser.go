package common

import (
	"log"

	"github.com/mzky/utils/memdb"
	"github.com/tidwall/gjson"
)

var (
	SysInfos = []string{"identity", "tags", "os_release", "proc_version", "lscpu", "uptime"}
	l2       = []string{"timestamp", "cpu_total", "stat_counters", "loadavg", "proc_meminfo"}
	l3       = []string{"cpus", "disks", "networks", "filesystems"}
	Occupys  = append(l2, l3...)
)

type Mem struct {
	*memdb.DB
}

func (m *Mem) InsertInfo(result gjson.Result, pkey string) {
	if i, err := m.DB.Get(pkey); err != nil && i == nil {
		m.InsertData(result, pkey)
	}
}

func (m *Mem) InsertData(result gjson.Result, pkey string) {
	for pn, r := range result.Map() {
		for sn, sr := range r.Map() {
			switch sr.Value().(type) {
			case string:
				Handle(m.DB.Append(sr.String(), pkey, pn, sn))
			case float64:
				Handle(m.DB.Append(sr.Float(), pkey, pn, sn))
			}
		}
		switch r.Value().(type) {
		case string:
			Handle(m.DB.Append(r.String(), pkey, pn))
		case float64:
			Handle(m.DB.Append(r.Float(), pkey, pn))
		}
	}
}

func (m *Mem) Parser(data []byte) {
	m.Parse(data)
}

func Handle(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (m *Mem) Parse(j []byte) {
	g := gjson.ParseBytes(j)
	for _, v := range SysInfos {
		m.InsertInfo(g.Get(v), "SysInfo."+v)
	}
	for _, v := range Occupys {
		m.InsertData(g.Get(v), "Occupy."+v)
	}
}
