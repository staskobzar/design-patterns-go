package main

import "fmt"

type PushService interface {
	sendPush(msg string)
}

type DeviceInfoService interface {
	getDeviceModel() string
}

type MobileService interface {
	createPushService() PushService
	createDeviceInfoService() DeviceInfoService
}

// Google

type FcmPushService struct{}

func (FcmPushService) sendPush(msg string) {
	fmt.Printf("Google: Sending push via FCM -> %q\n", msg)
}

type GoogleDeviceInfoService struct{}

func (GoogleDeviceInfoService) getDeviceModel() string {
	return "Google Pixel 8"
}

// factory implementation
type GoogleFactory struct{}

func (GoogleFactory) createPushService() PushService             { return FcmPushService{} }
func (GoogleFactory) createDeviceInfoService() DeviceInfoService { return GoogleDeviceInfoService{} }

// Huawei
type HuaweiPushService struct{}

func (HuaweiPushService) sendPush(msg string) {
	fmt.Printf("Huawei: Sending push via HMS -> %q\n", msg)
}

type HuaweiDeviceInfoService struct{}

func (HuaweiDeviceInfoService) getDeviceModel() string {
	return "Huawei P50 Pro"
}

// factory implementaion
type HuaweiFactory struct{}

func (HuaweiFactory) createPushService() PushService             { return HuaweiPushService{} }
func (HuaweiFactory) createDeviceInfoService() DeviceInfoService { return HuaweiDeviceInfoService{} }

// abstract creators

func RunMobileService(ms MobileService) {
	ms.createPushService().sendPush("Running Mobile Service")
	fmt.Println("... get device ...")
	fmt.Println(ms.createDeviceInfoService().getDeviceModel())
	fmt.Println("[+] done")
}

func main() {
	RunMobileService(HuaweiFactory{})
	RunMobileService(GoogleFactory{})
	fmt.Println("--------------------------------------------------")
}
