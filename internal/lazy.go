package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/shopspring/decimal"
)

type NmonReport struct {
	ScriptName string
	XAxisdatas string
	CPUUsers   string
	CPUSyss    string
	CPUWaits   string
	Memfrees   string
	Actives    string
	Memtotals  string
	NetReads   string
	NetWrites  string
	DiskReads  string
	DiskWrites string
}

func GenIndexPage(nr *NmonReport, fPath string) {
	tpl := template.Must(template.New("data.tpl").ParseFiles(filepath.Join("web", "chart", "data.tpl")))
	file, err := os.Create(filepath.Join(fPath, "data.json"))
	os.Chmod(filepath.Join(fPath, "data.json"), os.ModePerm)

	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(file, nr)
	if err != nil {
		log.Println(err)
	}
}

func GetNmonReport(filePath string, name string) {
	//fileName := filepath.Join(filePath, name)
	fileName := GetFiles(filePath, name)
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	hasZZZZ := false
	hasCPUAll := false
	hasDiskRead := false
	hasDiskWrite := false
	hasMem := false
	hasNet := false
	indexNetRead := make([]int, 0)
	indexNetWrite := make([]int, 0)
	sliceZZZZTime := make([]string, 0, 1024)
	// maxCPUUsage, averageCPUUsage, minCPUUsage := 0.0, 0.0, 100.0
	sliceCPUUser := make([]float64, 0, 1024)
	sliceCPUSys := make([]float64, 0, 1024)
	sliceCPUWait := make([]float64, 0, 1024)
	// sliceCPUUsage := make([]float64, 0, 1024)
	sliceDiskRead := make([]float64, 0, 1024)
	sliceDiskWrite := make([]float64, 0, 1024)
	// maxMemUsage, averageMemUsage, minMemUsage, memoryTotal := 0.0, 0.0, 100.0, 0.0
	sliceMemTotal := make([]float64, 0, 1024)
	sliceMemFree := make([]float64, 0, 1024)
	// sliceMemCached := make([]float64, 0, 1024)
	sliceMemActive := make([]float64, 0, 1024)
	// sliceMemBuffers := make([]float64, 0, 1024)
	// sliceMemUsage := make([]float64, 0, 1024)
	sliceNetReadTotal := make([]float64, 0, 1024)
	sliceNetWriteTotal := make([]float64, 0, 1024)
	for {
		line, isPrefix, err := reader.ReadLine()
		// 解决单行字节数大于4096的情况
		for isPrefix && err == nil {
			var bs []byte
			bs, isPrefix, err = reader.ReadLine()
			line = append(line, bs...)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return
		}
		strLine := string(line)
		arr := strings.Split(strLine, ",")
		if !hasZZZZ && strings.HasPrefix(strLine, "ZZZZ,") {
			hasZZZZ = true
			// ZZZZ没有标题栏，所以不需要continue
		}
		if !hasCPUAll && strings.HasPrefix(strLine, "CPU_ALL,CPU Total") {
			hasCPUAll = true
			continue
		}
		if !hasDiskRead && strings.HasPrefix(strLine, "DISKREAD,Disk Read KB/s") {
			hasDiskRead = true
			continue
		}
		if !hasDiskWrite && strings.HasPrefix(strLine, "DISKWRITE,Disk Write KB/s") {
			hasDiskWrite = true
			continue
		}
		if !hasMem && strings.HasPrefix(strLine, "MEM,Memory MB") {
			hasMem = true
			continue
		}
		if !hasNet && strings.HasPrefix(strLine, "NET,Network I/O") {
			hasNet = true
			// 或许存在多个网络适配器
			for i, columnName := range arr {
				if strings.HasSuffix(columnName, "-read-KB/s") {
					indexNetRead = append(indexNetRead, i)
				}
				if strings.HasSuffix(columnName, "-write-KB/s") {
					indexNetWrite = append(indexNetWrite, i)
				}
			}
			continue
		}
		if hasZZZZ && strings.HasPrefix(strLine, "ZZZZ,") {
			sliceZZZZTime = append(sliceZZZZTime, "\""+arr[2]+"\"")
			continue
		}
		if hasCPUAll && strings.HasPrefix(strLine, "CPU_ALL,") {
			uu := GetFloatFromString(arr[2])
			su := GetFloatFromString(arr[3])
			wu := GetFloatFromString(arr[4])
			// use := SumOfFloat(uu, su)
			sliceCPUUser = append(sliceCPUUser, uu)
			sliceCPUSys = append(sliceCPUSys, su)
			sliceCPUWait = append(sliceCPUWait, wu)
			// sliceCPUUsage = append(sliceCPUUsage, use)
			// if maxCPUUsage < use {
			// 	maxCPUUsage = use
			// }
			// if minCPUUsage > use {
			// 	minCPUUsage = use
			// }
			// averageCPUUsage = SumOfFloat(averageCPUUsage, use)
			continue
		}
		if hasDiskRead && strings.HasPrefix(strLine, "DISKREAD,") {
			sliceDiskRead = append(sliceDiskRead, SumOfEachColumns(strLine))
			continue
		}
		if hasDiskWrite && strings.HasPrefix(strLine, "DISKWRITE,") {
			sliceDiskWrite = append(sliceDiskWrite, SumOfEachColumns(strLine)*-1)
			continue
		}
		if hasMem && strings.HasPrefix(strLine, "MEM,") {
			mTotal, _ := decimal.NewFromString(arr[2])
			mFree, _ := decimal.NewFromString(arr[6])
			// mCached, _ := decimal.NewFromString(arr[11])
			mActive, _ := decimal.NewFromString(arr[12])
			// mBuffers, _ := decimal.NewFromString(arr[14])
			// mUsage := mTotal.Sub(mFree).Sub(mCached).Sub(mBuffers).DivRound(mTotal, 4).Mul(decimal.NewFromFloat32(100))
			// nu, _ := mUsage.Float64()
			// if memoryTotal == 0.0 {
			// 	memoryTotal = GetFloatFromDecimal(mTotal)
			// }
			// if maxMemUsage < nu {
			// 	maxMemUsage = nu
			// }
			// if minMemUsage > nu {
			// 	minMemUsage = nu
			// }
			// averageMemUsage = SumOfFloat(averageMemUsage, nu)
			sliceMemTotal = append(sliceMemTotal, GetFloatFromDecimal(mTotal))
			sliceMemFree = append(sliceMemFree, GetFloatFromDecimal(mFree))
			// sliceMemCached = append(sliceMemCached, GetFloatFromDecimal(mCached))
			sliceMemActive = append(sliceMemActive, GetFloatFromDecimal(mActive))
			// sliceMemBuffers = append(sliceMemBuffers, GetFloatFromDecimal(mBuffers))
			// sliceMemUsage = append(sliceMemUsage, nu)
			continue
		}
		if hasNet && strings.HasPrefix(strLine, "NET,") {
			sliceNetReadTotal = append(sliceNetReadTotal, SumOfSpecifiedColumns(strLine, indexNetRead))
			sliceNetWriteTotal = append(sliceNetWriteTotal, SumOfSpecifiedColumns(strLine, indexNetWrite)*-1)
			continue
		}

	}

	nr := new(NmonReport)

	//if !hasZZZZ {
	//	log.Println("解析nmon结果文件失败")
	//	return
	//} else {
	if hasZZZZ {
		nr.XAxisdatas = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceZZZZTime), "[]"), " ", ",")
	}

	if hasCPUAll {
		nr.CPUUsers = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceCPUUser), "[]"), " ", ",")
		nr.CPUSyss = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceCPUSys), "[]"), " ", ",")
		nr.CPUWaits = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceCPUWait), "[]"), " ", ",")
	}

	if hasMem {
		nr.Memfrees = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceMemFree), "[]"), " ", ",")
		nr.Actives = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceMemActive), "[]"), " ", ",")
		nr.Memtotals = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceMemTotal), "[]"), " ", ",")
	}

	if hasNet {
		nr.NetReads = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceNetReadTotal), "[]"), " ", ",")
		nr.NetWrites = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceNetWriteTotal), "[]"), " ", ",")
	}

	if hasDiskRead && hasDiskWrite {
		nr.DiskReads = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceDiskRead), "[]"), " ", ",")
		nr.DiskWrites = strings.ReplaceAll(strings.Trim(fmt.Sprint(sliceDiskWrite), "[]"), " ", ",")
	}
	if nr != nil {
		nr.ScriptName = name
		GenIndexPage(nr, filePath)
	}
}

// GetFloatFromString 字符串转float64
func GetFloatFromString(value string) float64 {
	decimal.DivisionPrecision = 2
	n, _ := decimal.NewFromString(value)
	ret, _ := n.Float64()
	return ret
}

// GetFloatFromDecimal decimal.Decimal转float64
func GetFloatFromDecimal(value decimal.Decimal) float64 {
	decimal.DivisionPrecision = 2
	ret, _ := value.Float64()
	return ret
}

// SumOfFloat 计算float的和
func SumOfFloat(value ...float64) float64 {
	decimal.DivisionPrecision = 2
	sum := decimal.NewFromFloat32(0)
	for _, v := range value {
		sum = sum.Add(decimal.NewFromFloat(v))
	}
	ret, _ := sum.Float64()
	return ret
}

// SumOfEachColumns 返回当前行的列之和(不包含前两列)
func SumOfEachColumns(line string) float64 {
	decimal.DivisionPrecision = 2
	arr := strings.Split(line, ",")
	sum := decimal.NewFromFloat32(0)
	for i := 2; i < len(arr); i++ {
		n, err := decimal.NewFromString(arr[i])
		if err != nil {
			log.Println(err, "该值将当作0完成后续计算")
			n = decimal.NewFromFloat32(0)
		}
		sum = sum.Add(n)
	}
	ret, _ := sum.Float64()
	return ret
}

// SumOfSpecifiedColumns 返回当前行的指定列之和
func SumOfSpecifiedColumns(line string, columns []int) float64 {
	decimal.DivisionPrecision = 2
	arr := strings.Split(line, ",")
	sum := decimal.NewFromFloat32(0)
	for _, index := range columns {
		n, _ := decimal.NewFromString(arr[index])
		sum = sum.Add(n)
	}
	ret, _ := sum.Float64()
	return ret
}
