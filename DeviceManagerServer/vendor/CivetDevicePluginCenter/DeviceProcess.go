// Package CivetDevicePluginCenter comment
// This file was generated by tars2go 1.1.4
// Generated from DeviceProcess.tars
package CivetDevicePluginCenter

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// Device struct implement
type Device struct {
	ID            int32  `json:"ID"`
	ProductID     int32  `json:"ProductID"`
	DeivceClassID int32  `json:"DeivceClassID"`
	DeivceID      string `json:"DeivceID"`
	Online        int32  `json:"online"`
}

func (st *Device) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *Device) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.ID, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ProductID, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.DeivceClassID, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.DeivceID, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Online, 4, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *Device) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Device, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *Device) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.ID, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ProductID, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.DeivceClassID, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.DeivceID, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Online, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *Device) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}