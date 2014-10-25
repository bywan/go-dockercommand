package dockercommand

type Container struct {
	ID         string
	Image      string
	Command    string
	Created    int64
	Status     string
	Ports      []Port
	SizeRw     int64
	SizeRootFs int64
	Names      []string
}

type Port struct {
	PrivatePort int64
	PublicPort  int64
	Type        string
	IP          string
}
