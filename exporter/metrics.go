package exporter

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "daikin_aircon"

	htemp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "htemp"),
		"sensor info, htemp",
		[]string{"actual", "target"}, nil,
	)

	hhum = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "hhum"),
		"sensor info, hhum",
		[]string{"actual", "target"}, nil,
	)

	otemp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "otemp"),
		"sensor info, otemp",
		[]string{"actual", "target"}, nil,
	)

	er = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "err"),
		"sensor info, err",
		[]string{"actual", "target"}, nil,
	)

	cmpfreq = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cmpfreq"),
		"sensor info, cmpfreq",
		[]string{"actual", "target"}, nil,
	)

	mompow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "mompow"),
		"sensor info, mompow",
		[]string{"actual", "target"}, nil,
	)

	filter_sign = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "filter_sign"),
		"sensor info, filter_sign",
		[]string{"actual", "target"}, nil,
	)

	pow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pow"),
		"control info, pow",
		[]string{"actual", "target"}, nil,
	)

	mode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "mode"),
		"control info, mode",
		[]string{"actual", "target"}, nil,
	)

	adv = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "adv"),
		"control info, adv",
		[]string{"actual", "target"}, nil,
	)

	stemp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "stemp"),
		"control info, stemp",
		[]string{"actual", "target"}, nil,
	)

	shum = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "shum"),
		"control info, shum",
		[]string{"actual", "target"}, nil,
	)

	dt1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt1"),
		"control info, dt1",
		[]string{"actual", "target"}, nil,
	)

	dt2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt2"),
		"control info, dt2",
		[]string{"actual", "target"}, nil,
	)

	dt3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt3"),
		"control info, dt3",
		[]string{"actual", "target"}, nil,
	)

	dt4 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt4"),
		"control info, dt4",
		[]string{"actual", "target"}, nil,
	)

	dt5 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt5"),
		"control info, dt5",
		[]string{"actual", "target"}, nil,
	)

	dt7 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt7"),
		"control info, dt7",
		[]string{"actual", "target"}, nil,
	)

	dh1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh1"),
		"control info, dh1",
		[]string{"actual", "target"}, nil,
	)

	dh2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh2"),
		"control info, dh2",
		[]string{"actual", "target"}, nil,
	)

	dh3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh3"),
		"control info, dh3",
		[]string{"actual", "target"}, nil,
	)

	dh4 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh4"),
		"control info, dh4",
		[]string{"actual", "target"}, nil,
	)

	dh5 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh5"),
		"control info, dh5",
		[]string{"actual", "target"}, nil,
	)

	dh7 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dh7"),
		"control info, dh7",
		[]string{"actual", "target"}, nil,
	)

	dhh = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dhh"),
		"control info, dhh",
		[]string{"actual", "target"}, nil,
	)

	b_mode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_mode"),
		"control info, b_mode",
		[]string{"actual", "target"}, nil,
	)

	b_stemp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_stemp"),
		"control info, b_stemp",
		[]string{"actual", "target"}, nil,
	)

	b_shum = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_shum"),
		"control info, b_shum",
		[]string{"actual", "target"}, nil,
	)

	alert = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "alert"),
		"control info, alert",
		[]string{"actual", "target"}, nil,
	)

	f_rate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "f_rate"),
		"control info, f_rate",
		[]string{"actual", "target"}, nil,
	)

	b_f_rate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_f_rate"),
		"control info, b_f_rate",
		[]string{"actual", "target"}, nil,
	)

	dfr1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr1"),
		"control info, dfr1",
		[]string{"actual", "target"}, nil,
	)

	dfr2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr2"),
		"control info, dfr2",
		[]string{"actual", "target"}, nil,
	)

	dfr3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr3"),
		"control info, dfr3",
		[]string{"actual", "target"}, nil,
	)

	dfr4 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr4"),
		"control info, dfr4",
		[]string{"actual", "target"}, nil,
	)

	dfr5 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr5"),
		"control info, dfr5",
		[]string{"actual", "target"}, nil,
	)

	dfr6 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr6"),
		"control info, dfr6",
		[]string{"actual", "target"}, nil,
	)

	dfr7 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfr7"),
		"control info, dfr7",
		[]string{"actual", "target"}, nil,
	)

	dfrh = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfrh"),
		"control info, dfrh",
		[]string{"actual", "target"}, nil,
	)

	f_dir = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "f_dir"),
		"control info, f_dir",
		[]string{"actual", "target"}, nil,
	)

	b_f_dir = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_f_dir"),
		"control info, b_f_dir",
		[]string{"actual", "target"}, nil,
	)

	dfd1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd1"),
		"control info, dfd1",
		[]string{"actual", "target"}, nil,
	)

	dfd2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd2"),
		"control info, dfd2",
		[]string{"actual", "target"}, nil,
	)

	dfd3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd3"),
		"control info, dfd3",
		[]string{"actual", "target"}, nil,
	)

	dfd4 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd4"),
		"control info, dfd4",
		[]string{"actual", "target"}, nil,
	)

	dfd5 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd5"),
		"control info, dfd5",
		[]string{"actual", "target"}, nil,
	)

	dfd6 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd6"),
		"control info, dfd6",
		[]string{"actual", "target"}, nil,
	)

	dfd7 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfd7"),
		"control info, dfd7",
		[]string{"actual", "target"}, nil,
	)

	dfdh = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dfdh"),
		"control info, dfdh",
		[]string{"actual", "target"}, nil,
	)

	stemp_a = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "stemp_a"),
		"control info, stemp_a",
		[]string{"actual", "target"}, nil,
	)

	dt1_a = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt1_a"),
		"control info, dt1_a",
		[]string{"actual", "target"}, nil,
	)

	dt7_a = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dt7_a"),
		"control info, dt7_a",
		[]string{"actual", "target"}, nil,
	)

	b_stemp_a = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_stemp_a"),
		"control info, b_stemp_a",
		[]string{"actual", "target"}, nil,
	)

	f_dir_ud = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "f_dir_ud"),
		"control info, f_dir_ud",
		[]string{"actual", "target"}, nil,
	)

	f_dir_lr = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "f_dir_lr"),
		"control info, f_dir_lr",
		[]string{"actual", "target"}, nil,
	)

	b_f_dir_ud = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_f_dir_ud"),
		"control info, b_f_dir_ud",
		[]string{"actual", "target"}, nil,
	)

	b_f_dir_lr = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "b_f_dir_lr"),
		"control info, b_f_dir_lr",
		[]string{"actual", "target"}, nil,
	)

	ndfd1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd1"),
		"control info, ndfd1",
		[]string{"actual", "target"}, nil,
	)

	ndfd2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd2"),
		"control info, ndfd2",
		[]string{"actual", "target"}, nil,
	)

	ndfd3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd3"),
		"control info, ndfd3",
		[]string{"actual", "target"}, nil,
	)

	ndfd4 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd4"),
		"control info, ndfd4",
		[]string{"actual", "target"}, nil,
	)

	ndfd5 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd5"),
		"control info, ndfd5",
		[]string{"actual", "target"}, nil,
	)

	ndfd6 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd6"),
		"control info, ndfd6",
		[]string{"actual", "target"}, nil,
	)

	ndfd7 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfd7"),
		"control info, ndfd7",
		[]string{"actual", "target"}, nil,
	)

	ndfdh = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ndfdh"),
		"control info, ndfdh",
		[]string{"actual", "target"}, nil,
	)
)

type Metrics struct {
	Target string
}

func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- htemp
	ch <- hhum
	ch <- otemp
	ch <- er
	ch <- cmpfreq
	ch <- mompow
	ch <- filter_sign
	ch <- pow
	ch <- mode
	ch <- adv
	ch <- stemp
	ch <- shum
	ch <- dt1
	ch <- dt2
	ch <- dt3
	ch <- dt4
	ch <- dt5
	ch <- dt7
	ch <- dh1
	ch <- dh2
	ch <- dh3
	ch <- dh4
	ch <- dh5
	ch <- dh7
	ch <- dhh
	ch <- b_mode
	ch <- b_stemp
	ch <- b_shum
	ch <- alert
	ch <- f_rate
	ch <- b_f_rate
	ch <- dfr1
	ch <- dfr2
	ch <- dfr3
	ch <- dfr4
	ch <- dfr5
	ch <- dfr6
	ch <- dfr7
	ch <- dfrh
	ch <- f_dir
	ch <- b_f_dir
	ch <- dfd1
	ch <- dfd2
	ch <- dfd3
	ch <- dfd4
	ch <- dfd5
	ch <- dfd6
	ch <- dfd7
	ch <- dfdh
	ch <- stemp_a
	ch <- dt1_a
	ch <- dt7_a
	ch <- b_stemp_a
	ch <- f_dir_ud
	ch <- f_dir_lr
	ch <- b_f_dir_ud
	ch <- b_f_dir_lr
	ch <- ndfd1
	ch <- ndfd2
	ch <- ndfd3
	ch <- ndfd4
	ch <- ndfd5
	ch <- ndfd6
	ch <- ndfd7
	ch <- ndfdh
}

func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	err := m.getInfo(path.Join("aircon", "get_sensor_info"), ch)
	if err != nil {
		log.Println("Failed to get info from aircon", err)
		return
	}
	err = m.getInfo(path.Join("aircon", "get_control_info"), ch)
	if err != nil {
		log.Println("Failed to get info from aircon", err)
		return
	}
}

func (m Metrics) getInfo(relativePath string, ch chan<- prometheus.Metric) (err error) {
	r, err := http.DefaultClient.Get("http://" + path.Join(m.Target, relativePath))
	if err != nil {
		return
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	for _, v := range strings.Split(string(b), ",") {
		t := strings.Split(v, "=")
		v, err := strconv.ParseFloat(t[1], 32)
		actual := t[1]
		if err != nil {
			continue
		}
		switch t[0] {
		case "htemp":
			ch <- prometheus.MustNewConstMetric(htemp, prometheus.GaugeValue, v, actual, m.Target)
		case "hhum":
			ch <- prometheus.MustNewConstMetric(hhum, prometheus.GaugeValue, v, actual, m.Target)
		case "otemp":
			ch <- prometheus.MustNewConstMetric(otemp, prometheus.GaugeValue, v, actual, m.Target)
		case "err":
			ch <- prometheus.MustNewConstMetric(er, prometheus.GaugeValue, v, actual, m.Target)
		case "cmpfreq":
			ch <- prometheus.MustNewConstMetric(cmpfreq, prometheus.GaugeValue, v, actual, m.Target)
		case "mompow":
			ch <- prometheus.MustNewConstMetric(mompow, prometheus.GaugeValue, v, actual, m.Target)
		case "filter_sign":
			ch <- prometheus.MustNewConstMetric(filter_sign, prometheus.GaugeValue, v, actual, m.Target)
		case "pow":
			ch <- prometheus.MustNewConstMetric(pow, prometheus.GaugeValue, v, actual, m.Target)
		case "mode":
			ch <- prometheus.MustNewConstMetric(mode, prometheus.GaugeValue, v, actual, m.Target)
		case "adv":
			ch <- prometheus.MustNewConstMetric(adv, prometheus.GaugeValue, v, actual, m.Target)
		case "stemp":
			ch <- prometheus.MustNewConstMetric(stemp, prometheus.GaugeValue, v, actual, m.Target)
		case "shum":
			ch <- prometheus.MustNewConstMetric(shum, prometheus.GaugeValue, v, actual, m.Target)
		case "dt1":
			ch <- prometheus.MustNewConstMetric(dt1, prometheus.GaugeValue, v, actual, m.Target)
		case "dt2":
			ch <- prometheus.MustNewConstMetric(dt2, prometheus.GaugeValue, v, actual, m.Target)
		case "dt3":
			ch <- prometheus.MustNewConstMetric(dt3, prometheus.GaugeValue, v, actual, m.Target)
		case "dt4":
			ch <- prometheus.MustNewConstMetric(dt4, prometheus.GaugeValue, v, actual, m.Target)
		case "dt5":
			ch <- prometheus.MustNewConstMetric(dt5, prometheus.GaugeValue, v, actual, m.Target)
		case "dt7":
			ch <- prometheus.MustNewConstMetric(dt7, prometheus.GaugeValue, v, actual, m.Target)
		case "dh1":
			ch <- prometheus.MustNewConstMetric(dh1, prometheus.GaugeValue, v, actual, m.Target)
		case "dh2":
			ch <- prometheus.MustNewConstMetric(dh2, prometheus.GaugeValue, v, actual, m.Target)
		case "dh3":
			ch <- prometheus.MustNewConstMetric(dh3, prometheus.GaugeValue, v, actual, m.Target)
		case "dh4":
			ch <- prometheus.MustNewConstMetric(dh4, prometheus.GaugeValue, v, actual, m.Target)
		case "dh5":
			ch <- prometheus.MustNewConstMetric(dh5, prometheus.GaugeValue, v, actual, m.Target)
		case "dh7":
			ch <- prometheus.MustNewConstMetric(dh7, prometheus.GaugeValue, v, actual, m.Target)
		case "dhh":
			ch <- prometheus.MustNewConstMetric(dhh, prometheus.GaugeValue, v, actual, m.Target)
		case "b_mode":
			ch <- prometheus.MustNewConstMetric(b_mode, prometheus.GaugeValue, v, actual, m.Target)
		case "b_stemp":
			ch <- prometheus.MustNewConstMetric(b_stemp, prometheus.GaugeValue, v, actual, m.Target)
		case "b_shum":
			ch <- prometheus.MustNewConstMetric(b_shum, prometheus.GaugeValue, v, actual, m.Target)
		case "alert":
			ch <- prometheus.MustNewConstMetric(alert, prometheus.GaugeValue, v, actual, m.Target)
		case "f_rate":
			ch <- prometheus.MustNewConstMetric(f_rate, prometheus.GaugeValue, v, actual, m.Target)
		case "b_f_rate":
			ch <- prometheus.MustNewConstMetric(b_f_rate, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr1":
			ch <- prometheus.MustNewConstMetric(dfr1, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr2":
			ch <- prometheus.MustNewConstMetric(dfr2, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr3":
			ch <- prometheus.MustNewConstMetric(dfr3, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr4":
			ch <- prometheus.MustNewConstMetric(dfr4, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr5":
			ch <- prometheus.MustNewConstMetric(dfr5, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr6":
			ch <- prometheus.MustNewConstMetric(dfr6, prometheus.GaugeValue, v, actual, m.Target)
		case "dfr7":
			ch <- prometheus.MustNewConstMetric(dfr7, prometheus.GaugeValue, v, actual, m.Target)
		case "dfrh":
			ch <- prometheus.MustNewConstMetric(dfrh, prometheus.GaugeValue, v, actual, m.Target)
		case "f_dir":
			ch <- prometheus.MustNewConstMetric(f_dir, prometheus.GaugeValue, v, actual, m.Target)
		case "b_f_dir":
			ch <- prometheus.MustNewConstMetric(b_f_dir, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd1":
			ch <- prometheus.MustNewConstMetric(dfd1, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd2":
			ch <- prometheus.MustNewConstMetric(dfd2, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd3":
			ch <- prometheus.MustNewConstMetric(dfd3, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd4":
			ch <- prometheus.MustNewConstMetric(dfd4, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd5":
			ch <- prometheus.MustNewConstMetric(dfd5, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd6":
			ch <- prometheus.MustNewConstMetric(dfd6, prometheus.GaugeValue, v, actual, m.Target)
		case "dfd7":
			ch <- prometheus.MustNewConstMetric(dfd7, prometheus.GaugeValue, v, actual, m.Target)
		case "dfdh":
			ch <- prometheus.MustNewConstMetric(dfdh, prometheus.GaugeValue, v, actual, m.Target)
		case "stemp_a":
			ch <- prometheus.MustNewConstMetric(stemp_a, prometheus.GaugeValue, v, actual, m.Target)
		case "dt1_a":
			ch <- prometheus.MustNewConstMetric(dt1_a, prometheus.GaugeValue, v, actual, m.Target)
		case "dt7_a":
			ch <- prometheus.MustNewConstMetric(dt7_a, prometheus.GaugeValue, v, actual, m.Target)
		case "b_stemp_a":
			ch <- prometheus.MustNewConstMetric(b_stemp_a, prometheus.GaugeValue, v, actual, m.Target)
		case "f_dir_ud":
			ch <- prometheus.MustNewConstMetric(f_dir_ud, prometheus.GaugeValue, v, actual, m.Target)
		case "f_dir_lr":
			ch <- prometheus.MustNewConstMetric(f_dir_lr, prometheus.GaugeValue, v, actual, m.Target)
		case "b_f_dir_ud":
			ch <- prometheus.MustNewConstMetric(b_f_dir_ud, prometheus.GaugeValue, v, actual, m.Target)
		case "b_f_dir_lr":
			ch <- prometheus.MustNewConstMetric(b_f_dir_lr, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd1":
			ch <- prometheus.MustNewConstMetric(ndfd1, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd2":
			ch <- prometheus.MustNewConstMetric(ndfd2, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd3":
			ch <- prometheus.MustNewConstMetric(ndfd3, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd4":
			ch <- prometheus.MustNewConstMetric(ndfd4, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd5":
			ch <- prometheus.MustNewConstMetric(ndfd5, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd6":
			ch <- prometheus.MustNewConstMetric(ndfd6, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfd7":
			ch <- prometheus.MustNewConstMetric(ndfd7, prometheus.GaugeValue, v, actual, m.Target)
		case "ndfdh":
			ch <- prometheus.MustNewConstMetric(ndfdh, prometheus.GaugeValue, v, actual, m.Target)
		}
	}
	return
}
