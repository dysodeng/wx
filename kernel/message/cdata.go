package message

// CDATAText XML文本域
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// value2CDATA 值转换为CDATA
func value2CDATA(value string) CDATAText {
	return CDATAText{"<![CDATA[" + value + "]]>"}
}

// ptrValue2CDATA 值转换为指针型CDATA
func ptrValue2CDATA(value string) *CDATAText {
	return &CDATAText{"<![CDATA[" + value + "]]>"}
}
