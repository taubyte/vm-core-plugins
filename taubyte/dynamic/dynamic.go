package dynamic

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/i2mv/memview"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (f *Factory) getModuleInstanceId() uint32 {
	return f.moduleInstanceIds.Add(1)
}

func (f *Factory) getModuleInstance(id uint32) (vm.ModuleInstance, errno.Error) {
	f.moduleInstanceLock.RLock()
	defer f.moduleInstanceLock.RUnlock()

	if modInstance, exists := f.moduleInstances[id]; exists {
		return modInstance, 0
	}

	return nil, errno.ErrorDynamicGetModuleInstanceFailed
}

func (f *Factory) getFunctionInstanceId() uint32 {
	return f.functionInstanceIds.Add(1)
}

func (f *Factory) getFunctionInstance(id uint32) (vm.FunctionInstance, errno.Error) {
	f.functionInstanceLock.RLock()
	defer f.functionInstanceLock.RUnlock()

	if funcInstance, exists := f.functionInstances[id]; exists {
		return funcInstance, 0
	}

	return nil, errno.ErrorDynamicGetFunctionInstanceFailed
}

func (f *Factory) W_dynamicModule(
	ctx context.Context,
	module vm.Module,
	moduleNamePtr,
	moduleNameLen,
	moduleIdPtr uint32,
) errno.Error {
	moduleName, err0 := f.ReadString(module, moduleNamePtr, moduleNameLen)
	if err0 != 0 {
		return err0
	}

	moduleInstance, err := f.instance.Module(moduleName)
	if err != nil {
		return errno.ErrorDynamicModuleInstanceFailed
	}

	id := f.getModuleInstanceId()
	f.moduleInstanceLock.Lock()
	f.moduleInstances[id] = moduleInstance
	f.moduleInstanceLock.Unlock()

	return f.WriteUint32Le(module, moduleIdPtr, id)
}

func (f *Factory) W_dynamicFunction(
	ctx context.Context,
	module vm.Module,
	moduleId,
	funcNamePtr,
	funcNameLen,
	funcIdPtr uint32,
) errno.Error {
	modInstance, err0 := f.getModuleInstance(moduleId)
	if err0 != 0 {
		return err0
	}

	funcName, err0 := f.ReadString(module, funcNamePtr, funcNameLen)
	if err0 != 0 {
		return err0
	}

	functionInstance, err := modInstance.Function(funcName)
	if err != nil {
		return errno.ErrorDynamicFunctionInstanceFailed
	}

	id := f.getFunctionInstanceId()
	f.functionInstanceLock.Lock()
	f.functionInstances[id] = functionInstance
	f.functionInstanceLock.Unlock()

	return f.WriteUint32Le(module, funcIdPtr, id)
}

func (f *Factory) W_dynamicCall(
	ctx context.Context,
	module vm.Module,
	functionId,
	argsPtr, // []u64 codec
	argsSize,
	outputSize,
	outputMVIdPtr uint32,
) errno.Error {
	functionInstance, err0 := f.getFunctionInstance(functionId)
	if err0 != 0 {
		return err0
	}

	args, err0 := f.ReadUint64Slice(module, argsPtr, argsSize)
	if err0 != 0 {
		return err0
	}

	argInterfaces := inputAsInterfaces(args)
	ret := functionInstance.Call(argInterfaces...)

	if ret.Error() != nil {
		return errno.ErrorDynamicCallFailed
	}

	output := make([]*uint64, outputSize)
	if err := ret.Reflect(argsAsInterfaces(output)...); err != nil {
		return errno.ErrorDynamicCallOutputFailed
	}

	var outputEncoded []byte
	if err := codec.Convert(argsDeRef(output)).To(&outputEncoded); err != nil {
		return errno.ErrorByteConversionFailed
	}

	memViewId, _, err := memview.New(outputEncoded, true)
	if err != nil {
		return errno.ErrorDynamicCallOutputMemviewFailed
	}

	return f.WriteUint32Le(module, outputMVIdPtr, memViewId)
}

func inputAsInterfaces(args []uint64) []interface{} {
	argsInterface := make([]interface{}, 0)

	for _, arg := range args {
		argsInterface = append(argsInterface, arg)
	}

	return argsInterface
}

func argsAsInterfaces(args []*uint64) []interface{} {
	argsInterface := make([]interface{}, 0)

	for _, arg := range args {
		argsInterface = append(argsInterface, arg)
	}

	return argsInterface
}

func argsDeRef(args []*uint64) []uint64 {
	deRef := make([]uint64, 0, len(args))

	for _, arg := range args {
		deRef = append(deRef, *arg)
	}

	return deRef
}
