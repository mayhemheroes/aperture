package fuzz

import "strconv"
import "github.com/fluxninja/aperture/pkg/filesystem"
import "github.com/fluxninja/aperture/pkg/config"
import "github.com/fluxninja/aperture/pkg/utils"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
            var test filesystem.FileInfo
            test.WriteByteBufferToFile(bytes)
            return 0

        case 1:
            content := string(bytes)
            filesystem.PurgeDirectory(content)
            return 0

        case 2:
            test, _ := config.NewProtobufUnmarshaller(bytes)
            var out interface{}
            test.Unmarshal(out)
            return 0

        case 3:
            var test config.ProtobufUnmarshaller
            test.Reload(bytes)
            return 0

        case 4:
            var test config.Duration
            test.UnmarshalJSON(bytes)
            return 0

        case 5:
            var out interface{}
            config.UnmarshalYAML(bytes, out)
            return 0

        // case 6:
        //     content := string(bytes)
        //     var test config.KoanfUnmarshaller
        //     test.Get(content)
        //     return 0

        // case 7:
        //     content := string(bytes)
        //     var test config.KoanfUnmarshaller
        //     test.IsSet(content)
        //     return 0

        case 8:
            content := string(bytes)
            var arr = []string{"mayhem", "fuzz"}
            utils.SliceFind(arr, content)
            return 0

        case 9:
            content := string(bytes)
            var arr = []string{"mayhem", "fuzz"}
            utils.SliceContains(arr, content)
            return 0

        default:
            content := string(bytes)
            utils.IsHTTPUrl(content)
            return 0

        }

    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}