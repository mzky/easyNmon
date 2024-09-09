package common

type MonData struct {
	Timestamp   Timestamp             `json:"timestamp"`
	Identity    Identity              `json:"identity"`
	OSRelease   OSRelease             `json:"os_release"`
	ProcVersion ProcVersion           `json:"proc_version"` //---
	LsCPU       LsCPU                 `json:"lscpu"`        // cpuInfo
	CPUInfo     map[string]CPUInfo    `json:"cpu_info"`
	CPUTotal    CPU                   `json:"cpu_total"`
	Cpus        map[string]CPU        `json:"cpus"`
	ProcMemInfo ProcMemInfo           `json:"proc_meminfo"`
	ProcVmStat  ProcVmStat            `json:"proc_vmstat"`
	Disks       map[string]Disks      `json:"disks"`
	Networks    map[string]Networks   `json:"networks"`
	Uptime      Uptime                `json:"uptime"`
	Filesystems map[string]Filesystem `json:"filesystems"`
}

type ProcMemInfo struct {
	MemTotal     int64 `json:"MemTotal"`
	MemFree      int64 `json:"MemFree"`
	MemAvailable int64 `json:"MemAvailable"`
	Buffers      int64 `json:"Buffers"`
	Cached       int64 `json:"Cached"`
	Active       int64 `json:"Active"`
	Inactive     int64 `json:"Inactive"`
	ActiveFile   int64 `json:"Active_file"`
	InactiveFile int64 `json:"Inactive_file"` // --以上/1024
	SwapTotal    int64 `json:"SwapTotal"`     // --以下/1024/1024
	SwapFree     int64 `json:"SwapFree"`
}

type ProcVmStat struct {
	PgpgIn  int64 `json:"pgpgin"`
	PgpgOut int64 `json:"pgpgout"`
	PswpIn  int64 `json:"pswpin"`
	PswpOut int64 `json:"pswpout"`
}

type CPUInfo struct {
	VendorId   string  `json:"vendor_id"`
	ModelName  string  `json:"model_name"`
	CpuMhz     float64 `json:"cpu_mhz"`
	CacheSize  float64 `json:"cache_size"`
	PhysicalId int     `json:"physical_id"`
	Siblings   int     `json:"siblings"`
	CoreId     int     `json:"core_id"`
	CpuCores   int     `json:"cpu_cores"`
}

type Disks struct {
	Reads  float64 `json:"reads"`
	Writes float64 `json:"writes"`
	RMerge float64 `json:"rmerge"`
	WMerge float64 `json:"wmerge"`
	Rkb    float64 `json:"rkb"`
	Wkb    float64 `json:"wkb"`
	RMsec  float64 `json:"rmsec"`
	WMsec  float64 `json:"wmsec"`
	Xfers  float64 `json:"xfers"`
}

type CPU struct {
	User      float64 `json:"user"`
	Nice      float64 `json:"nice"`
	Sys       float64 `json:"sys"`
	Idle      float64 `json:"idle"`
	IoWait    float64 `json:"iowait"`
	HardIrq   float64 `json:"hardirq"`
	SoftIrq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestnice"`
}

type Filesystem struct {
	FSFReqs       int64   `json:"fs_freqs"`
	FSPassNo      int64   `json:"fs_passno"`
	FSBSize       int64   `json:"fs_bsize"`
	FSBlocks      int64   `json:"fs_blocks"`
	FSBFree       int64   `json:"fs_bfree"`
	FSBAvail      int64   `json:"fs_bavail"`
	FSSizeMB      int64   `json:"fs_size_mb"`
	FSFreeMB      int64   `json:"fs_free_mb"`
	FSUsedMB      int64   `json:"fs_used_mb"`
	FSFullPercent float64 `json:"fs_full_percent"`
	FSAvail       int64   `json:"fs_avail"`
	FSFiles       int64   `json:"fs_files"`
	FSFilesFree   int64   `json:"fs_files_free"`
	FSNameLength  int64   `json:"fs_namelength"`
}

type Identity struct {
	Hostname     string `json:"hostname"`
	FullHostname string `json:"fullhostname"`
	Ipaddress    string `json:"ipaddress"`
	NjmonCommand string `json:"njmon_command"`
	NjmonMode    string `json:"njmon_mode"`
	NjmonVersion string `json:"njmon_version"`
	Username     string `json:"username"`
	Userid       int64  `json:"userid"`
	Model        string `json:"model"`
	Vendor       string `json:"vendor"`
}

type LsCPU struct {
	Architecture   string `json:"architecture"`
	ByteOrder      string `json:"byte_order"`
	Cpus           string `json:"cpus"`
	OnlineCPUList  string `json:"online_cpu_list"`
	VendorID       string `json:"vendor_id"`
	ModelName      string `json:"model_name"`
	CPUFamily      string `json:"cpu_family"`
	Model          string `json:"model"`
	ThreadsPerCore string `json:"threads_per_core"`
	CoresPerSocket string `json:"cores_per_socket"`
	Sockets        string `json:"sockets"`
	Stepping       string `json:"stepping"`
	BoGomIps       string `json:"bogomips"`
	NUMANodes      string `json:"numa_nodes"`
}

type Networks struct {
	IBytes   float64 `json:"ibytes"`
	OBytes   float64 `json:"obytes"`
	IPackets float64 `json:"ipackets"`
	OPackets float64 `json:"opackets"`
	IDrop    float64 `json:"idrop"`
	ODrop    float64 `json:"odrop"`
	Ififo    float64 `json:"ififo"`
	OFifo    float64 `json:"ofifo"`
}

type OSRelease struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	PrettyName string `json:"pretty_name"`
	VersionID  string `json:"version_id"`
}

type ProcVersion struct {
	Version string `json:"version"`
}

type Timestamp struct {
	Datetime         string  `json:"datetime"`
	UTC              string  `json:"UTC"`
	SnapshotSeconds  int64   `json:"snapshot_seconds"`
	SnapshotMaxLoops int64   `json:"snapshot_maxloops"`
	SnapshotLoop     int64   `json:"snapshot_loop"`
	Sleeping         float64 `json:"sleeping"`
	ExecuteTime      float64 `json:"execute_time"`
	SleepOverrun     float64 `json:"sleep_overrun"`
	Elapsed          float64 `json:"elapsed"`
}

type Uptime struct {
	Days    int64 `json:"days"`
	Hours   int64 `json:"hours"`
	Minutes int64 `json:"minutes"`
	Users   int64 `json:"users"`
}

// Occupy 定义一个输出的结构体来匹配JSON数据
type Occupy struct {
	CPUTotal struct {
		User   []float64 `json:"user"`
		Sys    []float64 `json:"sys"`
		Iowait []float64 `json:"iowait"`
	} `json:"cpu_total"`
	ProcMeminfo struct {
		MemFree  []float64 `json:"MemFree"`
		Active   []float64 `json:"Active"`
		MemTotal []float64 `json:"MemTotal"`
	} `json:"proc_meminfo"`
	Networks map[string]struct {
		Ibytes []float64 `json:"ibytes"`
		Obytes []float64 `json:"obytes"`
	} `json:"networks"`
	Disks map[string]struct {
		Reads  []float64 `json:"reads"`
		Writes []float64 `json:"writes"`
	} `json:"disks"`
	Name string `json:"name"`
}

type Rsp struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RspOK(message string, data interface{}) *Rsp {
	return &Rsp{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

var WebRoot = "/web"
