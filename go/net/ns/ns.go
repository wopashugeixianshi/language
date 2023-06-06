package ns

import (
	"runtime"

	"github.com/vishvananda/netns"
)

type NetNsMoudle struct {
	ns netns.NsHandle
}

func NewNetNsModule() (*NetNsMoudle, error) {
	runtime.LockOSThread()
	defaultNs, err := netns.Get()
	if err != nil {
		return nil, err
	}

	return &NetNsMoudle{defaultNs}, nil
}

func (module *NetNsMoudle) SetNetNs(pid int32) error {
	dockerns, err := netns.GetFromPid(int(pid))
	if err != nil {
		return err
	}
	defer dockerns.Close()

	err = netns.Set(dockerns)
	if err != nil {
		return err
	}

	return nil
}

// close时要恢复原来的ns
func (module *NetNsMoudle) Close() error {
	defer runtime.UnlockOSThread()
	defer module.ns.Close()
	return netns.Set(module.ns)
}


