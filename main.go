package main

import (
    "github.com/firedrake969/gopher-brain"
    "github.com/googollee/go-socket.io"

    "log"
    "strconv"
    "net/http"
)

const MAIN_ROOM = "server"

func emitToAll(so socketio.Socket, event string, data ...interface{}) {
    so.BroadcastTo(MAIN_ROOM, event, data...)
    so.Emit(event, data...)
}

// this loop is more suitable for sensors
// outputs := [](map[string]string){}
//             outputsUsed := make(map[string]bool)
//             for name, _ := range myNet.Outputs {
//                 baseName := name[0:len(name)-2]
//                 if _, exists := outputsUsed[baseName]; !exists {
//                     outputs = append(outputs, map[string]string{
//                         "name": baseName,
//                         "valueOne": strconv.FormatFloat(myNet.Outputs[fmt.Sprintf("%v-0", baseName)].Value, 'f', -1, 64),
//                         "valueTwo": strconv.FormatFloat(myNet.Outputs[fmt.Sprintf("%v-1", baseName)].Value, 'f', -1, 64),
//                     })
//                     outputsUsed[baseName] = true
//                 }
//             }
//             jsonRep, _ := json.Marshal(outputs)

func main() {
    frames := 0 // TODO - move this into a property on the Network struct
    myNet := brain.Brain([3]int{12, 25, 25}, []brain.SensorConstructor{
        brain.SensorConstructor{
            Name:"eye",
            R: 1,
            Count: 9,
            Plane: "y",
            Center: [3]int{8, 0, 12},
            OutputCount: 5,
            InputFunc: func(nodes []*brain.Node, influences map[string]*brain.Output) {
                for _, node := range nodes {
                    node.Value = 1
                }
            },
        },
        brain.SensorConstructor{
            Name:"ear",
            R: 1,
            Count: 9,
            Plane: "y",
            Center: [3]int{8, 8, 12},
            OutputCount: 3,
            InputFunc: func(nodes []*brain.Node, influences map[string]*brain.Output) {
                for _, node := range nodes {
                    node.Value = 1
                }
            },
        },
    })

    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        so.Join(MAIN_ROOM)
        so.Emit("connected", true)
        so.Emit("cycle", frames)

        // jsonRep, _ := json.Marshal(outputs)
        so.Emit("outputs", SerializeOutputs(myNet))

        // todo - function in gopher-brain pkg to allow conversion of net to fully JSON-able state
        // sensors := make(map[string]float64)
        // for name, sensor := range myNet.Sensors {
        //     sensors[name] = sensor.Value
        // }
        // jsonRep, _ := json.Marshal(sensors)
        // so.Emit("sensors", string(jsonRep))
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })
    server.On("cycle", func(so socketio.Socket, cycles int, saveFrames bool) {
        // should I save before or after cycle?
        for i := 0; i < cycles; i++ {
            myNet.Cycle()
            if saveFrames {
                myNet.DumpJSON(strconv.Itoa(frames))
            }
            frames++

            emitToAll(so, "cycle", frames)

            outputs := SerializeOutputs(myNet)
            emitToAll(so, "outputs", outputs)
        }
    })
    server.On("save", func(so socketio.Socket, saveName string) {
        myNet.SaveState(saveName)
        emitToAll(so, "saved")
    })
    server.On("autorun", func(so socketio.Socket, autorun bool) {
        emitToAll(so, "autorun", autorun)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./client/build")))
    log.Println("-- Serving localhost:3000 --")
    log.Fatal(http.ListenAndServe(":3000", nil))
}