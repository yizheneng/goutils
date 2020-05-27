package goutils

/*
#include <time.h>
#include <stdint.h>
#ifdef _WIN32
#include <windows.h>
//#pragma comment(lib, "winmm.lib")
#cgo LDFLAGS: -lwinmm
#endif

uint64_t systemUptime()
{
#ifdef _WIN32
	return timeGetTime();
#else
	struct timespec tp;
	clock_gettime(CLOCK_MONOTONIC, &tp);
	return (long)tp.tv_sec * 1000 + tp.tv_nsec / 1000000;
#endif
}
*/
import "C"

type ElapsedTimer struct {
	timeval uint64
	lastTimer uint64
	stopFlag bool
}

func NewElapsedTimerWithVal(val uint64) (this *ElapsedTimer) {
	this = &ElapsedTimer{timeval:val, lastTimer:SystemUpTime(), stopFlag:false}
	this.timeval = val
	return
}

func (this *ElapsedTimer)Restart() {
	this.stopFlag = false
	this.Start()
}

func (this *ElapsedTimer)IsTimeout() (bool) {
	if(this.stopFlag) {
		return false
	} else {
		return (SystemUpTime() - this.lastTimer) > this.timeval
	}
}

func (this *ElapsedTimer)SetTimeoutVal(timeval uint64){
	this.timeval = timeval
	this.Restart()
}

func (this *ElapsedTimer)Stop() {
	this.stopFlag = true
}

func (this *ElapsedTimer)Start() {
	this.lastTimer = SystemUpTime()
}

func (this *ElapsedTimer)Elapsed() (val uint64) {
	val = SystemUpTime() - this.lastTimer
	return
}

func SystemUpTime() uint64 {
	return uint64(C.systemUptime())
}
