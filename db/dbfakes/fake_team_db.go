// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

type FakeTeamDB struct {
	GetPipelinesStub        func() ([]db.SavedPipeline, error)
	getPipelinesMutex       sync.RWMutex
	getPipelinesArgsForCall []struct{}
	getPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	getPipelinesReturnsOnCall map[int]struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetPublicPipelinesStub        func() ([]db.SavedPipeline, error)
	getPublicPipelinesMutex       sync.RWMutex
	getPublicPipelinesArgsForCall []struct{}
	getPublicPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	getPublicPipelinesReturnsOnCall map[int]struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetPrivateAndAllPublicPipelinesStub        func() ([]db.SavedPipeline, error)
	getPrivateAndAllPublicPipelinesMutex       sync.RWMutex
	getPrivateAndAllPublicPipelinesArgsForCall []struct{}
	getPrivateAndAllPublicPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	getPrivateAndAllPublicPipelinesReturnsOnCall map[int]struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetPipelineByNameStub        func(pipelineName string) (db.SavedPipeline, bool, error)
	getPipelineByNameMutex       sync.RWMutex
	getPipelineByNameArgsForCall []struct {
		pipelineName string
	}
	getPipelineByNameReturns struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}
	getPipelineByNameReturnsOnCall map[int]struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}
	OrderPipelinesStub        func([]string) error
	orderPipelinesMutex       sync.RWMutex
	orderPipelinesArgsForCall []struct {
		arg1 []string
	}
	orderPipelinesReturns struct {
		result1 error
	}
	orderPipelinesReturnsOnCall map[int]struct {
		result1 error
	}
	GetTeamStub        func() (db.SavedTeam, bool, error)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct{}
	getTeamReturns     struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	getTeamReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	UpdateBasicAuthStub        func(basicAuth *db.BasicAuth) (db.SavedTeam, error)
	updateBasicAuthMutex       sync.RWMutex
	updateBasicAuthArgsForCall []struct {
		basicAuth *db.BasicAuth
	}
	updateBasicAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	updateBasicAuthReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateGitHubAuthStub        func(gitHubAuth *db.GitHubAuth) (db.SavedTeam, error)
	updateGitHubAuthMutex       sync.RWMutex
	updateGitHubAuthArgsForCall []struct {
		gitHubAuth *db.GitHubAuth
	}
	updateGitHubAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	updateGitHubAuthReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateUAAAuthStub        func(uaaAuth *db.UAAAuth) (db.SavedTeam, error)
	updateUAAAuthMutex       sync.RWMutex
	updateUAAAuthArgsForCall []struct {
		uaaAuth *db.UAAAuth
	}
	updateUAAAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	updateUAAAuthReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateGenericOAuthStub        func(genericOAuth *db.GenericOAuth) (db.SavedTeam, error)
	updateGenericOAuthMutex       sync.RWMutex
	updateGenericOAuthArgsForCall []struct {
		genericOAuth *db.GenericOAuth
	}
	updateGenericOAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	updateGenericOAuthReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 error
	}
	GetConfigStub        func(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct {
		pipelineName string
	}
	getConfigReturns struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}
	getConfigReturnsOnCall map[int]struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}
	SaveConfigToBeDeprecatedStub        func(string, atc.Config, db.ConfigVersion, db.PipelinePausedState) (db.SavedPipeline, bool, error)
	saveConfigToBeDeprecatedMutex       sync.RWMutex
	saveConfigToBeDeprecatedArgsForCall []struct {
		arg1 string
		arg2 atc.Config
		arg3 db.ConfigVersion
		arg4 db.PipelinePausedState
	}
	saveConfigToBeDeprecatedReturns struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}
	saveConfigToBeDeprecatedReturnsOnCall map[int]struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}
	CreateOneOffBuildStub        func() (db.Build, error)
	createOneOffBuildMutex       sync.RWMutex
	createOneOffBuildArgsForCall []struct{}
	createOneOffBuildReturns     struct {
		result1 db.Build
		result2 error
	}
	createOneOffBuildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 error
	}
	GetPrivateAndPublicBuildsStub        func(page db.Page) ([]db.Build, db.Pagination, error)
	getPrivateAndPublicBuildsMutex       sync.RWMutex
	getPrivateAndPublicBuildsArgsForCall []struct {
		page db.Page
	}
	getPrivateAndPublicBuildsReturns struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	getPrivateAndPublicBuildsReturnsOnCall map[int]struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}
	GetContainerStub        func(handle string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		handle string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	getContainerReturnsOnCall map[int]struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	FindContainersByDescriptorsStub        func(id db.Container) ([]db.SavedContainer, error)
	findContainersByDescriptorsMutex       sync.RWMutex
	findContainersByDescriptorsArgsForCall []struct {
		id db.Container
	}
	findContainersByDescriptorsReturns struct {
		result1 []db.SavedContainer
		result2 error
	}
	findContainersByDescriptorsReturnsOnCall map[int]struct {
		result1 []db.SavedContainer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamDB) GetPipelines() ([]db.SavedPipeline, error) {
	fake.getPipelinesMutex.Lock()
	ret, specificReturn := fake.getPipelinesReturnsOnCall[len(fake.getPipelinesArgsForCall)]
	fake.getPipelinesArgsForCall = append(fake.getPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetPipelines", []interface{}{})
	fake.getPipelinesMutex.Unlock()
	if fake.GetPipelinesStub != nil {
		return fake.GetPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPipelinesReturns.result1, fake.getPipelinesReturns.result2
}

func (fake *FakeTeamDB) GetPipelinesCallCount() int {
	fake.getPipelinesMutex.RLock()
	defer fake.getPipelinesMutex.RUnlock()
	return len(fake.getPipelinesArgsForCall)
}

func (fake *FakeTeamDB) GetPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetPipelinesStub = nil
	fake.getPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPipelinesReturnsOnCall(i int, result1 []db.SavedPipeline, result2 error) {
	fake.GetPipelinesStub = nil
	if fake.getPipelinesReturnsOnCall == nil {
		fake.getPipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.SavedPipeline
			result2 error
		})
	}
	fake.getPipelinesReturnsOnCall[i] = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPublicPipelines() ([]db.SavedPipeline, error) {
	fake.getPublicPipelinesMutex.Lock()
	ret, specificReturn := fake.getPublicPipelinesReturnsOnCall[len(fake.getPublicPipelinesArgsForCall)]
	fake.getPublicPipelinesArgsForCall = append(fake.getPublicPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetPublicPipelines", []interface{}{})
	fake.getPublicPipelinesMutex.Unlock()
	if fake.GetPublicPipelinesStub != nil {
		return fake.GetPublicPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPublicPipelinesReturns.result1, fake.getPublicPipelinesReturns.result2
}

func (fake *FakeTeamDB) GetPublicPipelinesCallCount() int {
	fake.getPublicPipelinesMutex.RLock()
	defer fake.getPublicPipelinesMutex.RUnlock()
	return len(fake.getPublicPipelinesArgsForCall)
}

func (fake *FakeTeamDB) GetPublicPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetPublicPipelinesStub = nil
	fake.getPublicPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPublicPipelinesReturnsOnCall(i int, result1 []db.SavedPipeline, result2 error) {
	fake.GetPublicPipelinesStub = nil
	if fake.getPublicPipelinesReturnsOnCall == nil {
		fake.getPublicPipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.SavedPipeline
			result2 error
		})
	}
	fake.getPublicPipelinesReturnsOnCall[i] = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPrivateAndAllPublicPipelines() ([]db.SavedPipeline, error) {
	fake.getPrivateAndAllPublicPipelinesMutex.Lock()
	ret, specificReturn := fake.getPrivateAndAllPublicPipelinesReturnsOnCall[len(fake.getPrivateAndAllPublicPipelinesArgsForCall)]
	fake.getPrivateAndAllPublicPipelinesArgsForCall = append(fake.getPrivateAndAllPublicPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetPrivateAndAllPublicPipelines", []interface{}{})
	fake.getPrivateAndAllPublicPipelinesMutex.Unlock()
	if fake.GetPrivateAndAllPublicPipelinesStub != nil {
		return fake.GetPrivateAndAllPublicPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPrivateAndAllPublicPipelinesReturns.result1, fake.getPrivateAndAllPublicPipelinesReturns.result2
}

func (fake *FakeTeamDB) GetPrivateAndAllPublicPipelinesCallCount() int {
	fake.getPrivateAndAllPublicPipelinesMutex.RLock()
	defer fake.getPrivateAndAllPublicPipelinesMutex.RUnlock()
	return len(fake.getPrivateAndAllPublicPipelinesArgsForCall)
}

func (fake *FakeTeamDB) GetPrivateAndAllPublicPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetPrivateAndAllPublicPipelinesStub = nil
	fake.getPrivateAndAllPublicPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPrivateAndAllPublicPipelinesReturnsOnCall(i int, result1 []db.SavedPipeline, result2 error) {
	fake.GetPrivateAndAllPublicPipelinesStub = nil
	if fake.getPrivateAndAllPublicPipelinesReturnsOnCall == nil {
		fake.getPrivateAndAllPublicPipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.SavedPipeline
			result2 error
		})
	}
	fake.getPrivateAndAllPublicPipelinesReturnsOnCall[i] = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPipelineByName(pipelineName string) (db.SavedPipeline, bool, error) {
	fake.getPipelineByNameMutex.Lock()
	ret, specificReturn := fake.getPipelineByNameReturnsOnCall[len(fake.getPipelineByNameArgsForCall)]
	fake.getPipelineByNameArgsForCall = append(fake.getPipelineByNameArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.recordInvocation("GetPipelineByName", []interface{}{pipelineName})
	fake.getPipelineByNameMutex.Unlock()
	if fake.GetPipelineByNameStub != nil {
		return fake.GetPipelineByNameStub(pipelineName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getPipelineByNameReturns.result1, fake.getPipelineByNameReturns.result2, fake.getPipelineByNameReturns.result3
}

func (fake *FakeTeamDB) GetPipelineByNameCallCount() int {
	fake.getPipelineByNameMutex.RLock()
	defer fake.getPipelineByNameMutex.RUnlock()
	return len(fake.getPipelineByNameArgsForCall)
}

func (fake *FakeTeamDB) GetPipelineByNameArgsForCall(i int) string {
	fake.getPipelineByNameMutex.RLock()
	defer fake.getPipelineByNameMutex.RUnlock()
	return fake.getPipelineByNameArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetPipelineByNameReturns(result1 db.SavedPipeline, result2 bool, result3 error) {
	fake.GetPipelineByNameStub = nil
	fake.getPipelineByNameReturns = struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetPipelineByNameReturnsOnCall(i int, result1 db.SavedPipeline, result2 bool, result3 error) {
	fake.GetPipelineByNameStub = nil
	if fake.getPipelineByNameReturnsOnCall == nil {
		fake.getPipelineByNameReturnsOnCall = make(map[int]struct {
			result1 db.SavedPipeline
			result2 bool
			result3 error
		})
	}
	fake.getPipelineByNameReturnsOnCall[i] = struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) OrderPipelines(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.orderPipelinesMutex.Lock()
	ret, specificReturn := fake.orderPipelinesReturnsOnCall[len(fake.orderPipelinesArgsForCall)]
	fake.orderPipelinesArgsForCall = append(fake.orderPipelinesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("OrderPipelines", []interface{}{arg1Copy})
	fake.orderPipelinesMutex.Unlock()
	if fake.OrderPipelinesStub != nil {
		return fake.OrderPipelinesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.orderPipelinesReturns.result1
}

func (fake *FakeTeamDB) OrderPipelinesCallCount() int {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return len(fake.orderPipelinesArgsForCall)
}

func (fake *FakeTeamDB) OrderPipelinesArgsForCall(i int) []string {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return fake.orderPipelinesArgsForCall[i].arg1
}

func (fake *FakeTeamDB) OrderPipelinesReturns(result1 error) {
	fake.OrderPipelinesStub = nil
	fake.orderPipelinesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTeamDB) OrderPipelinesReturnsOnCall(i int, result1 error) {
	fake.OrderPipelinesStub = nil
	if fake.orderPipelinesReturnsOnCall == nil {
		fake.orderPipelinesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.orderPipelinesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTeamDB) GetTeam() (db.SavedTeam, bool, error) {
	fake.getTeamMutex.Lock()
	ret, specificReturn := fake.getTeamReturnsOnCall[len(fake.getTeamArgsForCall)]
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct{}{})
	fake.recordInvocation("GetTeam", []interface{}{})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3
}

func (fake *FakeTeamDB) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeTeamDB) GetTeamReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetTeamReturnsOnCall(i int, result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	if fake.getTeamReturnsOnCall == nil {
		fake.getTeamReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 bool
			result3 error
		})
	}
	fake.getTeamReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) UpdateBasicAuth(basicAuth *db.BasicAuth) (db.SavedTeam, error) {
	fake.updateBasicAuthMutex.Lock()
	ret, specificReturn := fake.updateBasicAuthReturnsOnCall[len(fake.updateBasicAuthArgsForCall)]
	fake.updateBasicAuthArgsForCall = append(fake.updateBasicAuthArgsForCall, struct {
		basicAuth *db.BasicAuth
	}{basicAuth})
	fake.recordInvocation("UpdateBasicAuth", []interface{}{basicAuth})
	fake.updateBasicAuthMutex.Unlock()
	if fake.UpdateBasicAuthStub != nil {
		return fake.UpdateBasicAuthStub(basicAuth)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateBasicAuthReturns.result1, fake.updateBasicAuthReturns.result2
}

func (fake *FakeTeamDB) UpdateBasicAuthCallCount() int {
	fake.updateBasicAuthMutex.RLock()
	defer fake.updateBasicAuthMutex.RUnlock()
	return len(fake.updateBasicAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateBasicAuthArgsForCall(i int) *db.BasicAuth {
	fake.updateBasicAuthMutex.RLock()
	defer fake.updateBasicAuthMutex.RUnlock()
	return fake.updateBasicAuthArgsForCall[i].basicAuth
}

func (fake *FakeTeamDB) UpdateBasicAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateBasicAuthStub = nil
	fake.updateBasicAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateBasicAuthReturnsOnCall(i int, result1 db.SavedTeam, result2 error) {
	fake.UpdateBasicAuthStub = nil
	if fake.updateBasicAuthReturnsOnCall == nil {
		fake.updateBasicAuthReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 error
		})
	}
	fake.updateBasicAuthReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateGitHubAuth(gitHubAuth *db.GitHubAuth) (db.SavedTeam, error) {
	fake.updateGitHubAuthMutex.Lock()
	ret, specificReturn := fake.updateGitHubAuthReturnsOnCall[len(fake.updateGitHubAuthArgsForCall)]
	fake.updateGitHubAuthArgsForCall = append(fake.updateGitHubAuthArgsForCall, struct {
		gitHubAuth *db.GitHubAuth
	}{gitHubAuth})
	fake.recordInvocation("UpdateGitHubAuth", []interface{}{gitHubAuth})
	fake.updateGitHubAuthMutex.Unlock()
	if fake.UpdateGitHubAuthStub != nil {
		return fake.UpdateGitHubAuthStub(gitHubAuth)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateGitHubAuthReturns.result1, fake.updateGitHubAuthReturns.result2
}

func (fake *FakeTeamDB) UpdateGitHubAuthCallCount() int {
	fake.updateGitHubAuthMutex.RLock()
	defer fake.updateGitHubAuthMutex.RUnlock()
	return len(fake.updateGitHubAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateGitHubAuthArgsForCall(i int) *db.GitHubAuth {
	fake.updateGitHubAuthMutex.RLock()
	defer fake.updateGitHubAuthMutex.RUnlock()
	return fake.updateGitHubAuthArgsForCall[i].gitHubAuth
}

func (fake *FakeTeamDB) UpdateGitHubAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateGitHubAuthStub = nil
	fake.updateGitHubAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateGitHubAuthReturnsOnCall(i int, result1 db.SavedTeam, result2 error) {
	fake.UpdateGitHubAuthStub = nil
	if fake.updateGitHubAuthReturnsOnCall == nil {
		fake.updateGitHubAuthReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 error
		})
	}
	fake.updateGitHubAuthReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateUAAAuth(uaaAuth *db.UAAAuth) (db.SavedTeam, error) {
	fake.updateUAAAuthMutex.Lock()
	ret, specificReturn := fake.updateUAAAuthReturnsOnCall[len(fake.updateUAAAuthArgsForCall)]
	fake.updateUAAAuthArgsForCall = append(fake.updateUAAAuthArgsForCall, struct {
		uaaAuth *db.UAAAuth
	}{uaaAuth})
	fake.recordInvocation("UpdateUAAAuth", []interface{}{uaaAuth})
	fake.updateUAAAuthMutex.Unlock()
	if fake.UpdateUAAAuthStub != nil {
		return fake.UpdateUAAAuthStub(uaaAuth)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateUAAAuthReturns.result1, fake.updateUAAAuthReturns.result2
}

func (fake *FakeTeamDB) UpdateUAAAuthCallCount() int {
	fake.updateUAAAuthMutex.RLock()
	defer fake.updateUAAAuthMutex.RUnlock()
	return len(fake.updateUAAAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateUAAAuthArgsForCall(i int) *db.UAAAuth {
	fake.updateUAAAuthMutex.RLock()
	defer fake.updateUAAAuthMutex.RUnlock()
	return fake.updateUAAAuthArgsForCall[i].uaaAuth
}

func (fake *FakeTeamDB) UpdateUAAAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateUAAAuthStub = nil
	fake.updateUAAAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateUAAAuthReturnsOnCall(i int, result1 db.SavedTeam, result2 error) {
	fake.UpdateUAAAuthStub = nil
	if fake.updateUAAAuthReturnsOnCall == nil {
		fake.updateUAAAuthReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 error
		})
	}
	fake.updateUAAAuthReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateGenericOAuth(genericOAuth *db.GenericOAuth) (db.SavedTeam, error) {
	fake.updateGenericOAuthMutex.Lock()
	ret, specificReturn := fake.updateGenericOAuthReturnsOnCall[len(fake.updateGenericOAuthArgsForCall)]
	fake.updateGenericOAuthArgsForCall = append(fake.updateGenericOAuthArgsForCall, struct {
		genericOAuth *db.GenericOAuth
	}{genericOAuth})
	fake.recordInvocation("UpdateGenericOAuth", []interface{}{genericOAuth})
	fake.updateGenericOAuthMutex.Unlock()
	if fake.UpdateGenericOAuthStub != nil {
		return fake.UpdateGenericOAuthStub(genericOAuth)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateGenericOAuthReturns.result1, fake.updateGenericOAuthReturns.result2
}

func (fake *FakeTeamDB) UpdateGenericOAuthCallCount() int {
	fake.updateGenericOAuthMutex.RLock()
	defer fake.updateGenericOAuthMutex.RUnlock()
	return len(fake.updateGenericOAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateGenericOAuthArgsForCall(i int) *db.GenericOAuth {
	fake.updateGenericOAuthMutex.RLock()
	defer fake.updateGenericOAuthMutex.RUnlock()
	return fake.updateGenericOAuthArgsForCall[i].genericOAuth
}

func (fake *FakeTeamDB) UpdateGenericOAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateGenericOAuthStub = nil
	fake.updateGenericOAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateGenericOAuthReturnsOnCall(i int, result1 db.SavedTeam, result2 error) {
	fake.UpdateGenericOAuthStub = nil
	if fake.updateGenericOAuthReturnsOnCall == nil {
		fake.updateGenericOAuthReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 error
		})
	}
	fake.updateGenericOAuthReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetConfig(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	ret, specificReturn := fake.getConfigReturnsOnCall[len(fake.getConfigArgsForCall)]
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.recordInvocation("GetConfig", []interface{}{pipelineName})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub(pipelineName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3, fake.getConfigReturns.result4
}

func (fake *FakeTeamDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeTeamDB) GetConfigArgsForCall(i int) string {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return fake.getConfigArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetConfigReturns(result1 atc.Config, result2 atc.RawConfig, result3 db.ConfigVersion, result4 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTeamDB) GetConfigReturnsOnCall(i int, result1 atc.Config, result2 atc.RawConfig, result3 db.ConfigVersion, result4 error) {
	fake.GetConfigStub = nil
	if fake.getConfigReturnsOnCall == nil {
		fake.getConfigReturnsOnCall = make(map[int]struct {
			result1 atc.Config
			result2 atc.RawConfig
			result3 db.ConfigVersion
			result4 error
		})
	}
	fake.getConfigReturnsOnCall[i] = struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTeamDB) SaveConfigToBeDeprecated(arg1 string, arg2 atc.Config, arg3 db.ConfigVersion, arg4 db.PipelinePausedState) (db.SavedPipeline, bool, error) {
	fake.saveConfigToBeDeprecatedMutex.Lock()
	ret, specificReturn := fake.saveConfigToBeDeprecatedReturnsOnCall[len(fake.saveConfigToBeDeprecatedArgsForCall)]
	fake.saveConfigToBeDeprecatedArgsForCall = append(fake.saveConfigToBeDeprecatedArgsForCall, struct {
		arg1 string
		arg2 atc.Config
		arg3 db.ConfigVersion
		arg4 db.PipelinePausedState
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SaveConfigToBeDeprecated", []interface{}{arg1, arg2, arg3, arg4})
	fake.saveConfigToBeDeprecatedMutex.Unlock()
	if fake.SaveConfigToBeDeprecatedStub != nil {
		return fake.SaveConfigToBeDeprecatedStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.saveConfigToBeDeprecatedReturns.result1, fake.saveConfigToBeDeprecatedReturns.result2, fake.saveConfigToBeDeprecatedReturns.result3
}

func (fake *FakeTeamDB) SaveConfigToBeDeprecatedCallCount() int {
	fake.saveConfigToBeDeprecatedMutex.RLock()
	defer fake.saveConfigToBeDeprecatedMutex.RUnlock()
	return len(fake.saveConfigToBeDeprecatedArgsForCall)
}

func (fake *FakeTeamDB) SaveConfigToBeDeprecatedArgsForCall(i int) (string, atc.Config, db.ConfigVersion, db.PipelinePausedState) {
	fake.saveConfigToBeDeprecatedMutex.RLock()
	defer fake.saveConfigToBeDeprecatedMutex.RUnlock()
	return fake.saveConfigToBeDeprecatedArgsForCall[i].arg1, fake.saveConfigToBeDeprecatedArgsForCall[i].arg2, fake.saveConfigToBeDeprecatedArgsForCall[i].arg3, fake.saveConfigToBeDeprecatedArgsForCall[i].arg4
}

func (fake *FakeTeamDB) SaveConfigToBeDeprecatedReturns(result1 db.SavedPipeline, result2 bool, result3 error) {
	fake.SaveConfigToBeDeprecatedStub = nil
	fake.saveConfigToBeDeprecatedReturns = struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) SaveConfigToBeDeprecatedReturnsOnCall(i int, result1 db.SavedPipeline, result2 bool, result3 error) {
	fake.SaveConfigToBeDeprecatedStub = nil
	if fake.saveConfigToBeDeprecatedReturnsOnCall == nil {
		fake.saveConfigToBeDeprecatedReturnsOnCall = make(map[int]struct {
			result1 db.SavedPipeline
			result2 bool
			result3 error
		})
	}
	fake.saveConfigToBeDeprecatedReturnsOnCall[i] = struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) CreateOneOffBuild() (db.Build, error) {
	fake.createOneOffBuildMutex.Lock()
	ret, specificReturn := fake.createOneOffBuildReturnsOnCall[len(fake.createOneOffBuildArgsForCall)]
	fake.createOneOffBuildArgsForCall = append(fake.createOneOffBuildArgsForCall, struct{}{})
	fake.recordInvocation("CreateOneOffBuild", []interface{}{})
	fake.createOneOffBuildMutex.Unlock()
	if fake.CreateOneOffBuildStub != nil {
		return fake.CreateOneOffBuildStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createOneOffBuildReturns.result1, fake.createOneOffBuildReturns.result2
}

func (fake *FakeTeamDB) CreateOneOffBuildCallCount() int {
	fake.createOneOffBuildMutex.RLock()
	defer fake.createOneOffBuildMutex.RUnlock()
	return len(fake.createOneOffBuildArgsForCall)
}

func (fake *FakeTeamDB) CreateOneOffBuildReturns(result1 db.Build, result2 error) {
	fake.CreateOneOffBuildStub = nil
	fake.createOneOffBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) CreateOneOffBuildReturnsOnCall(i int, result1 db.Build, result2 error) {
	fake.CreateOneOffBuildStub = nil
	if fake.createOneOffBuildReturnsOnCall == nil {
		fake.createOneOffBuildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 error
		})
	}
	fake.createOneOffBuildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPrivateAndPublicBuilds(page db.Page) ([]db.Build, db.Pagination, error) {
	fake.getPrivateAndPublicBuildsMutex.Lock()
	ret, specificReturn := fake.getPrivateAndPublicBuildsReturnsOnCall[len(fake.getPrivateAndPublicBuildsArgsForCall)]
	fake.getPrivateAndPublicBuildsArgsForCall = append(fake.getPrivateAndPublicBuildsArgsForCall, struct {
		page db.Page
	}{page})
	fake.recordInvocation("GetPrivateAndPublicBuilds", []interface{}{page})
	fake.getPrivateAndPublicBuildsMutex.Unlock()
	if fake.GetPrivateAndPublicBuildsStub != nil {
		return fake.GetPrivateAndPublicBuildsStub(page)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getPrivateAndPublicBuildsReturns.result1, fake.getPrivateAndPublicBuildsReturns.result2, fake.getPrivateAndPublicBuildsReturns.result3
}

func (fake *FakeTeamDB) GetPrivateAndPublicBuildsCallCount() int {
	fake.getPrivateAndPublicBuildsMutex.RLock()
	defer fake.getPrivateAndPublicBuildsMutex.RUnlock()
	return len(fake.getPrivateAndPublicBuildsArgsForCall)
}

func (fake *FakeTeamDB) GetPrivateAndPublicBuildsArgsForCall(i int) db.Page {
	fake.getPrivateAndPublicBuildsMutex.RLock()
	defer fake.getPrivateAndPublicBuildsMutex.RUnlock()
	return fake.getPrivateAndPublicBuildsArgsForCall[i].page
}

func (fake *FakeTeamDB) GetPrivateAndPublicBuildsReturns(result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.GetPrivateAndPublicBuildsStub = nil
	fake.getPrivateAndPublicBuildsReturns = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetPrivateAndPublicBuildsReturnsOnCall(i int, result1 []db.Build, result2 db.Pagination, result3 error) {
	fake.GetPrivateAndPublicBuildsStub = nil
	if fake.getPrivateAndPublicBuildsReturnsOnCall == nil {
		fake.getPrivateAndPublicBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.Build
			result2 db.Pagination
			result3 error
		})
	}
	fake.getPrivateAndPublicBuildsReturnsOnCall[i] = struct {
		result1 []db.Build
		result2 db.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetContainer(handle string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	ret, specificReturn := fake.getContainerReturnsOnCall[len(fake.getContainerArgsForCall)]
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("GetContainer", []interface{}{handle})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
}

func (fake *FakeTeamDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeTeamDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].handle
}

func (fake *FakeTeamDB) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetContainerReturnsOnCall(i int, result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	if fake.getContainerReturnsOnCall == nil {
		fake.getContainerReturnsOnCall = make(map[int]struct {
			result1 db.SavedContainer
			result2 bool
			result3 error
		})
	}
	fake.getContainerReturnsOnCall[i] = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) FindContainersByDescriptors(id db.Container) ([]db.SavedContainer, error) {
	fake.findContainersByDescriptorsMutex.Lock()
	ret, specificReturn := fake.findContainersByDescriptorsReturnsOnCall[len(fake.findContainersByDescriptorsArgsForCall)]
	fake.findContainersByDescriptorsArgsForCall = append(fake.findContainersByDescriptorsArgsForCall, struct {
		id db.Container
	}{id})
	fake.recordInvocation("FindContainersByDescriptors", []interface{}{id})
	fake.findContainersByDescriptorsMutex.Unlock()
	if fake.FindContainersByDescriptorsStub != nil {
		return fake.FindContainersByDescriptorsStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findContainersByDescriptorsReturns.result1, fake.findContainersByDescriptorsReturns.result2
}

func (fake *FakeTeamDB) FindContainersByDescriptorsCallCount() int {
	fake.findContainersByDescriptorsMutex.RLock()
	defer fake.findContainersByDescriptorsMutex.RUnlock()
	return len(fake.findContainersByDescriptorsArgsForCall)
}

func (fake *FakeTeamDB) FindContainersByDescriptorsArgsForCall(i int) db.Container {
	fake.findContainersByDescriptorsMutex.RLock()
	defer fake.findContainersByDescriptorsMutex.RUnlock()
	return fake.findContainersByDescriptorsArgsForCall[i].id
}

func (fake *FakeTeamDB) FindContainersByDescriptorsReturns(result1 []db.SavedContainer, result2 error) {
	fake.FindContainersByDescriptorsStub = nil
	fake.findContainersByDescriptorsReturns = struct {
		result1 []db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) FindContainersByDescriptorsReturnsOnCall(i int, result1 []db.SavedContainer, result2 error) {
	fake.FindContainersByDescriptorsStub = nil
	if fake.findContainersByDescriptorsReturnsOnCall == nil {
		fake.findContainersByDescriptorsReturnsOnCall = make(map[int]struct {
			result1 []db.SavedContainer
			result2 error
		})
	}
	fake.findContainersByDescriptorsReturnsOnCall[i] = struct {
		result1 []db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPipelinesMutex.RLock()
	defer fake.getPipelinesMutex.RUnlock()
	fake.getPublicPipelinesMutex.RLock()
	defer fake.getPublicPipelinesMutex.RUnlock()
	fake.getPrivateAndAllPublicPipelinesMutex.RLock()
	defer fake.getPrivateAndAllPublicPipelinesMutex.RUnlock()
	fake.getPipelineByNameMutex.RLock()
	defer fake.getPipelineByNameMutex.RUnlock()
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	fake.updateBasicAuthMutex.RLock()
	defer fake.updateBasicAuthMutex.RUnlock()
	fake.updateGitHubAuthMutex.RLock()
	defer fake.updateGitHubAuthMutex.RUnlock()
	fake.updateUAAAuthMutex.RLock()
	defer fake.updateUAAAuthMutex.RUnlock()
	fake.updateGenericOAuthMutex.RLock()
	defer fake.updateGenericOAuthMutex.RUnlock()
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	fake.saveConfigToBeDeprecatedMutex.RLock()
	defer fake.saveConfigToBeDeprecatedMutex.RUnlock()
	fake.createOneOffBuildMutex.RLock()
	defer fake.createOneOffBuildMutex.RUnlock()
	fake.getPrivateAndPublicBuildsMutex.RLock()
	defer fake.getPrivateAndPublicBuildsMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.findContainersByDescriptorsMutex.RLock()
	defer fake.findContainersByDescriptorsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamDB) recordInvocation(key string, args []interface{}) {
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

var _ db.TeamDB = new(FakeTeamDB)
