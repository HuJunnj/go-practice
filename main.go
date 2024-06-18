package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type CommonJson struct {
	ServiceDeploymentLocation ServiceDeploymentLocation `json:"service_deployment_location"`
	DeploymentLocationConfig  DeploymentLocationConfig  `json:"deployment_location_config"`
	ApplicationSettingDebug   bool                      `json:"application_setting_debug"`
}
type ServiceDeploymentLocation struct {
	MqttBroker                    string `json:"mqtt_broker"`
	ControllerWebService          string `json:"controller_web_service"`
	ActiveAlarmService            string `json:"active_alarm_service"`
	LogHistoryService             string `json:"log_history_service"`
	ControllerRoutineService      string `json:"controller_routine_service"`
	BroadcastListenerOnController string `json:"broadcast_listener_on_controller"`
	VisionSocket                  string `json:"vision_socket"`
	ControllerGrpcService         string `json:"controller_grpc_service"`
	PureWebServer                 string `json:"pure_web_server"`
	SubpubService                 string `json:"subpub_service"`
	StatusServerWeb               string `json:"status_server_web"`
	TpHardwareWebService          string `json:"tp_hardware_web_service"`
	TpDevServer                   string `json:"tp_dev_server"`
	ControllerReqProxy            string `json:"controller_req_proxy"`
	TpHardwareRoutineNode         string `json:"tp_hardware_routine_node"`
	MqttMonitorNode               string `json:"mqtt_monitor_node"`
	TpTeach                       string `json:"tp_teach"`
}

type DeploymentLocationConfig struct {
	Tp         DeploymentLocationConfigItem `json:"tp"`
	Controller DeploymentLocationConfigItem `json:"controller"`
	Pc         DeploymentLocationConfigItem `json:"pc"`
}

type DeploymentLocationConfigItem struct {
	CommunicationPartnerIp CommunicationPartnerIp `json:"communication_partner_ip"`
}

type CommunicationPartnerIp struct {
	Tp         string `json:"tp"`
	Controller string `json:"controller"`
	Pc         string `json:"pc"`
}

type Item struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}

var c CommonJson

func main() {
	ip, err := getWirelessLANIP()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wireless LAN IP: %s\n", ip)
	}

}

// ReadCommonJson 读取配置文件
func ReadCommonJson() {
	for {
		time.Sleep(3 * time.Second)
		ReadCommonJsontest()

	}
}
func getWirelessLANIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		if inter.Flags&net.FlagUp == 0 {
			continue // 接口必须是启用状态
		}
		if inter.Flags&net.FlagLoopback != 0 {
			continue // 排除loopback接口
		}

		addrs, err := inter.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ip := addr.(*net.IPNet).IP
			if ip.To4() != nil && !ip.IsLoopback() {
				return ip.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no wireless LAN IP address found")
}
func ReadCommonJsontest() error {
	file, err := os.Open("D:\\GIT\\commom.json")
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return errors.New("空文件")
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
