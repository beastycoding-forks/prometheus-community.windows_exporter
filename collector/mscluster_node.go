package collector

import (
	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("mscluster_node", newMSCluster_NodeCollector)
}

// A MSCluster_NodeCollector is a Prometheus collector for WMI MSCluster_Node metrics
type MSCluster_NodeCollector struct {
	BuildNumber           *prometheus.Desc
	Characteristics       *prometheus.Desc
	DetectedCloudPlatform *prometheus.Desc
	DynamicWeight         *prometheus.Desc
	Flags                 *prometheus.Desc
	MajorVersion          *prometheus.Desc
	MinorVersion          *prometheus.Desc
	NeedsPreventQuorum    *prometheus.Desc
	NodeDrainStatus       *prometheus.Desc
	NodeHighestVersion    *prometheus.Desc
	NodeLowestVersion     *prometheus.Desc
	NodeWeight            *prometheus.Desc
	State                 *prometheus.Desc
	StatusInformation     *prometheus.Desc
}

func newMSCluster_NodeCollector() (Collector, error) {
	const subsystem = "mscluster_node"
	return &MSCluster_NodeCollector{
		BuildNumber: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "build_number"),
			"(BuildNumber)",
			[]string{"name"},
			nil,
		),
		Characteristics: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "characteristics"),
			"(Characteristics)",
			[]string{"name"},
			nil,
		),
		DetectedCloudPlatform: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "detected_cloud_platform"),
			"(DetectedCloudPlatform)",
			[]string{"name"},
			nil,
		),
		DynamicWeight: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "dynamic_weight"),
			"(DynamicWeight)",
			[]string{"name"},
			nil,
		),
		Flags: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "flags"),
			"(Flags)",
			[]string{"name"},
			nil,
		),
		MajorVersion: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "major_version"),
			"(MajorVersion)",
			[]string{"name"},
			nil,
		),
		MinorVersion: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "minor_version"),
			"(MinorVersion)",
			[]string{"name"},
			nil,
		),
		NeedsPreventQuorum: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "needs_prevent_quorum"),
			"(NeedsPreventQuorum)",
			[]string{"name"},
			nil,
		),
		NodeDrainStatus: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "node_drain_status"),
			"(NodeDrainStatus)",
			[]string{"name"},
			nil,
		),
		NodeHighestVersion: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "node_highest_version"),
			"(NodeHighestVersion)",
			[]string{"name"},
			nil,
		),
		NodeLowestVersion: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "node_lowest_version"),
			"(NodeLowestVersion)",
			[]string{"name"},
			nil,
		),
		NodeWeight: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "node_weight"),
			"(NodeWeight)",
			[]string{"name"},
			nil,
		),
		State: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "state"),
			"(State)",
			[]string{"name"},
			nil,
		),
		StatusInformation: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "status_information"),
			"(StatusInformation)",
			[]string{"name"},
			nil,
		),
	}, nil
}

// MSCluster_Node docs:
// - <add link to documentation here>
type MSCluster_Node struct {
	Name string

	BuildNumber           uint
	Characteristics       uint
	DetectedCloudPlatform uint
	DynamicWeight         uint
	Flags                 uint
	MajorVersion          uint
	MinorVersion          uint
	NeedsPreventQuorum    uint
	NodeDrainStatus       uint
	NodeHighestVersion    uint
	NodeLowestVersion     uint
	NodeWeight            uint
	State                 uint
	StatusInformation     uint
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *MSCluster_NodeCollector) Collect(ctx *ScrapeContext, ch chan<- prometheus.Metric) error {
	var dst []MSCluster_Node
	q := queryAll(&dst)
	if err := wmi.QueryNamespace(q, &dst, "root/MSCluster"); err != nil {
		return err
	}

	for _, v := range dst {

		ch <- prometheus.MustNewConstMetric(
			c.BuildNumber,
			prometheus.GaugeValue,
			float64(v.BuildNumber),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Characteristics,
			prometheus.GaugeValue,
			float64(v.Characteristics),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.DetectedCloudPlatform,
			prometheus.GaugeValue,
			float64(v.DetectedCloudPlatform),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.DynamicWeight,
			prometheus.GaugeValue,
			float64(v.DynamicWeight),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Flags,
			prometheus.GaugeValue,
			float64(v.Flags),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.MajorVersion,
			prometheus.GaugeValue,
			float64(v.MajorVersion),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.MinorVersion,
			prometheus.GaugeValue,
			float64(v.MinorVersion),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.NeedsPreventQuorum,
			prometheus.GaugeValue,
			float64(v.NeedsPreventQuorum),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.NodeDrainStatus,
			prometheus.GaugeValue,
			float64(v.NodeDrainStatus),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.NodeHighestVersion,
			prometheus.GaugeValue,
			float64(v.NodeHighestVersion),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.NodeLowestVersion,
			prometheus.GaugeValue,
			float64(v.NodeLowestVersion),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.NodeWeight,
			prometheus.GaugeValue,
			float64(v.NodeWeight),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.State,
			prometheus.GaugeValue,
			float64(v.State),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.StatusInformation,
			prometheus.GaugeValue,
			float64(v.StatusInformation),
			v.Name,
		)
	}

	return nil
}
