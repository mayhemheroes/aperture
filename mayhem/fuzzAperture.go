package fuzzAperture

import "strconv"
import "github.com/fluxninja/aperture/pkg/filesystem"
import "github.com/fluxninja/aperture/pkg/config"
import "github.com/fluxninja/aperture/pkg/utils"
import fuzz "github.com/AdaLogics/go-fuzz-headers"

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
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }

            filesystem.PurgeDirectory(content)
            return 0

        case 2:
            test, _ := config.NewProtobufUnmarshaller(bytes)

            fuzzConsumerInterface := fuzz.NewConsumer(bytes)
            var content interface{}
            err := fuzzConsumerInterface.CreateSlice(&content)
            if err != nil {
                return 0
            }

            test.Unmarshal(content)
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
            fuzzConsumerInterface := fuzz.NewConsumer(bytes)
            var content interface{}
            err := fuzzConsumerInterface.CreateSlice(&content)
            if err != nil {
                return 0
            }

            config.UnmarshalYAML(bytes, content)
            return 0

        case 6:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumerString.CreateSlice(&content)
            if err != nil {
                return 0
            }

            fuzzConsumerArr := fuzz.NewConsumer(bytes)
            var content2 []string
            err = fuzzConsumerArr.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            utils.SliceFind(content2, content)
            return 0

        case 7:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumerString.CreateSlice(&content)
            if err != nil {
                return 0
            }

            fuzzConsumerArr := fuzz.NewConsumer(bytes)
            var content2 []string
            err = fuzzConsumerArr.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            utils.SliceContains(content2, content)
            return 0

        default:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }

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