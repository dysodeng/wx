package message

// CDATAText XML文本域
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// Value2CDATA 值转换为CDATA
func Value2CDATA(value string) CDATAText {
	return CDATAText{"<![CDATA[" + value + "]]>"}
}

// PtrValue2CDATA 值转换为指针型CDATA
func PtrValue2CDATA(value string) *CDATAText {
	return &CDATAText{"<![CDATA[" + value + "]]>"}
}
