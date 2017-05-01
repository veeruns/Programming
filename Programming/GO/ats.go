package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "net"
    "time"
    "runtime"
    "sort"
)

// Configuration Files:
// ErrorCodes, API, 
//
//**********************************************************************
type days []uintptr
type ByteSize float64

const perYrMinute int    = 1*24*60
// var perYrMinute int   = 1*24*time.Minute
const ipv4_class uint8   = 255
var httpCodes = map[string]uint{
    "100":  0,
    "200":  0,
    "400":  0,
    "500":  0,
}

var httpCodeStrings = map[string]string{
    "100":  "Redirect",
    "200":  "Ok", 
    "400":  "Not Found",
    "500":  "Illegal Call",
}

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

const (
  maxProxyRxNodes = 1024  // Gates that allow/get traffic from remote ATS.
)

//**********************************************************************
type edgeResult struct {
    hname     string
    dConn     net.Conn
    connStatus bool
    nodeAddr  net.TCPAddr
    ecode     error
    line      []byte
}

type atsParse struct {
    minutes [perYrMinute] uint32
    daemonStartTime       uint32
}

type EdgeHosts struct {
  edgeHostsAll          map[string] []string
  edgeHostsColo         map[string] uint16
  edgeHostsColoSorted   [] string
  edgeIPV4Hosts         map[string] []string
  edgeIPV6Hosts         map[string] []string
}

type atsReadLine struct {
  currentLineNo   uint32;
  currentIP       string;
  totalLineNo     uint32;
  thisMinute      int
  thisMinuteText  string;
  logLineIP       string;
}

type flowstats struct {
    client_24 [ipv4_class] uint
    client_16 [ipv4_class] uint
    client_8  [ipv4_class] uint
    client_0  [ipv4_class] uint
    ResponseCode      map[string]uint
    ymailAPI          map[string]uint
    UserAgent         map[string]uint
    EdgeNodeBW        map[string] uint64

    clientbw_24 [ipv4_class] uint
    clientbw_16 [ipv4_class] uint
    clientbw_8  [ipv4_class] uint
    clientbw_0  [ipv4_class] uint

    latency  [ipv4_class][32] uint
}

type ipv4_clientstats interface {
    ipv4ClassStats (ipv4Count int, flow *flowstats)
}

type httpResponseStats interface {
    httpStats (CodeStats int, flow *flowstats)
}

type AccessLogStats interface {
    ipv4_clientstats
    httpResponseStats
}

//**********************************************************************
// https://golang.org/ref/spec

var atsStats  [perYrMinute]       flowstats
var atsParser [maxProxyRxNodes]   atsReadLine
var EdgeNodes EdgeHosts
var edgeFile = "/Users/veeru/Programming/GO/ycpi.csv"

var mgHosts    map[string] net.IP
var mgFile = "/Users/veeru/Programming/GO/mg.csv"
var mgport int = 4080

var httpCodeCount = len(httpCodes)
var tickChan *time.Ticker

//**********************************************************************
func initBoot() {
    EdgeNodes.edgeHostsAll  = make(map[string] []string)
    EdgeNodes.edgeHostsColo = make(map[string] uint16)
    EdgeNodes.edgeIPV4Hosts = make(map[string] []string)
    EdgeNodes.edgeIPV6Hosts = make(map[string] []string)
    mgHosts                 = map[string] net.IP {}

    // atsStats[0].EdgeReference  = &EdgeNodes.edgeIPV4Hosts
    // fmt.Println("Done Init boot")
}
//**********************************************************************
func initPhase1() {
  rec_count, err := edgeNodesCsv(edgeFile)
  if err != nil || rec_count == 0 {
    fmt.Println(rec_count, err)
  } else {
    fmt.Println("Number of Edge nodes: ", rec_count)
    // for h := range EdgeNodes.edgeIPV4Hosts {
    //  fmt.Printf("%24s %s\n", h, EdgeNodes.edgeIPV4Hosts[h])
    // }

    //******************************************
    // Edge Details.
    lenAts  := len(EdgeNodes.edgeIPV4Hosts)
    lenColo := len(EdgeNodes.edgeHostsColo)
    EdgeNodes.edgeHostsColoSorted = make([]string, lenColo);
    
    i := 0
    for k, _ := range EdgeNodes.edgeHostsColo {
      EdgeNodes.edgeHostsColoSorted[i] = k
      i++
    }
    sort.Strings(EdgeNodes.edgeHostsColoSorted)
    fmt.Println("Edge Nodes: ", lenAts, "; Colo: ", lenColo)
  }

  //**********************************************
  // First Collect the inner edge node details where ATS is running.
  if err, mgC := collectInnerEdge(mgFile); err != nil {
    fmt.Println("Cannot InnerEdge Nodes: ", err)
    os.Exit(0)
  } else {
    fmt.Println("mgHosts = ", len(mgHosts), mgC)
  }
}
//**********************************************************************
// Start a goroutine to read from our net connection
func readLoop(connDetails edgeResult, dCh chan edgeResult, eCh chan edgeResult) {

  conn := connDetails.dConn
  name := connDetails.hname
  fmt.Fprintf(conn, "Vatsan\n")

  reader := bufio.NewReader(conn)
  for {
    strLine,err := reader.ReadString('\n')
    res         := new(edgeResult)
    res.dConn    = conn
    res.hname    = name
    res.nodeAddr = connDetails.nodeAddr

    if err != nil {
      // send an error if it's encountered
      res.ecode = err
      eCh <- *res
      return 
    }

    res.line = []byte(strLine)
    // send data if we read some.
    dCh <- *res
  }
}
//**********************************************************************
func tickTock(chan2 chan string) {
  for t := range tickChan.C {
    chan2 <- "Timer Ticked " + fmt.Sprint(t)
  }
}
//**********************************************************************
func initTimer() {
    // timeChan := time.NewTimer(time.Second).C
    tickChan = time.NewTicker(time.Millisecond * 1000)
}
//**********************************************************************
func  edgeNodesCsv(file string) (uint32, error) {
    var count uint32

    fd, err := os.Open(file)
    if err != nil {
        return  0, err
    }

    defer fd.Close()
    scanner := bufio.NewScanner(fd)
    count = 0

    for scanner.Scan() {
        hname := scanner.Text()
        ipAddr, err := net.LookupHost(hname)
        if err != nil {
            fmt.Println(err)
            continue
        }

        // r10.xxx.<colo>.domain.net. we will extract <colo> here so cluster level report
        // can be done for bandwidth from edge nodes.
        fqdnEdge := strings.Split(hname, ".")
        x1       := fqdnEdge[len(fqdnEdge)-3]
        EdgeNodes.edgeHostsColo[x1]++;

        EdgeNodes.edgeHostsAll[hname] = ipAddr
        for v := range ipAddr {
            // fmt.Println(hname, ipAddr[v])

            // Extract array of strings (IP address) to string that has 1 field.
            x :=  strings.Fields(ipAddr[v])

            // Contains wont work with single string.
            // Group IPV4 and IPV6 separately and hash map hostname to IP address.
            if strings.Contains(ipAddr[v], ":") == false {
                EdgeNodes.edgeIPV4Hosts[hname] = x
                count++
            } else {
                EdgeNodes.edgeIPV6Hosts[hname] = x
            }
        }
    }

    return count, err
}
//**********************************************************************
// func edgeConnect(src map[string]net.IP, dst map[string] edgeResult ) {
// Since map will not work with updating elements, using slice for dst.

func edgeConnect(src map[string]net.IP, dst [] edgeResult ) uint32 {
  var index uint32 = 0

  for name, _ := range src {
    hname := name + ":" + strconv.Itoa(mgport)
    tcpAddr, err := net.ResolveTCPAddr("tcp", hname)
    if err != nil { 
        panic(err) 
    } 
    fmt.Printf("Connecting...%19s", tcpAddr)

    // connect to this socket
    tcpAddr.Port = mgport
    // conn, rc := net.DialTCP("tcp", nil, tcpAddr)
    conn, rc := net.DialTimeout("tcp", hname, time.Second)
    if rc != nil {
        fmt.Printf(" - Disconnect - %19s\n", tcpAddr)
        continue

        // os.Exit(0) 
    }
    
    dst[index].hname      = name
    dst[index].dConn      = conn
    dst[index].nodeAddr   = *tcpAddr
    dst[index].connStatus = true
    index++
    fmt.Printf(" - Connected  - %19s\n", tcpAddr)
  }

  fmt.Println("Connected: ", index, "Nodes")
  return index
}
//**********************************************************************
func collectInnerEdge(file string) (error, uint32) {
  var count uint32

  fd, err := os.Open(file)
  if err != nil {
      return  err, 0
  }

  defer fd.Close()
  scanner := bufio.NewScanner(fd)

  count = 0
  for scanner.Scan() {
      name := scanner.Text()
      hname := name + ":" + strconv.Itoa(mgport)
      tcpAddr, err := net.ResolveTCPAddr("tcp", hname)
      if err != nil {
          fmt.Println(err)
          continue
      }

      // fmt.Println("Host = ", name, tcpAddr.IP)
      mgHosts[name] = tcpAddr.IP
      count++
  }

  return nil, count
}
//**********************************************************************
func readLineAtsLogs(thisNode atsReadLine) uint32 {
  var lineCount uint32 = 0
  var lmin string;

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)       
  for scanner.Scan() {
    lineCount++
    words := strings.Split(scanner.Text(), "\t")
    for i := range words {
      // fmt.Println(words[i])

      // First Field is time with NO "=" ; Rest of the fields have "k=v" model.
      // many fields have k=v model and let us get it done first.
      kv := strings.Split(words[i], "=")
      k, v := kv[0], kv[1:]

      // First Field is Time. Test it and Extract the logged time.
      // Extract the minute field for our running index of 1 minute stats.
      if len(v) != 0 {
            atsCountFlds(k,v, thisNode)
      } else {
        kv    = strings.Split(k, " ")
        tlog := strings.Split(kv[1], ":")
        lmin  = tlog[1]

        if (lmin != thisNode.thisMinuteText)  {
          x1,err1 := strconv.Atoi(tlog[0])
          x2,err2 := strconv.Atoi(tlog[1])

          if err1 == nil && err2 == nil {
            //thisMinute     = x1*60 + x2
            thisNode.thisMinute     = x1*60 + x2
            thisNode.thisMinuteText = lmin
            thisNode.logLineIP      = ""

            fmt.Println("Line: ", lineCount, k, thisNode.thisMinuteText, thisNode.thisMinute)
          }
        } 
      }
    }
    // fmt.Println("Total Line Count = ", lineCount)
  }

  return lineCount
}
//**********************************************************************
func atsCountFlds(c string, v []string, thisNode atsReadLine) {
    switch (c) {
        case "H"  :
            // fmt.Println("H = ", c,v)
            thisNode.logLineIP = v[0]
            ipv4 := strings.Split(v[0], ".")

            x1,err1 := strconv.Atoi(ipv4[0]); x2,err2 := strconv.Atoi(ipv4[1]); 
            x3,err3 := strconv.Atoi(ipv4[2]); x4,err4 := strconv.Atoi(ipv4[3]);

            if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
              atsStats[thisNode.thisMinute].client_24[x1]++
              atsStats[thisNode.thisMinute].client_16[x2]++
              atsStats[thisNode.thisMinute].client_8[x3]++
              atsStats[thisNode.thisMinute].client_0[x4]++
              // fmt.Println(err1,err2,err3,err4)
            }
        case "A"  :
            // fmt.Println("A = ", c,v)
        case "J"  :
            // fmt.Println("J = ", c,v)
        case "N"  :
            // fmt.Println("N = ", c,v)
        case "n"  :
            // fmt.Println("n = ", c,v)
        case "O"  :
            // fmt.Println("O = ", c,v)
        case "Q"  :
            // fmt.Println("Q = ", c,v)
        case "v"  :
            // fmt.Println("v = ", c,v)
        case "w"  :
            // fmt.Println("w = ", c,v)
        case "g"  :
            // fmt.Println("g = ", c,v)
        case "K"  :
            // fmt.Println("K = ", c,v)
        case "s"  :
            // fmt.Println("s = ", c,v)
        case "m"  :
            // fmt.Println("m = ", c,v)
        case "cr" :
            // fmt.Println("rest = ", c,v)
    }
}
//**********************************************************************
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
//**********************************************************************
func main() {
  dCh := make(chan edgeResult)
  eCh := make(chan edgeResult)
  tCh := make(chan string)
  mgHosts = map[string] net.IP {}

  var lineCounted uint32 =  0 
  var bsize int =  0 
  var rxNode edgeResult

  cpuMax:= runtime.NumCPU()
  runtime.GOMAXPROCS(cpuMax)

  initBoot()
  initPhase1()

  //**********************************************
  // Try to connect to edge nodes using Slice Only.
  txNode := make([]edgeResult, len(mgHosts))
  rc := edgeConnect(mgHosts, txNode) 

  if rc == 0 {
    fmt.Println("Cannot Connect to any node")
    os.Exit(0)
  }

  fmt.Println("CPU Max: ", cpuMax, "Total Inner Nodes: ", len(txNode) )

  for hin, edge  := range txNode {
    // Run a goroutine loop to read.
    if (edge.connStatus == true) {
    	fmt.Println("Go Routine: ", hin, edge.hname)
    	go readLoop(txNode[hin], dCh, eCh)
    }
  }

  tickChan = time.NewTicker(time.Second * 5)
  go tickTock(tCh) 

  // continuously read from the connection
  for {
    select {
    // This case means we recieved data on the connection
    case rxNode = <-dCh:
        // fmt.Println("Received Size: ", len(rxNode.atsLine) )
        lineCounted++
        bsize = bsize + len(rxNode.line)
        break
    // This case means we got an error and the goroutine has finished
    case rxNode := <-eCh:
      // handle our error then exit for loop
      fmt.Println("Error is: ", rxNode.hname, rxNode.ecode)
      break;
     // This will timeout on the read.
    case tout := <- tCh:
      fmt.Println("Timeout is: ", tout, ": ", lineCounted, ": ", bsize, "= ", rxNode.nodeAddr.IP, rxNode.nodeAddr.Port)
    }
  }


/***********************************************************************
    time.Sleep(time.Millisecond * 1000)
    lineC := readLineAtsLogs(atsParser[0])
    fmt.Println("Len of line", lineC, cap(atsStats) )

    for i := range atsStats {
        for j := range atsStats[i].client_24 {
          if  atsStats[i].client_24[j] > 0 {
            fmt.Printf("ip24 = %4d %10d\n", i, atsStats[i].client_24[j])
          }
      }
    }
***********************************************************************/

}
