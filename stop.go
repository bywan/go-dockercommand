package dockercommand

type StopOptions struct {
	ID      string
	Timeout uint
}

func (dock *Docker) Stop(options *StopOptions) error {
	return dock.client.StopContainer(options.ID, options.Timeout)
}
