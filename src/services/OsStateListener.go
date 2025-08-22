package services

import (
	"runtime"
	"time"
	"voxesis/src/core/entity"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetOsStateListener() (*entity.OsState, error) {
	state := &entity.OsState{}

	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	if len(cpuPercent) > 0 {
		state.CpuUsage = cpuPercent[0]
	}

	// 获取物理CPU核心数和CPU名称
	cpuInfo, err := cpu.Info()
	if err != nil {
		// 如果无法获取CPU详细信息，则回退到逻辑CPU数
		state.CpuCores = int64(runtime.NumCPU())
	} else if len(cpuInfo) > 0 {

		// 获取物理CPU核心数
		// 使用 false 参数获取物理核心数，true 获取逻辑核心数
		physicalCores, err := cpu.Counts(false)
		if err != nil {
			// 回退到使用 Info 中的信息
			state.CpuCores = int64(cpuInfo[0].Cores)
		} else {
			state.CpuCores = int64(physicalCores)
		}
	} else {
		// 如果没有获取到CPU信息，也回退到逻辑CPU数
		state.CpuCores = int64(runtime.NumCPU())
	}

	// 获取内存信息
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	state.MemoryUsage = memStat.UsedPercent
	state.OsMemory = float64(memStat.Total) / 1024 / 1024 // 转换为MB

	return state, nil
}
