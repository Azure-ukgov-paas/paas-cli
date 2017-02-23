// This file was generated by counterfeiter
package sshfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/cli/cf/ssh"
	"golang.org/x/crypto/ssh"
)

type FakeSecureSession struct {
	RequestPtyStub        func(term string, height, width int, termModes ssh.TerminalModes) error
	requestPtyMutex       sync.RWMutex
	requestPtyArgsForCall []struct {
		term      string
		height    int
		width     int
		termModes ssh.TerminalModes
	}
	requestPtyReturns struct {
		result1 error
	}
	SendRequestStub        func(name string, wantReply bool, payload []byte) (bool, error)
	sendRequestMutex       sync.RWMutex
	sendRequestArgsForCall []struct {
		name      string
		wantReply bool
		payload   []byte
	}
	sendRequestReturns struct {
		result1 bool
		result2 error
	}
	StdinPipeStub        func() (io.WriteCloser, error)
	stdinPipeMutex       sync.RWMutex
	stdinPipeArgsForCall []struct{}
	stdinPipeReturns     struct {
		result1 io.WriteCloser
		result2 error
	}
	StdoutPipeStub        func() (io.Reader, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct{}
	stdoutPipeReturns     struct {
		result1 io.Reader
		result2 error
	}
	StderrPipeStub        func() (io.Reader, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct{}
	stderrPipeReturns     struct {
		result1 io.Reader
		result2 error
	}
	StartStub        func(command string) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		command string
	}
	startReturns struct {
		result1 error
	}
	ShellStub        func() error
	shellMutex       sync.RWMutex
	shellArgsForCall []struct{}
	shellReturns     struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecureSession) RequestPty(term string, height int, width int, termModes ssh.TerminalModes) error {
	fake.requestPtyMutex.Lock()
	fake.requestPtyArgsForCall = append(fake.requestPtyArgsForCall, struct {
		term      string
		height    int
		width     int
		termModes ssh.TerminalModes
	}{term, height, width, termModes})
	fake.recordInvocation("RequestPty", []interface{}{term, height, width, termModes})
	fake.requestPtyMutex.Unlock()
	if fake.RequestPtyStub != nil {
		return fake.RequestPtyStub(term, height, width, termModes)
	} else {
		return fake.requestPtyReturns.result1
	}
}

func (fake *FakeSecureSession) RequestPtyCallCount() int {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	return len(fake.requestPtyArgsForCall)
}

func (fake *FakeSecureSession) RequestPtyArgsForCall(i int) (string, int, int, ssh.TerminalModes) {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	return fake.requestPtyArgsForCall[i].term, fake.requestPtyArgsForCall[i].height, fake.requestPtyArgsForCall[i].width, fake.requestPtyArgsForCall[i].termModes
}

func (fake *FakeSecureSession) RequestPtyReturns(result1 error) {
	fake.RequestPtyStub = nil
	fake.requestPtyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) SendRequest(name string, wantReply bool, payload []byte) (bool, error) {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.sendRequestMutex.Lock()
	fake.sendRequestArgsForCall = append(fake.sendRequestArgsForCall, struct {
		name      string
		wantReply bool
		payload   []byte
	}{name, wantReply, payloadCopy})
	fake.recordInvocation("SendRequest", []interface{}{name, wantReply, payloadCopy})
	fake.sendRequestMutex.Unlock()
	if fake.SendRequestStub != nil {
		return fake.SendRequestStub(name, wantReply, payload)
	} else {
		return fake.sendRequestReturns.result1, fake.sendRequestReturns.result2
	}
}

func (fake *FakeSecureSession) SendRequestCallCount() int {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return len(fake.sendRequestArgsForCall)
}

func (fake *FakeSecureSession) SendRequestArgsForCall(i int) (string, bool, []byte) {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return fake.sendRequestArgsForCall[i].name, fake.sendRequestArgsForCall[i].wantReply, fake.sendRequestArgsForCall[i].payload
}

func (fake *FakeSecureSession) SendRequestReturns(result1 bool, result2 error) {
	fake.SendRequestStub = nil
	fake.sendRequestReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdinPipe() (io.WriteCloser, error) {
	fake.stdinPipeMutex.Lock()
	fake.stdinPipeArgsForCall = append(fake.stdinPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdinPipe", []interface{}{})
	fake.stdinPipeMutex.Unlock()
	if fake.StdinPipeStub != nil {
		return fake.StdinPipeStub()
	} else {
		return fake.stdinPipeReturns.result1, fake.stdinPipeReturns.result2
	}
}

func (fake *FakeSecureSession) StdinPipeCallCount() int {
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	return len(fake.stdinPipeArgsForCall)
}

func (fake *FakeSecureSession) StdinPipeReturns(result1 io.WriteCloser, result2 error) {
	fake.StdinPipeStub = nil
	fake.stdinPipeReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdoutPipe() (io.Reader, error) {
	fake.stdoutPipeMutex.Lock()
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdoutPipe", []interface{}{})
	fake.stdoutPipeMutex.Unlock()
	if fake.StdoutPipeStub != nil {
		return fake.StdoutPipeStub()
	} else {
		return fake.stdoutPipeReturns.result1, fake.stdoutPipeReturns.result2
	}
}

func (fake *FakeSecureSession) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

func (fake *FakeSecureSession) StdoutPipeReturns(result1 io.Reader, result2 error) {
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StderrPipe() (io.Reader, error) {
	fake.stderrPipeMutex.Lock()
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct{}{})
	fake.recordInvocation("StderrPipe", []interface{}{})
	fake.stderrPipeMutex.Unlock()
	if fake.StderrPipeStub != nil {
		return fake.StderrPipeStub()
	} else {
		return fake.stderrPipeReturns.result1, fake.stderrPipeReturns.result2
	}
}

func (fake *FakeSecureSession) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

func (fake *FakeSecureSession) StderrPipeReturns(result1 io.Reader, result2 error) {
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) Start(command string) error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		command string
	}{command})
	fake.recordInvocation("Start", []interface{}{command})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(command)
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeSecureSession) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeSecureSession) StartArgsForCall(i int) string {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].command
}

func (fake *FakeSecureSession) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) Shell() error {
	fake.shellMutex.Lock()
	fake.shellArgsForCall = append(fake.shellArgsForCall, struct{}{})
	fake.recordInvocation("Shell", []interface{}{})
	fake.shellMutex.Unlock()
	if fake.ShellStub != nil {
		return fake.ShellStub()
	} else {
		return fake.shellReturns.result1
	}
}

func (fake *FakeSecureSession) ShellCallCount() int {
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	return len(fake.shellArgsForCall)
}

func (fake *FakeSecureSession) ShellReturns(result1 error) {
	fake.ShellStub = nil
	fake.shellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeSecureSession) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureSession) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeSecureSession) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureSession) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSecureSession) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ sshCmd.SecureSession = new(FakeSecureSession)