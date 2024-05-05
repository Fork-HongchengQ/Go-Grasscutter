package handler

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/consts"
	"Go-Grasscutter/src/server/event/dispatch"
	"Go-Grasscutter/src/server/http/object"
	"Go-Grasscutter/src/utils"
	"Go-Grasscutter/src/utils/crypto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/golang/protobuf/proto"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	regions              sync.Map // map[string]RegionData
	regionListResponse   string
	regionListResponseCN string
	doOnce               sync.Once
)

// Initialize Configures region data according to configuration.
func Initialize() {
	var head string
	c := config.Conf
	httpInfo := c.Server.Http
	// todo use https
	if httpInfo.Encryption.UseInRouting {
		head = "https://"
	} else {
		head = "http://"
	}
	head = "http://"

	dispatchDomain :=
		head +
			config.LrString(httpInfo.AccessAddress, httpInfo.BindAddress) +
			":" +
			strconv.Itoa(config.LrInt(httpInfo.AccessPort, httpInfo.BindPort))
	var (
		servers           = make([]*pb.RegionSimpleInfo, 0)
		usedNames         = make(map[string]struct{})
		configuredRegions = make([]*config.Region, 0) // todo raed config data
	)

	// todo getRunMode() != ServerRunMode.HYBRID && set const "HYBRID"
	if c.Server.RunMode != "HYBRID" && len(configuredRegions) == 0 {
		log.Fatal("[Dispatch] There are no game servers available. Exiting due to unplayable state.")
	} else if len(configuredRegions) == 0 {
		gameInfo := c.Server.Game
		configuredRegions = append(configuredRegions, &config.Region{
			Name:  "os_usa",
			Title: c.Server.Dispatch.DefaultName,
			Ip:    config.LrString(gameInfo.AccessAddress, gameInfo.BindAddress),
			Port:  config.LrInt(gameInfo.AccessPort, gameInfo.BindPort),
		})
	}

	for _, region := range configuredRegions {
		if _, ok := usedNames[region.Name]; ok {
			log.Println("Region name already in use.")
			continue
		}
		// Create a region identifier.
		identifier := &pb.RegionSimpleInfo{
			Name:        region.Name,
			Title:       region.Title,
			Type:        "DEV_PUBLIC",
			DispatchUrl: dispatchDomain + "/query_cur_region/" + region.Name,
		}
		usedNames[region.Name] = struct{}{}
		servers = append(servers, identifier)

		// Create a region info object.
		// todo load config region.Name
		updatedQuery := &pb.QueryCurrRegionHttpRsp{
			RegionInfo: &pb.RegionInfo{
				GateserverIp:   region.Ip,
				GateserverPort: uint32(region.Port),
			},
			ClientSecretKey: crypto.DispatchSeed,
		}
		data, err := proto.Marshal(updatedQuery)
		if err != nil {
			panic(err)
		}
		regions.Store(region.Name, &RegionData{
			RegionQuery: updatedQuery,
			Base64:      utils.Base64Encode(data),
		})
	}

	// Determine config settings.
	// Create a config object.
	customConfig := struct {
		Sdkenv        string `json:"sdkenv"`
		CheckDevice   string `json:"checkdevice"`
		LoadPatch     string `json:"loadPatch"`
		ShowException string `json:"showexception"`
		RegionConfig  string `json:"regionConfig"`
		DownloadMode  string `json:"downloadMode"`
		CodeSwitch    []int  `json:"codeSwitch"`
		CoverSwitch   []int  `json:"coverSwitch"`
	}{
		Sdkenv:        "2",
		CheckDevice:   "false",
		LoadPatch:     "false",
		ShowException: "false", // todo GameConstants.DEBUG
		RegionConfig:  "pm|fk|add",
		DownloadMode:  "0",
		CodeSwitch:    []int{3628}, // codeSwitch
		CoverSwitch:   []int{40},   // hiddenIcons
	}

	// XOR the config with the key.

	encodedConfig, err := json.Marshal(customConfig)

	encodedConfig = crypto.Xor(encodedConfig, crypto.DispatchKey)

	// Create an updated region list.
	updatedRegionList := pb.QueryRegionListHttpRsp{
		RegionList:                  servers,
		ClientSecretKey:             crypto.DispatchSeed,
		ClientCustomConfigEncrypted: encodedConfig,
		EnableLoginPc:               true,
	}
	// Set the region list response.
	regionListResponse = utils.Base64Encode([]byte(updatedRegionList.String()))

	// CN
	// Modify the existing config option.
	customConfig.Sdkenv = "0"
	// XOR the config with the key.
	encodedConfig, err = json.Marshal(customConfig)
	if err != nil {
		log.Println(err)
	}

	encodedConfig = crypto.Xor(encodedConfig, crypto.DispatchKey)
	// Create an updated region list.
	updatedRegionListCN := pb.QueryRegionListHttpRsp{
		RegionList:                  servers,
		ClientSecretKey:             crypto.DispatchSeed,
		ClientCustomConfigEncrypted: encodedConfig,
		EnableLoginPc:               true,
	}

	// todo fix it
	// Set the region list response.
	regionListResponseCN = utils.Base64Encode([]byte(updatedRegionListCN.String()))

}

type RegionData struct {
	RegionQuery *pb.QueryCurrRegionHttpRsp
	Base64      string
}

func ApplyRegionHandler(r *server.Hertz) {
	r.GET("/query_region_list", queryRegionList)
	// login/v2 -> /query_cur_region
	r.GET("/query_cur_region", queryCurrentRegion)
	r.GET("/query_cur_region/:region", queryCurrentRegion)
}

// todo fix it
func queryRegionList(c context.Context, ctx *app.RequestContext) {
	doOnce.Do(func() {
		Initialize()
	})
	if ctx.QueryArgs().Has("version") && ctx.QueryArgs().Has("platform") {
		versionName := ctx.Query("version")
		versionCode := versionName[0:8]
		//platformName := ctx.Query("platform")
		switch versionCode {
		case "CNRELiOS", "CNRELWin", "CNRELAnd":
			// Use the CN region list.
			// todo QueryAllRegionsEvent

			// Respond with the event result.
			ctx.String(200, regionListResponseCN)
		case "OSRELiOS", "OSRELWin", "OSRELAnd":
			// Use the OS region list.

			// Respond with the event result.
			ctx.String(200, regionListResponse)
		default:
			ctx.String(200, regionListResponse)
		}
	} else {
		// Use the default region list.

		// Respond with the event result.
		ctx.String(200, regionListResponse)
	}
	// Log the request to the console.

}

// GET /query_cur_region/
// os_usa?
// version=CNRELWin4.0.0
// &lang=2
// &platform=3
// &binary=1
// &time=778
// &channel_id=1
// &sub_channel_id=1
// &account_type=1
// &dispatchSeed=f586f0569733d33a
// &key_id=4
// &aid=10001
// todo check make game can connect to game server
func queryCurrentRegion(c context.Context, ctx *app.RequestContext) {
	doOnce.Do(func() {
		Initialize()
	})
	fmt.Println(ctx.ClientIP(), "visit QueryCurrentRegion")
	// Get region to query.
	regionName := ctx.Param("region")
	versionName := ctx.Query("version")
	// todo var region = regions.get(regionName);
	// Get region data.
	regionData := "CAESGE5vdCBGb3VuZCB2ZXJzaW9uIGNvbmZpZw=="
	val, ok := regions.Load(regionName)
	if ctx.QueryArgs().Len() > 0 {
		if ok {
			regionData = val.(*RegionData).Base64
		}
	}

	re := regexp.MustCompile("[a-zA-Z]")
	clientVersion := re.ReplaceAllString(versionName, "")
	versionCode := strings.Split(clientVersion, ".")

	versionMajor, _ := strconv.Atoi(versionCode[0])
	versionMinor, _ := strconv.Atoi(versionCode[1])
	versionFix, _ := strconv.Atoi(versionCode[2])
	if versionMajor >= 3 ||
		(versionMajor == 2 && versionMinor == 7 && versionFix >= 50) ||
		(versionMajor == 2 && versionMinor == 8) {
		event := dispatch.NewQueryCurrentRegionEvent(regionData)
		// event.call();
		key_id := ctx.Query("key_id")
		// The 'fix' or 'patch' version is not checked because it is only used
		// when miHoYo is desperate and fucks up big time.
		if versionMajor != consts.VersionParts[0] ||
			versionMinor != consts.VersionParts[1] {
			// Reject clients when there is a version mismatch
			updateClient := strings.Compare(consts.Version, clientVersion) > 0
			var contentMsg string
			if updateClient {
				contentMsg = fmt.Sprintf("\nVersion mismatch outdated client! \n\nServer version: %s\nClient version: %s",
					consts.Version, clientVersion)
			} else {
				contentMsg = fmt.Sprintf("\nVersion mismatch outdated server! \n\nServer version: %s\nClient version: %s",
					consts.Version, clientVersion)
			}
			rsp := pb.QueryCurrRegionHttpRsp{
				Retcode:    int32(pb.Retcode_RET_STOP_SERVER),
				Msg:        "Connection Failed!",
				RegionInfo: &pb.RegionInfo{},
				Detail: &pb.QueryCurrRegionHttpRsp_StopServer{
					StopServer: &pb.StopServerInfo{
						StopBeginTime: uint32(time.Now().Unix()),
						StopEndTime:   uint32(time.Now().Unix() + 1),
						Url:           "https://discord.gg/T5vZU6UyeG",
						ContentMsg:    contentMsg,
					},
				},
			}
			log.Printf("Connection denied for %s due to %t.\n", utils.Address(ctx), updateClient)
			ctx.JSON(200, &rsp)
			return
		}
		if len(ctx.Query("dispatchSeed")) == 0 {
			rsp := object.QueryCurRegionRspJson{
				Content: event.RegionInfo,
				Sign:    "TW9yZSBsb3ZlIGZvciBVQSBQYXRjaCBwbGF5ZXJz",
			}
			ctx.JSON(200, rsp)
			return
		}
		regionInfo := utils.Base64Decode(event.RegionInfo)
		rsp, err := crypto.EncryptAndSignRegionData(regionInfo, key_id)
		if err != nil {
			log.Println("An error occurred while handling query_cur_region.", rsp)
			return
		}
		ctx.JSON(200, rsp)
	} else {
		// todo Invoke event.
		event := dispatch.NewQueryCurrentRegionEvent(regionData)
		// Respond with event result.
		ctx.String(200, event.RegionInfo)
	}
}
