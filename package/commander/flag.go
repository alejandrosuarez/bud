package commander

type Flag struct {
	name  string
	usage string
	value value
	short byte
}

func (f *Flag) Short(short byte) *Flag {
	f.short = short
	return f
}

func (f *Flag) Int(target *int) *Int {
	value := &Int{target: target}
	f.value = &intValue{inner: value}
	return value
}

func (f *Flag) String(target *string) *String {
	value := &String{target: target}
	f.value = &stringValue{inner: value}
	return value
}

func (f *Flag) Strings(target *[]string) *Strings {
	value := &Strings{target: target}
	f.value = &stringsValue{inner: value}
	return value
}

func (f *Flag) StringMap(target *map[string]string) *StringMap {
	value := &StringMap{target: target}
	f.value = &stringMapValue{inner: value}
	return value
}

func (f *Flag) Bool(target *bool) *Bool {
	value := &Bool{target: target}
	f.value = &boolValue{inner: value}
	return value
}

func (f *Flag) verify(name string) error {
	return f.value.verify("--" + name)
}

func verifyFlags(flags []*Flag) error {
	for _, flag := range flags {
		if err := flag.verify(flag.name); err != nil {
			return err
		}
	}
	return nil
}
