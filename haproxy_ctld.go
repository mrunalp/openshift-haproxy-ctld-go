package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

type HAProxyAttr struct {
	pxname, svname, qcur, qmax, scur, smax, slim, stot, bin, bout, dreq, dresp, ereq, econ, eresp, wretr, wredis, status, weight, act, bck, chkfail, chkdown, lastchg, downtime, qlimit, pid, iid, sid, throttle, lbtot, tracked, htype, rate, rate_lim, rate_max, check_status, check_code, check_duration, hrsp_1xx, hrsp_2xx, hrsp_3xx, hrsp_4xx, hrsp_5xx, hrsp_other, hanafail, req_rate, req_rate_max, req_tot, cli_abrt, srv_abrt string
}

func (h *HAProxyAttr) Initialize(sl []string) {
	h.pxname, h.svname, h.qcur, h.qmax, h.scur, h.smax, h.slim, h.stot, h.bin, h.bout, h.dreq, h.dresp, h.ereq, h.econ, h.eresp, h.wretr, h.wredis, h.status, h.weight, h.act, h.bck, h.chkfail, h.chkdown, h.lastchg, h.downtime, h.qlimit, h.pid, h.iid, h.sid, h.throttle, h.lbtot, h.tracked, h.htype, h.rate, h.rate_lim, h.rate_max, h.check_status, h.check_code, h.check_duration, h.hrsp_1xx, h.hrsp_2xx, h.hrsp_3xx, h.hrsp_4xx, h.hrsp_5xx, h.hrsp_other, h.hanafail, h.req_rate, h.req_rate_max, h.req_tot, h.cli_abrt, h.srv_abrt = sl[0], sl[1], sl[2], sl[3], sl[4], sl[5], sl[6], sl[7], sl[8], sl[9], sl[10], sl[11], sl[12], sl[13], sl[14], sl[15], sl[16], sl[17], sl[18], sl[19], sl[20], sl[21], sl[22], sl[23], sl[24], sl[25], sl[26], sl[27], sl[28], sl[29], sl[30], sl[31], sl[32], sl[33], sl[34], sl[35], sl[36], sl[37], sl[38], sl[39], sl[40], sl[41], sl[42], sl[43], sl[44], sl[45], sl[46], sl[47], sl[48], sl[49], sl[50]
}

func main() {
	fmt.Println("Starting HA control daemon")
	haproxy_dir_path := os.Getenv("OPENSHIFT_HAPROXY_DIR")
	fmt.Println(haproxy_dir_path)
	stats_socket := path.Join(haproxy_dir_path, "run", "stats")
	fmt.Println(stats_socket)

	for {
		c, err := net.Dial("unix", stats_socket)
		handleError(err)
		defer c.Close()

		fmt.Println("Connected to stats_socket")
		fmt.Fprintf(c, "show stat\n")
		response, err := ioutil.ReadAll(c)
		handleError(err)
		respString := string(response)
		lines := strings.Split(respString, "\n")
		//amap := make(string[HAProxyAttr])
		for nr, line := range lines {
			fmt.Println(line)
			if nr == 0 || line == "" {
				continue
			}
			hattr := new(HAProxyAttr)
			sp := string.Split(line, ",")
			hattr.Initialize(sp)
			fmt.Println(hattr)
		}
		time.Sleep(3 * time.Second)
	}
}
