package main

import (
    "fmt"
    "net"
    "time"
    "mynotes/src/net/rtp"
)

var (
    localPort = 5220
    local, _ = net.ResolveIPAddr("ip", "192.168.1.178")
    remotePort = 5222
    remote, _ = net.ResolveIPAddr("ip", "192.168.1.178")

    rsLocal *rtp.Session
    rsRemote *rtp.Session

    localPay [160]byte
    remotePay [160]byte

    stop bool
    stopLocalRecv chan bool
    stopRemoteRecv chan bool
    stopLocalCtrl chan bool
    stopRemoteCtrl chan bool

    eventNamesNew = []string{"NewStreamData", "NewStreamCtrl"}
    eventNamesRtcp = []string{"SR", "RR", "SDES", "BYE"}

    localZone = ""
    remoteZone = ""
)

// Create a RTP packet suitable for standard stream (index 0) with a payload length of 160 bytes
// The method initializes the RTP packet with SSRC, sequence number, and RTP version number. 
// If the payload type was set with the RTP stream then the payload type is also set in
// the RTP packet   
func sendLocalToRemote() {
    var cnt int
    stamp := uint32(0)
    for !stop {
        rp := rsLocal.NewDataPacket(stamp)
        rp.SetPayload(localPay[:])
        rsLocal.WriteData(rp)
        rp.FreePacket()
        if (cnt % 50) == 0 {
            fmt.Println("Local sent :   ", cnt)
        }
        cnt++
        stamp += 160
        time.Sleep(20e6)
    }
}

func sendLocalToRemoteIdx(index uint32) {
    var cnt int
    stamp := uint32(0)
    for !stop {
        rp := rsLocal.NewDataPacketForStream(index, stamp)
        rp.SetPayload(localPay[:])
        rsLocal.WriteData(rp)
        rp.FreePacket()
        if (cnt % 50) == 0 {
			fmt.Println("Local with index sent:   ", cnt)
        }
        cnt++
        stamp += 160
        time.Sleep(20e6)
    }
}

func sendRemoteToLocal() {
    var cnt int
    stamp := uint32(0)
    for !stop {
        rp := rsRemote.NewDataPacket(stamp)
        rp.SetPayload(remotePay[:])
        rsRemote.WriteData(rp)
        rp.FreePacket()
        if (cnt % 50) == 0 {
            fmt.Println("Remote sent:   ", cnt)
        }
        cnt++
        stamp += 160
        time.Sleep(20e6)
    }
}

func receivePacketLocal() {
    // Create and store the data receive channel.
    dataReceiver := rsLocal.CreateDataReceiveChan()
    var cnt int

    for {
        select {
        case rp := <-dataReceiver: // just get a packet - maybe we add some tests later
            if (cnt % 50) == 0 {
                fmt.Printf("Remote receiver got %d packets\n", cnt)
            }
            cnt++
            rp.FreePacket()
        case <-stopLocalRecv:
            return
        }
    }
}

func receivePacketRemote() {
    // Create and store the data receive channel.
    dataReceiver := rsRemote.CreateDataReceiveChan()
    var cnt int

    for {
        select {
        case rp := <-dataReceiver: // just get a packet - maybe we add some tests later
            if (cnt % 50) == 0 {
                fmt.Printf("Remote receiver got: %d packets\n", cnt)
            }
            cnt++
            rp.FreePacket()
        case <-stopRemoteRecv:
            return
        }
    }
}

func receiveCtrlLocal() {
    // Create and store the control event channel.
    ctrlReceiver := rsLocal.CreateCtrlEventChan()
    for {
        select {
        case evSlice := <-ctrlReceiver: // get an event
            fmt.Println("Local Length:   ", len(evSlice))
            for _, event := range evSlice {
                if event != nil {
                    var eventName string
                    if event.EventType < 200 {
                        eventName = eventNamesNew[event.EventType]
                    } else {
                        eventName = eventNamesRtcp[event.EventType-200]
                    }
                    fmt.Printf("Local: received ctrl event, type: %s, ssrc: %d, %s\n", eventName, event.Ssrc, event.Reason)
                } else {
                    fmt.Println("Local: unexpected nil event")
                }
            } 
        case <-stopLocalCtrl:
            return
        }
    }
}

func receiveCtrlRemote() {
    // Create and store the control event channel.
    ctrlReceiver := rsRemote.CreateCtrlEventChan()
    for {
        select {
        case evSlice := <-ctrlReceiver: // get an event
            fmt.Println("Remote: Length of event slice:", len(evSlice))
            for _, event := range evSlice {
                if event != nil {
                    var eventName string
                    if event.EventType < 200 {
                        eventName = eventNamesNew[event.EventType]
                    } else {
                        eventName = eventNamesRtcp[event.EventType-200]
                    }
                    fmt.Printf("Remote: received ctrl event, type: %s, ssrc: %d, %s\n", eventName, event.Ssrc, event.Reason)
                } else {
                    fmt.Println("Remote: unexpected nil event")
                }
            } 
        case <-stopRemoteCtrl:
            return
        }
    }
}

func initialize() {
    // Some initialization for payload byte arrays
    for i := range localPay {
        localPay[i] = byte(i)
    }
    for i := range remotePay {
        remotePay[i] = byte(len(remotePay) - i)
    }
    stopLocalRecv = make(chan bool, 1)
    stopRemoteRecv = make(chan bool, 1)
    stopLocalCtrl = make(chan bool, 1)
    stopRemoteCtrl = make(chan bool, 1)
}

func fullDuplex() {
    fmt.Println("Starting full duplex test.")

    // Create a UDP transport with "local" address and use this for a "local" RTP session
    // The RTP session uses the transport to receive and send RTP packets to the remote peer.
    tpLocal, _ := rtp.NewTransportUDP(local, localPort, localZone)

    // TransportUDP implements TransportWrite and TransportRecv interfaces thus
    // use it to initialize the Session for both interfaces.
    rsLocal = rtp.NewSession(tpLocal, tpLocal)

    // Add address of a remote peer (participant)
    rsLocal.AddRemote(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone})

    // Create a media stream. 
    // The SSRC identifies the stream. Each stream has its own sequence number and other 
    // context. A RTP session can have several RTP stream for example to send several
    // streams of the same media.
    //
    strLocalIdx, _ := rsLocal.NewSsrcStreamOut(&rtp.Address{local.IP, localPort, localPort + 1, localZone}, 1020304, 4711)
    rsLocal.SsrcStreamOutForIndex(strLocalIdx).SetPayloadType(0)

    // Create the same set for a "remote" peer and use the "local" as its remote peer
    tpRemote, _ := rtp.NewTransportUDP(remote, remotePort, remoteZone)
    rsRemote = rtp.NewSession(tpRemote, tpRemote)
    rsRemote.AddRemote(&rtp.Address{local.IP, localPort, localPort + 1, localZone})

    strRemoteIdx, _ := rsRemote.NewSsrcStreamOut(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone}, 4030201, 815)
    rsRemote.SsrcStreamOutForIndex(strRemoteIdx).SetPayloadType(0)

    go receivePacketLocal()
    go receivePacketRemote()

    go receiveCtrlLocal()
    go receiveCtrlRemote()

    rsLocal.StartSession()
    rsRemote.StartSession()

    go sendLocalToRemote()
    go sendRemoteToLocal()

    time.Sleep(8e9)

    stop = true
    time.Sleep(30e6) // allow the sender to drain

    stopRemoteRecv <- true
    stopLocalRecv <- true
    stopRemoteCtrl <- true
    stopLocalCtrl <- true

    rsLocal.CloseSession()
    rsRemote.CloseSession()

    time.Sleep(10e6)

    fmt.Println("Full duplex test done.")
}

func fullDuplexTwoStreams() {
    fmt.Println("Starting full duplex test with two output streams from local to remote.")

    // Create a UDP transport with "local" address and use this for a "local" RTP session
    // The RTP session uses the transport to receive and send RTP packets to the remote peer.
    tpLocal, _ := rtp.NewTransportUDP(local, localPort, localZone)

    // TransportUDP implements TransportWrite and TransportRecv interfaces thus
    // use it to initialize the Session for both interfaces.
    rsLocal = rtp.NewSession(tpLocal, tpLocal)

    // Add address of a remote peer (participant)
    rsLocal.AddRemote(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone})

    // Create a media stream. 
    // The SSRC identifies the stream. Each stream has its own sequence number and other 
    // context. A RTP session can have several RTP stream for example to send several
    // streams of the same media.
    //
    strLocalIdx, _ := rsLocal.NewSsrcStreamOut(&rtp.Address{local.IP, localPort, localPort + 1, localZone}, 1020304, 4711)
    rsLocal.SsrcStreamOutForIndex(strLocalIdx).SetPayloadType(0)

    // create a second output stream
    strLocalIdx, _ = rsLocal.NewSsrcStreamOut(&rtp.Address{local.IP, localPort, localPort + 1, localZone}, 11223344, 1234)
    rsLocal.SsrcStreamOutForIndex(strLocalIdx).SetPayloadType(0)

    // Create the same set for a "remote" peer and use the "local" as its remote peer. Remote peer has one output stream only.
    tpRemote, _ := rtp.NewTransportUDP(remote, remotePort, remoteZone)
    rsRemote = rtp.NewSession(tpRemote, tpRemote)
    rsRemote.AddRemote(&rtp.Address{local.IP, localPort, localPort + 1, localZone})

    strRemoteIdx, _ := rsRemote.NewSsrcStreamOut(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone}, 4030201, 815)
    rsRemote.SsrcStreamOutForIndex(strRemoteIdx).SetPayloadType(0)

    go receivePacketLocal()
    go receivePacketRemote()

    go receiveCtrlLocal()
    go receiveCtrlRemote()

    rsLocal.StartSession()
    rsRemote.StartSession()

    go sendLocalToRemote()
    go sendLocalToRemoteIdx(strLocalIdx)
    go sendRemoteToLocal()

    time.Sleep(8e9)

    stop = true
    time.Sleep(30e6) // allow  the sender to drain

    stopRemoteRecv <- true
    stopLocalRecv <- true
    stopRemoteCtrl <- true
    stopLocalCtrl <- true

    rsLocal.CloseSession()
    rsRemote.CloseSession()

    time.Sleep(10e6)

    fmt.Printf("Full duplex test with 2 output streams done.")
}

func simpleRtp() {
    fmt.Println("Starting simple RTP test.")
    // Create a UDP transport with "local" address and use this for a "local" RTP session
    // The RTP session uses the transport to receive and send RTP packets to the remote peer.
    tpLocal, _ := rtp.NewTransportUDP(local, localPort, localZone)
    // TransportUDP implements TransportWrite and TransportRecv interfaces thus
    // use it to initialize the Session for both interfaces.
    rsLocal = rtp.NewSession(tpLocal, tpLocal)
    // Add address of a remote peer (participant)
    rsLocal.AddRemote(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone})
    // Create a media stream. 
    // The SSRC identifies the stream. Each stream has its own sequence number and other 
    // context. A RTP session can have several RTP stream for example to send several
    // streams of the same media.
    //
    strLocalIdx, _ := rsLocal.NewSsrcStreamOut(&rtp.Address{local.IP, localPort, localPort + 1, localZone}, 1020304, 4711)
    rsLocal.SsrcStreamOutForIndex(strLocalIdx).SetPayloadType(0)
    // Create the same set for a "remote" peer and use the "local" as its remote peer.
    tpRemote, _ := rtp.NewTransportUDP(remote, remotePort, remoteZone)
    rsRemote = rtp.NewSession(tpRemote, tpRemote)
    rsRemote.AddRemote(&rtp.Address{local.IP, localPort, localPort + 1, localZone})
    strRemoteIdx, _ := rsRemote.NewSsrcStreamOut(&rtp.Address{remote.IP, remotePort, remotePort + 1, remoteZone}, 4030201, 815)
    rsRemote.SsrcStreamOutForIndex(strRemoteIdx).SetPayloadType(0)
    go receivePacketLocal()
    go receivePacketRemote()
    // simple RTP: just listen on the RTP and RTCP receive transports. Do not start Session.
    rsLocal.ListenOnTransports()
    rsRemote.ListenOnTransports()
    // Just connect to control event channel, however in simple RTP mode GoRTP does not report any events.
    go sendLocalToRemote()
    go sendRemoteToLocal()
    time.Sleep(8e9)
    stop = true
    time.Sleep(30e6) // allow the sender to drain
    stopRemoteRecv <- true
    stopLocalRecv <- true
    stopRemoteCtrl <- true
    stopLocalCtrl <- true
    // Just close the receivers, no need to close a session.
    rsLocal.CloseRecv()
    rsRemote.CloseRecv()
    time.Sleep(10e6)
    fmt.Printf("Simple RTP test done.")
}

func main() {
    initialize()
    fullDuplex()
//    fullDuplexTwoStreams()
//    simpleRtp()
}
