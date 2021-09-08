package main

func fsum(output *[]int64, in *[]int64) {
	for i := 0; i < 10; i += 2 {
		(*output)[i] += (*in)[i]
		(*output)[i+1] += (*in)[i+1]
	}
}
func fdifference(output *[]int64, in *[]int64) {
	for i := 0; i < 10; i++ {
		(*output)[i] = (*in)[i] - (*output)[i]
	}
}

func fscalar_product(output *[]int64, in *[]int64, scalar int64) {
	for i := 0; i < 10; i++ {
		(*output)[i] = (*in)[i] * scalar
	}
}

func fproduct(output *[]int64, in2 *[]int64, in *[]int64) {
	(*output)[0] = int64(int32((*in2)[0]) * int32((*in)[0]))

	(*output)[1] = int64(int32((*in2)[0])*int32((*in)[1])) +
		int64(int32((*in2)[1])*int32((*in)[0]))

	(*output)[2] = 2*int64(int32((*in2)[1])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[2])) +
		int64(int32((*in2)[2])*int32((*in)[0]))

	(*output)[3] = int64(int32((*in2)[1])*int32((*in)[2])) +
		int64(int32((*in2)[2])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[3])) +
		int64(int32((*in2)[3])*int32((*in)[0]))

	(*output)[4] = int64(int32((*in2)[2])*int32((*in)[2])) +
		2*int64(int32((*in2)[1])*int32((*in)[3])) +
		int64(int32((*in2)[3])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[4]))

	(*output)[5] = int64(int32((*in2)[2])*int32((*in)[3])) +
		int64(int32((*in2)[3])*int32((*in)[2])) +
		int64(int32((*in2)[1])*int32((*in)[4])) +
		int64(int32((*in2)[4])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[5])) +
		int64(int32((*in2)[5])*int32((*in)[0]))

	(*output)[6] = 2*int64(int32((*in2)[3])*int32((*in)[3])) +
		int64(int32((*in2)[1])*int32((*in)[5])) +
		int64(int32((*in2)[5])*int32((*in)[1])) +
		int64(int32((*in2)[2])*int32((*in)[4])) +
		int64(int32((*in2)[4])*int32((*in)[2])) +
		int64(int32((*in2)[0])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[0]))

	(*output)[7] = int64(int32((*in2)[3])*int32((*in)[4])) +
		int64(int32((*in2)[4])*int32((*in)[3])) +
		int64(int32((*in2)[2])*int32((*in)[5])) +
		int64(int32((*in2)[5])*int32((*in)[2])) +
		int64(int32((*in2)[1])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[1]))

	(*output)[8] = int64(int32((*in2)[4])*int32((*in)[4])) +
		2*int64(int32((*in2)[3])*int32((*in)[5])) +
		int64(int32((*in2)[5])*int32((*in)[3])) +
		int64(int32((*in2)[1])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[1])) +
		int64(int32((*in2)[2])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[2])) +
		int64(int32((*in2)[0])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[0]))

	(*output)[9] = int64(int32((*in2)[4])*int32((*in)[5])) +
		int64(int32((*in2)[5])*int32((*in)[4])) +
		int64(int32((*in2)[3])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[3])) +
		int64(int32((*in2)[2])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[2])) +
		int64(int32((*in2)[1])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[1])) +
		int64(int32((*in2)[0])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[0]))

	(*output)[10] = 2*int64(int32((*in2)[5])*int32((*in)[5])) +
		int64(int32((*in2)[3])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[3])) +
		int64(int32((*in2)[1])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[1])) +
		int64(int32((*in2)[4])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[4])) +
		int64(int32((*in2)[2])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[2]))

	(*output)[11] = int64(int32((*in2)[5])*int32((*in)[6])) +
		int64(int32((*in2)[6])*int32((*in)[5])) +
		int64(int32((*in2)[4])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[4])) +
		int64(int32((*in2)[3])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[3])) +
		int64(int32((*in2)[2])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[2]))

	(*output)[12] = int64(int32((*in2)[6])*int32((*in)[6])) +
		2*int64(int32((*in2)[5])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[5])) +
		int64(int32((*in2)[3])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[3])) +
		int64(int32((*in2)[4])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[4]))

	(*output)[13] = int64(int32((*in2)[6])*int32((*in)[7])) +
		int64(int32((*in2)[7])*int32((*in)[6])) +
		int64(int32((*in2)[5])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[5])) +
		int64(int32((*in2)[4])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[4]))

	(*output)[14] = 2*int64(int32((*in2)[7])*int32((*in)[7])) +
		int64(int32((*in2)[5])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[5])) +
		int64(int32((*in2)[6])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[6]))

	(*output)[15] = int64(int32((*in2)[7])*int32((*in)[8])) +
		int64(int32((*in2)[8])*int32((*in)[7])) +
		int64(int32((*in2)[6])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[6]))

	(*output)[16] = int64(int32((*in2)[8])*int32((*in)[8])) +
		2*int64(int32((*in2)[7])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[7]))

	(*output)[17] = int64(int32((*in2)[8])*int32((*in)[9])) +
		int64(int32((*in2)[9])*int32((*in)[8]))

	(*output)[18] = 2 * int64(int32((*in2)[9])*int32((*in)[9]))
}
