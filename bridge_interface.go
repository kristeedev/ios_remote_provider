package main

import (
  "fmt"
  uj "github.com/nanoscopic/ujsonin/v2/mod"
)

type TunPair struct {
  from int
  to int
}

func (self TunPair) String() string {
  return fmt.Sprintf("%d:%d",self.from,self.to)
}

type Screenshot struct {
  format string
  data []byte
}

type BridgeDevInfo struct {
  udid string
}

// detect( onDevConnect func( bridge CliBridge ) )

type BridgeRoot interface {
  //OnConnect( dev BridgeDev )
  //OnDisconnect( dev BridgeDev )
  list() []BridgeDevInfo
}

type iProc struct {
  pid int32
  name string
}

type BridgeDev interface {
  getUdid() string
  tunnel( pairs []TunPair )
  info( names []string ) map[string]string
  gestalt( names []string ) map[string]string
  ps() []iProc
  screenshot() Screenshot
  wda( name string, port int, onStart func(), onStop func(interface{}) )
  destroy()
  setProcTracker( procTracker ProcTracker )
  NewBackupVideo( port int, onStop func( interface{} ) ) ( *BackupVideo )
  GetPid( appname string ) int
  AppInfo( bundleId string ) uj.JNode
  InstallApp( appPath string ) bool
  NewSyslogMonitor( handleLogItem func( uj.JNode ) )
}