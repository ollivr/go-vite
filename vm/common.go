package vm

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/vm/util"
	"github.com/vitelabs/go-vite/vm_context/vmctxt_interface"
)

const (
	Retry   = true
	NoRetry = false
)

func IsUserAccount(db vmctxt_interface.VmDatabase, addr types.Address) bool {
	if IsPrecompiledContractAddress(addr) {
		return false
	}
	_, code := util.GetContractCode(db, &addr)
	return len(code) == 0
}
