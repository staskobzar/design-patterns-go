package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type AnalyticsTracker interface{
	track(event string, data map[string]any)
}

type LegacyAnalyticsService struct{}

func (*LegacyAnalyticsService) logEvent(name, data string) {
	fmt.Printf("[+] LegacyAnalyticsService: Logging %q with data:\n%q\n", name,data)
}

type LegacyAnalyticsAdapter struct {
	service *LegacyAnalyticsService
}

func NewLegacyAnalyticsAdapter() *LegacyAnalyticsAdapter{
	return &LegacyAnalyticsAdapter{service: &LegacyAnalyticsService{}}
}

func (a*LegacyAnalyticsAdapter) track(event string, data map[string]any) {
	jdata,err:=json.Marshal(data)
	if err !=nil{
		panic(err)
	}

	a.service.logEvent(event,string(jdata))
}

// consumer function
func signUp(user string, at AnalyticsTracker) {
	at.track("UserSignup",
		map[string]any{
		"username": user,
			"timestamp": time.Now(),
			"referral": "LandingPage",
	})
}

func main() {
	fmt.Println("=== ADAPTER")

	signUp("Maxim123", NewLegacyAnalyticsAdapter())

}
