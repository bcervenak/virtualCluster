package discovery

import (
	"github.com/golang/glog"

	"github.com/songbinliu/containerChain/pkg/registration"
	"github.com/songbinliu/containerChain/pkg/target"

	sdkprobe "github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"fmt"
)

type DiscoveryClient struct {
	targetConfig *TargetConf
	cluster      *target.Cluster
}

func NewDiscoveryClient(targetConfig *TargetConf, cluster *target.Cluster) *DiscoveryClient {
	return &DiscoveryClient{
		targetConfig: targetConfig,
		cluster:      cluster,
	}
}

func (dc *DiscoveryClient) GetAccountValues() *sdkprobe.TurboTargetInfo {
	var accountValues []*proto.AccountValue

	targetConf := dc.targetConfig
	// Convert all parameters in clientConf to AccountValue list
	targetID := registration.TargetIdentifierField
	accVal := &proto.AccountValue{
		Key:         &targetID,
		StringValue: &targetConf.Address,
	}
	accountValues = append(accountValues, accVal)

	username := registration.Username
	accVal = &proto.AccountValue{
		Key:         &username,
		StringValue: &targetConf.Username,
	}
	accountValues = append(accountValues, accVal)

	password := registration.Password
	accVal = &proto.AccountValue{
		Key:         &password,
		StringValue: &targetConf.Password,
	}
	accountValues = append(accountValues, accVal)

	targetInfo := sdkprobe.NewTurboTargetInfoBuilder(targetConf.ProbeCategory, targetConf.TargetType, targetID, accountValues).Create()

	glog.V(2).Infof("Got AccountValues for target:%v", targetConf.Address)
	return targetInfo
}

func (dc *DiscoveryClient) Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error) {
	glog.V(2).Infof("begin to validating target...")

	return &proto.ValidationResponse{}, nil
}

func printDTOs(dtos []*proto.EntityDTO) {
	msg := ""
	for _, dto := range dtos {
		line := fmt.Sprintf("%+v", dto)
		msg = msg + "\n" + line
	}

	glog.V(3).Infof("%s", msg)
}

func (dc *DiscoveryClient) Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	//TODO: implement it

	glog.V(2).Infof("begin to discovery target...")

	resultDTOs, err := dc.cluster.GenerateDTOs()
	if err != nil {
		glog.Errorf("failed to generate DTOs: %v", err)
		resultDTOs = []*proto.EntityDTO{}
	}

	printDTOs(resultDTOs)

	response := &proto.DiscoveryResponse{
		EntityDTO: resultDTOs,
	}

	glog.V(2).Infof("end of discoverying target. [%d]", len(resultDTOs))

	return response, nil
}