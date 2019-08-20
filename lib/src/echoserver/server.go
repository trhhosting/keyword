package echoserver

import (
	"fmt"
	"gitlab.com/labstack/echo"
	"gitlab.com/labstack/echo/middleware"
	"log"
	"net"
	"os"
	"path/filepath"
	_ "github.com/trhhosting/keyword/lib/src/database"
	"strings"
	"time"
	"strconv"
	"sc.tpnfc.us/AdFreeZone/afzAccessSever/lib/src/devices"
	"sc.tpnfc.us/AdFreeZone/afzAccessSever/lib/src/nameservers"

)

func check(ip string) bool {
	trial := net.ParseIP(ip)
	if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo" {
		fmt.Println(trial)
		fmt.Println(trial.To4())
	}

	if trial.To4() == nil {
		//fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	return true
}

type TLMSMSServer struct {
	Port      string
	IPAddress string
}

func (t TLMSMSServer) TLMServer() {
	// Trial User to get JWT
	e.POST("/msg", afzHandler)
	// Paid User to get JWT
	e.POST("/dns", dnsServiceHandler)
	// Report Adverts
	e.POST("/ar", advertReportHandler)
	// Add device to a current user account
	e.POST("/add/:account", addDeviceToAccountHandler)
	// Paid User Account Management
	e.POST("/user/manage/:account", userAccountManagementHandler)
	// Paid User Accounts
	e.POST("/user/:account", userAccountServiceHandler)
	// Shop
	e.POST("/shop/:plan/:account", shopAllPlansHandler)
	// Shop
	e.POST("/shop/:plan/:account", shopAllPlansHandler)

	// Shop Coffee
	e.GET("/shop/coffee/:account", shopAllCoffeeHandler)
	// Shop Meal
	e.GET("/shop/meal/:account", shopAllMealHandler)

	//// Home Page
	//e.GET("/index.html", homeHandler)
	// Home Page
	e.GET("/home", homeHandler)
	// Help Page send to forum
	e.POST("/help/:account", helpHandler)

	// static files
	//path := "/Source/golang/src/sc.tpnfc.us/oyoshi/tlmtext"
	path, _ := os.Getwd()
	path = filepath.Join(path, "frontend", "deploy")
	path = strings.TrimPrefix(path, "C:")
	path = strings.Replace(path, "\\", "/", 100)
	path = strings.Replace(path, "lib/src/echoserver/", "", 1)
	if os.Getenv("nmpw3xEcjogyzbfA4uqr") == "zj9shEHitCdFwvmarygo" {
		fmt.Println(path)

	}
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   path,
		Browse: true,
		HTML5:  true,
		Index:  "index.html",
	}))
	//e.Static("/main.dart.js", path + "app/published/web/main.dart.js")
	//e.Static("/static/", path + "app/published/web/static")
	//e.Static("/packages/", path + "app/published/web/packages")

	if t.IPAddress != "" {
		x := check(t.IPAddress)
		if x {
			startStr := fmt.Sprintf("%v:%v", t.IPAddress, t.Port)
			e.Start(startStr)

		} else {
			startStr := fmt.Sprintf("%v:%v", "0.0.0.0", t.Port)
			e.Start(startStr)
		}
	} else {
		startStr := fmt.Sprintf("%v:%v", "localhost", t.Port)
		e.Start(startStr)
	}

}

func afzHandler(c echo.Context) (err error) {
	dat, _ := c.FormParams()
	var deviceData = new(devices.DeviceReceived)
	for key, value := range dat {

		switch key {
		case "UUID":
			deviceData.UUID = value[0]
		}
		log.Printf("key:\t %v\nvalue: %v\n", key, value[0])
	}
	var dvc = devices.Device{
		Received: *deviceData,
	}

	data, err := dvc.AddDeviceEntry()
	if err != nil {
		log.Print(err)
	}

	//TODO add some msg to Message2 its currently blank
	var html string
	if data.Expires.After(time.Now()) {
		html = fmt.Sprintf("SerialNo:%v\nOwnerKey:%v\nMessage:%v\nExpires:%vMessage2:%v\nuuid:%v", data.SerialNo, data.OwnerKey, data.Message, data.Expires, "", deviceData.UUID)
		fmt.Println(html)
	} else {
		var NoTActive = "dns0=1.1.1.1\ndns1=80.80.80.80\ndns2=80.80.81.81\ndns3=1.1.1.1\ndns4=1.1.1.1\ndns5=1.1.1.1\ndns6=1.1.1.1\n"
		html = fmt.Sprintf("SerialNo:%v\nOwnerKey:%v\nMessage:%v\nExpires:%v\n%vMessage2:%v\nuuid:%v", data.SerialNo, data.OwnerKey, data.Message, data.Expires, NoTActive, "", deviceData.UUID)
		fmt.Println(html)
	}
	return c.HTML(200, html)
}

func dnsServiceHandler(c echo.Context) (err error) {
	dat, _ := c.FormParams()
	var deviceData = new(devices.DeviceReceived)
	for key, value := range dat {

		switch key {
		case "UUID":
			deviceData.UUID = value[0]
		case "SerialNo":
			deviceData.SerialNo = value[0]
		case "OwnerKey":
			deviceData.OwnerKey = value[0]
		case "InUse":
			if value[0] == "true" || value[0] == "True" {
				deviceData.InUse = true
			} else {
				deviceData.InUse = false
			}
		}
		log.Printf("key:\t %v\nvalue: %v\n", key, value[0])
	}
	var dvc = devices.Device{
		Received: *deviceData,
	}

	data, err := dvc.AddDeviceEntry()
	if err != nil {
		log.Print(err)
	}
	ns, er := nameservers.MyDnsList("")
	if er != nil {
		log.Print(err)
	}

	var dnsLts string
	for key, value := range ns {
		dnsLts = fmt.Sprintf("%vdns%v=%v\n", dnsLts, key, value.IPAddress)

	}
	//TODO add some msg to Message2 its currently blank
	var html string
	if data.Expires.After(time.Now()) {
		html = fmt.Sprintf("Message:%v\nMessage2:%v\n%v",data.Message, "", dnsLts)
		fmt.Println(html)
	} else {
		var NoTActive = "dns0=1.1.1.1\ndns1=80.80.80.80\ndns2=80.80.81.81\ndns3=1.1.1.1\ndns4=1.1.1.1\ndns5=1.1.1.1\ndns6=1.1.1.1\n"
		html = fmt.Sprintf("Message:%v\nMessage2:%v\n%v",data.Message, "Buy me a meal to upgrade.", NoTActive)
		fmt.Println(html)
	}

	return c.HTML(200, html)
	return
}

func advertReportHandler(c echo.Context)  ( err error) {
	dat, _ := c.FormParams()
	var deviceData = new(devices.Advert)
	for key, value := range dat {
		switch key {
		case "UUID":
			deviceData.UUID = value[0]
		case "TimeFrame":
			tmp, _  := strconv.Atoi(value[0])
			if tmp == 5 || tmp == 10 || tmp == 15 {
				deviceData.TimeFrame = tmp
			}
		}
		log.Printf("key:\t %v\nvalue: %v\n", key, value[0])
	}
	deviceData.IPAddress = c.RealIP()
	devices.AdvertReportCreate(*deviceData)
	return
}
func addDeviceToAccountHandler(c echo.Context) (err error) {
	return c.HTML(200, "addDeviceToAccountHandler")
	return
}
func userAccountManagementHandler(c echo.Context) (err error) {
	return c.HTML(200, "userAccountManagementHandler")
	return
}
func userAccountServiceHandler(c echo.Context) (err error) {
	return c.HTML(200, "User Service")
	return
}
func shopAllPlansHandler(c echo.Context) (err error) {
	return c.HTML(200, "shopAllPlansHandler")
	return
}
func shopAllCoffeeHandler(c echo.Context) (err error) {
	return c.HTML(200, "shopAllPlansHandler")
	return
}
func shopAllMealHandler(c echo.Context) (err error) {
	return c.HTML(200, "shopAllPlansHandler")
	return
}
func shopAllGroceriesHandler(c echo.Context) (err error) {
	return c.HTML(200, "shopAllPlansHandler")
	return
}

func homeHandler(c echo.Context) (err error) {

	return c.HTML(200, "home")
}
func helpHandler(c echo.Context) (err error) {

	return c.HTML(200, "Help ")
}
