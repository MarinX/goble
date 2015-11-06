#Bluetooth Low Energy for Go (based od CC2541)

#Description
Bluetooth Low Enery (HM10) module for Go

#Installation
    go get github.com/MarinX/goble

#Example
    See the example directory

#Enable BeaconMode
    hm.FactoryReset()
	hm.SetBeaconMode(goble.BEACON_ENABLE)
	hm.GetBeaconMode()
	hm.SetAdvertasingInterval(5)
	hm.Reset()
<img src=https://raw.github.com/MarinX/goble/master/beacon.png 
width=300 />


#License
This library is under the MIT License
#Author
Marin Basic 
