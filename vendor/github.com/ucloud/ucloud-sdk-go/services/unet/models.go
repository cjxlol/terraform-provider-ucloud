// Code is generated by ucloud-model, DO NOT EDIT IT.

package unet

/*
UnetEIPAddrSet - DescribeEIP
*/
type UnetEIPAddrSet struct {

	// IP地址
	IP string

	// 运营商信息如: 电信: Telecom, 联通: Unicom, 国际: International, Duplet: 双线IP（电信+联通), BGP: Bgp
	OperatorName string
}

/*
UnetAllocateEIPSet - AllocateEIP
*/
type UnetAllocateEIPSet struct {

	// 申请到的IPv4地址.
	EIPAddr []UnetEIPAddrSet

	// 申请到的EIP资源ID
	EIPId string
}

/*
EIPAddrSet - DescribeBandwidthPackage
*/
type EIPAddrSet struct {

	// 弹性IP地址
	IP string

	// 运营商信息, 枚举值为:  BGP: BGP; International: 国际.
	OperatorName string
}

/*
UnetBandwidthPackageSet - DescribeBandwidthPackage
*/
type UnetBandwidthPackageSet struct {

	// 带宽包的临时带宽值, 单位Mbps
	Bandwidth int

	// 带宽包的资源ID
	BandwidthPackageId string

	// 创建时间, 格式为 Unix Timestamp
	CreateTime int

	// 失效时间, 格式为 Unix Timestamp
	DisableTime int

	// 带宽包所绑定弹性IP的详细信息,只有当EIPId对应双线IP时, EIPAddr的长度为2, 其他情况, EIPAddr长度均为1.参见 EIPAddrSet
	EIPAddr []EIPAddrSet

	// 带宽包所绑定弹性IP的资源ID
	EIPId string

	// 生效时间, 格式为 Unix Timestamp
	EnableTime int
}

/*
UnetBandwidthUsageEIPSet - DescribeBandwidthUsage
*/
type UnetBandwidthUsageEIPSet struct {

	// 最近5分钟带宽用量, 单位Mbps
	CurBandwidth float64

	// 弹性IP资源ID
	EIPId string
}

/*
ShareBandwidthSet - DescribeEIP
*/
type ShareBandwidthSet struct {

	// 共享带宽带宽值
	ShareBandwidth int

	// 共享带宽ID
	ShareBandwidthId string

	// 共享带宽的资源名称
	ShareBandwidthName string
}

/*
UnetEIPResourceSet - DescribeEIP
*/
type UnetEIPResourceSet struct {

	// 弹性IP的资源ID
	EIPId string

	// 已绑定资源的资源ID
	ResourceID string

	//
	ResourceId string `deprecated:"true"`

	// 已绑定的资源名称
	ResourceName string

	// 已绑定的资源类型, 枚举值为: uhost, 云主机；natgw：NAT网关；ulb：负载均衡器；upm: 物理机; hadoophost: 大数据集群;fortresshost：堡垒机；udockhost：容器；udhost：私有专区主机；vpngw：IPSec VPN；ucdr：云灾备；dbaudit：数据库审计，uni：虚拟网卡。
	ResourceType string

	// 资源绑定的虚拟网卡的ID
	SubResourceId string

	// 资源绑定的虚拟网卡的名称
	SubResourceName string

	// 资源绑定的虚拟网卡的类型。uni，虚拟网卡。
	SubResourceType string
}

/*
UnetEIPSet - DescribeEIP
*/
type UnetEIPSet struct {

	// 弹性IP的带宽, 单位为Mbps, 当BandwidthType=1时, 该处显示为共享带宽值. 当BandwidthType=0时, 该处显示这个弹性IP的带宽.
	Bandwidth int

	// 带宽模式, 枚举值为: 0: 非共享带宽模式, 1: 共享带宽模式
	BandwidthType int

	// 付费方式, 枚举值为: Year, 按年付费; Month, 按月付费; Dynamic, 按小时付费; Trial, 试用. 按小时付费和试用这两种付费模式需要开通权限.
	ChargeType string

	// 弹性IP的创建时间, 格式为Unix Timestamp
	CreateTime int

	// 弹性IP的详细信息列表, 具体结构见下方 UnetEIPAddrSet
	EIPAddr []UnetEIPAddrSet

	// 弹性IP的资源ID
	EIPId string

	// 弹性IP是否到期
	Expire bool

	// 弹性IP的到期时间, 格式为Unix Timestamp
	ExpireTime int

	// 弹性IP的名称,缺省值为 "EIP"
	Name string

	// 弹性IP的计费模式, 枚举值为: "Bandwidth", 带宽计费; "Traffic", 流量计费; "ShareBandwidth",共享带宽模式. 默认为 "Bandwidth".
	PayMode string

	// 弹性IP的备注, 缺省值为 ""
	Remark string

	// 弹性IP的详细信息列表, 具体结构见下方 UnetEIPResourceSet
	Resource UnetEIPResourceSet

	// 共享带宽信息 参见 ShareBandwidthSet
	ShareBandwidthSet ShareBandwidthSet

	// 弹性IP的资源绑定状态, 枚举值为: used: 已绑定, free: 未绑定, freeze: 已冻结
	Status string

	// 弹性IP的业务组标识, 缺省值为 "Default"
	Tag string

	// 外网出口权重, 默认为50, 范围[0-100]
	Weight int
}

/*
FirewallRuleSet - DescribeFirewall
*/
type FirewallRuleSet struct {

	// 目标端口
	DstPort string

	// 优先级
	Priority string

	// 协议类型
	ProtocolType string

	// 防火墙规则备注
	Remark string

	// 防火墙动作
	RuleAction string

	// 源地址
	SrcIP string
}

/*
FirewallDataSet - DescribeFirewall
*/
type FirewallDataSet struct {

	// 防火墙组创建时间，格式为Unix Timestamp
	CreateTime int

	// 防火墙ID
	FWId string

	// 安全组ID（即将废弃）
	GroupId string

	// 防火墙名称
	Name string

	// 防火墙备注
	Remark string

	// 防火墙绑定资源数量
	ResourceCount int

	// 防火墙组中的规则列表，参见 FirewallRuleSet
	Rule []FirewallRuleSet

	// 防火墙业务组
	Tag string

	// 防火墙组类型，枚举值为： "user defined", 用户自定义防火墙； "recommend web", 默认Web防火墙； "recommend non web", 默认非Web防火墙
	Type string
}

/*
ResourceSet - 资源信息
*/
type ResourceSet struct {

	// 名称
	Name string

	// 内网IP
	PrivateIP string

	// 备注
	Remark string

	// 绑定该防火墙的资源id
	ResourceID string

	// 绑定防火墙组的资源类型。"unatgw"，NAT网关； "uhost"，云主机； "upm"，物理云主机； "hadoophost"，hadoop节点； "fortresshost"，堡垒机； "udhost"，私有专区主机；"udockhost"，容器；"dbaudit"，数据库审计，“uni”，虚拟网卡。
	ResourceType string

	// 状态
	Status int

	// 资源绑定的虚拟网卡的ID
	SubResourceId string

	// 资源绑定的虚拟网卡的名称
	SubResourceName string

	// 资源绑定的虚拟网卡的类型，“uni”，虚拟网卡。
	SubResourceType string

	// 业务组
	Tag string

	// 可用区
	Zone int
}

/*
EIPSetData - describeShareBandwidth
*/
type EIPSetData struct {

	// EIP带宽值
	Bandwidth int

	// EIP的IP信息，详情见EIPAddrSet
	EIPAddr []EIPAddrSet

	// EIP资源Id
	EIPId string
}

/*
UnetShareBandwidthSet - DescribeShareBandwidth
*/
type UnetShareBandwidthSet struct {

	//
	BandwidthGuarantee int `deprecated:"true"`

	// 付费方式, 预付费:Year 按年,Month 按月,Dynamic 按需;后付费:PostPay(按月)
	ChargeType string

	// 创建时间, 格式为Unix Timestamp
	CreateTime int

	// EIP信息,详情见 EIPSetData
	EIPSet []EIPSetData

	// 过期时间, 格式为Unix Timestamp
	ExpireTime int

	// 共享带宽类型
	IPVersion string

	// 共享带宽名称
	Name string

	//
	PostPayStartTime int `deprecated:"true"`

	// 共享带宽值(预付费)/共享带宽峰值(后付费), 单位Mbps
	ShareBandwidth int

	// 共享带宽的资源ID
	ShareBandwidthId string
}

/*
EIPPayModeSet - GetEIPPayModeEIP
*/
type EIPPayModeSet struct {

	// EIP的资源ID
	EIPId string

	// EIP的计费模式. 枚举值为：Bandwidth, 带宽计费;Traffic, 流量计费; "ShareBandwidth",共享带宽模式
	EIPPayMode string
}

/*
EIPPriceDetailSet - GetEIPPrice
*/
type EIPPriceDetailSet struct {

	// 弹性IP付费方式
	ChargeType string

	// 弹性IP的原价，单位“元”
	OriginalPrice float64

	// 购买弹性IP的实际价格, 单位"元"
	Price float64

	// 资源有效期, 以Unix Timestamp表示
	PurchaseValue int
}

/*
ThroughputDailyBillingInfo - 流量计费EIP计费信息
*/
type ThroughputDailyBillingInfo struct {

	// 是否已计费，“Yes”或者“No”
	BillingState string

	// 计费结束时间
	EndTime int

	// 计费流量，单位“GB”
	QuantityOut int

	// 计费开始时间
	StartTime int
}
