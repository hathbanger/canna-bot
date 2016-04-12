package models

import ()

type Sensor interface{
    TypeString() string
}

type Sensor_Base struct {
    Type        string  `json:"type"`
}

type PSISensor struct {
    Sensor_Base
}

// PSISensor Constructor
func NewPSISensor() *PSISensor {
    ret := &PSISensor{}
    ret.Type = "psi_sensor"

    return ret
}

// Method of the PSISensor, like a method of an object.
// 'psis PSISensor' to points the functions to the PSISensor Struct
func (psis PSISensor) TypeString() string {
    return psis.Type
}

// VVV More Types of Sensors VVV




