package httpGo

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Container) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "objs":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Objs == nil && msz > 0 {
				z.Objs = make(map[string]bool, msz)
			} else if len(z.Objs) > 0 {
				for key, _ := range z.Objs {
					delete(z.Objs, key)
				}
			}
			for msz > 0 {
				msz--
				var xvk string
				var bzg bool
				xvk, err = dc.ReadString()
				if err != nil {
					return
				}
				bzg, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Objs[xvk] = bzg
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Container) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "objs"
	err = en.Append(0xa4, 0x6f, 0x62, 0x6a, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Objs)))
	if err != nil {
		return
	}
	for xvk, bzg := range z.Objs {
		err = en.WriteString(xvk)
		if err != nil {
			return
		}
		err = en.WriteBool(bzg)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Container) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "objs"
	o = append(o, 0xa4, 0x6f, 0x62, 0x6a, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Objs)))
	for xvk, bzg := range z.Objs {
		o = msgp.AppendString(o, xvk)
		o = msgp.AppendBool(o, bzg)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Container) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "objs":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Objs == nil && msz > 0 {
				z.Objs = make(map[string]bool, msz)
			} else if len(z.Objs) > 0 {
				for key, _ := range z.Objs {
					delete(z.Objs, key)
				}
			}
			for msz > 0 {
				var xvk string
				var bzg bool
				msz--
				xvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				bzg, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					return
				}
				z.Objs[xvk] = bzg
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Container) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.MapHeaderSize
	if z.Objs != nil {
		for xvk, bzg := range z.Objs {
			_ = bzg
			s += msgp.StringPrefixSize + len(xvk) + msgp.BoolSize
		}
	}
	return
}
