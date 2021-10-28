package gojson

type rawProperty struct {
	Name  string
	Bytes []byte
}

type rawProperties []rawProperty

func (rp *rawProperties) Names() []string {
	names := []string{}
	for _, p := range *rp {
		names = append(names, p.Name)
	}
	return names
}

func (rp *rawProperties) Add(oProp rawProperty) {
	self := *rp
	found := false
	for j, sProp := range self {
		if oProp.Name == sProp.Name {
			self[j] = oProp
			found = true
			break
		}
	}
	if !found {
		self = append(self, oProp)
	}
	*rp = self
}

func (rp *rawProperties) Merge(other rawProperties) {
	if (other == nil) || (len(other) == 0) {
		return
	}
	self := *rp
	for _, oProp := range other {
		found := false
		for j, sProp := range self {
			if oProp.Name == sProp.Name {
				self[j] = oProp
				found = true
				break
			}
		}
		if !found {
			self = append(self, oProp)
		}
	}
	*rp = self
}
