// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package etl

import conv "github.com/cstockton/go-conv"

// TransformMultiplyByFloat32 :
func TransformMultiplyByFloat32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float32(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByFloat32 :
func TransformDivideByFloat32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float32(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByFloat32 :
func TransformAddByFloat32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float32(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByFloat32 :
func TransformSubtractByFloat32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float32(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByFloat64 :
func TransformMultiplyByFloat64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float64(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByFloat64 :
func TransformDivideByFloat64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float64(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByFloat64 :
func TransformAddByFloat64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float64(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByFloat64 :
func TransformSubtractByFloat64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Float64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Float64(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByInt :
func TransformMultiplyByInt(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByInt :
func TransformDivideByInt(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByInt :
func TransformAddByInt(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByInt :
func TransformSubtractByInt(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByInt8 :
func TransformMultiplyByInt8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int8(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByInt8 :
func TransformDivideByInt8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int8(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByInt8 :
func TransformAddByInt8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int8(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByInt8 :
func TransformSubtractByInt8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int8(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByInt16 :
func TransformMultiplyByInt16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int16(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByInt16 :
func TransformDivideByInt16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int16(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByInt16 :
func TransformAddByInt16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int16(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByInt16 :
func TransformSubtractByInt16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int16(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByInt32 :
func TransformMultiplyByInt32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int32(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByInt32 :
func TransformDivideByInt32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int32(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByInt32 :
func TransformAddByInt32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int32(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByInt32 :
func TransformSubtractByInt32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int32(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByInt64 :
func TransformMultiplyByInt64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int64(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByInt64 :
func TransformDivideByInt64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int64(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByInt64 :
func TransformAddByInt64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int64(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByInt64 :
func TransformSubtractByInt64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Int64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Int64(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByUint :
func TransformMultiplyByUint(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByUint :
func TransformDivideByUint(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByUint :
func TransformAddByUint(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByUint :
func TransformSubtractByUint(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByUint8 :
func TransformMultiplyByUint8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint8(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByUint8 :
func TransformDivideByUint8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint8(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByUint8 :
func TransformAddByUint8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint8(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByUint8 :
func TransformSubtractByUint8(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint8(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint8(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByUint16 :
func TransformMultiplyByUint16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint16(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByUint16 :
func TransformDivideByUint16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint16(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByUint16 :
func TransformAddByUint16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint16(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByUint16 :
func TransformSubtractByUint16(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint16(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint16(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByUint32 :
func TransformMultiplyByUint32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint32(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByUint32 :
func TransformDivideByUint32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint32(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByUint32 :
func TransformAddByUint32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint32(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByUint32 :
func TransformSubtractByUint32(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint32(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint32(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}

// TransformMultiplyByUint64 :
func TransformMultiplyByUint64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint64(value)
		if err != nil {
			return nil, err
		}
		return left * number, nil
	}
}

// TransformDivideByUint64 :
func TransformDivideByUint64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint64(value)
		if err != nil {
			return nil, err
		}
		return left / number, nil
	}
}

// TransformAddByUint64 :
func TransformAddByUint64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint64(value)
		if err != nil {
			return nil, err
		}
		return left + number, nil
	}
}

// TransformSubtractByUint64 :
func TransformSubtractByUint64(right interface{}) func(interface{}) (interface{}, error) {
	number, err := conv.DefaultConv.Uint64(right)
	return func(value interface{}) (interface{}, error) {
		if err != nil {
			return nil, err
		}
		left, err := conv.DefaultConv.Uint64(value)
		if err != nil {
			return nil, err
		}
		return left - number, nil
	}
}