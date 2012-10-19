package graph

type Cost interface {
	Add(Cost) Cost
	Less(Cost) bool
}

type CostI int

func (a CostI) Add(b Cost) Cost { return a + b.(CostI) }
func (a CostI) Less(b Cost) bool { return a < b.(CostI) }

type CostI8 int8

func (a CostI8) Add(b Cost) Cost { return a + b.(CostI8) }
func (a CostI8) Less(b Cost) bool { return a < b.(CostI8) }

type CostI16 int16

func (a CostI16) Add(b Cost) Cost { return a + b.(CostI16) }
func (a CostI16) Less(b Cost) bool { return a < b.(CostI16) }

type CostI32 int32

func (a CostI32) Add(b Cost) Cost { return a + b.(CostI32) }
func (a CostI32) Less(b Cost) bool { return a < b.(CostI32) }

type CostI64 int64

func (a CostI64) Add(b Cost) Cost { return a + b.(CostI64) }
func (a CostI64) Less(b Cost) bool { return a < b.(CostI64) }

type CostU uint

func (a CostU) Add(b Cost) Cost { return a + b.(CostU) }
func (a CostU) Less(b Cost) bool { return a < b.(CostU) }

type CostU8 uint8

func (a CostU8) Add(b Cost) Cost { return a + b.(CostU8) }
func (a CostU8) Less(b Cost) bool { return a < b.(CostU8) }

type CostU16 uint16

func (a CostU16) Add(b Cost) Cost { return a + b.(CostU16) }
func (a CostU16) Less(b Cost) bool { return a < b.(CostU16) }

type CostU32 uint32

func (a CostU32) Add(b Cost) Cost { return a + b.(CostU32) }
func (a CostU32) Less(b Cost) bool { return a < b.(CostU32) }

type CostU64 uint64

func (a CostU64) Add(b Cost) Cost { return a + b.(CostU64) }
func (a CostU64) Less(b Cost) bool { return a < b.(CostU64) }

type CostF64 float64

func (a CostF64) Add(b Cost) Cost { return a + b.(CostF64) }
func (a CostF64) Less(b Cost) bool { return a < b.(CostF64) }

type CostF32 float32

func (a CostF32) Add(b Cost) Cost { return a + b.(CostF32) }
func (a CostF32) Less(b Cost) bool { return a < b.(CostF32) }

// NewCost() returns a new Cost interface
func NewCost(i interface{}) Cost {
	switch i.(type) {
	default:
		return i.(Cost)
	case int:
		return CostI(i.(int))
	case int8:
		return CostI8(i.(int8))
	case int16:           
		return CostI16(i.(int16))
	case int32:           
		return CostI32(i.(int32))
	case int64:           
		return CostI64(i.(int64))
	case uint:            
		return CostU(i.(uint))
	case uint8:           
		return CostU8(i.(uint8))
	case uint16:          
		return CostU16(i.(uint16))
	case uint32:          
		return CostU32(i.(uint32))
	case uint64:          
		return CostU64(i.(uint64))
	case float32:         
		return CostF32(i.(float32))
	case float64:         
		return CostF64(i.(float64))
	}
	return nil		// Never get here.
}
