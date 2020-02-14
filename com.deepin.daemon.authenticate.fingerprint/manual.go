package fingerprint

func (v *device) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("deviceIfcName")
	ifcNameStr, ok := ifcName.(string)
	if ok {
		return ifcNameStr
	}
	return ""
}

func (v *device) SetInterfaceName(name string)  {
	v.GetObject_().SetExtra("deviceIfcName", name)
}

