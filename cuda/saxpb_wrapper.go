package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var saxpb_code cu.Function

type saxpb_args struct {
	arg_dst unsafe.Pointer
	arg_src unsafe.Pointer
	arg_fac float32
	arg_y   float32
	arg_N   int
	argptr  [5]unsafe.Pointer
}

// Wrapper for saxpb CUDA kernel, asynchronous.
func k_saxpb_async(dst unsafe.Pointer, src unsafe.Pointer, fac float32, y float32, N int, cfg *config, str cu.Stream) {
	if saxpb_code == 0 {
		saxpb_code = fatbinLoad(saxpb_map, "saxpb")
	}

	var _a_ saxpb_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_src = src
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_fac = fac
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_fac)
	_a_.arg_y = y
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_y)
	_a_.arg_N = N
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_N)

	args := _a_.argptr[:]
	cu.LaunchKernel(saxpb_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for saxpb CUDA kernel, synchronized.
func k_saxpb(dst unsafe.Pointer, src unsafe.Pointer, fac float32, y float32, N int, cfg *config) {
	str := stream()
	k_saxpb_async(dst, src, fac, y, N, cfg, str)
	syncAndRecycle(str)
}

var saxpb_map = map[int]string{0: "",
	20: saxpb_ptx_20,
	30: saxpb_ptx_30,
	35: saxpb_ptx_35}

const (
	saxpb_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry saxpb(
	.param .u64 saxpb_param_0,
	.param .u64 saxpb_param_1,
	.param .f32 saxpb_param_2,
	.param .f32 saxpb_param_3,
	.param .u32 saxpb_param_4
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<8>;


	ld.param.u64 	%rd3, [saxpb_param_0];
	ld.param.u64 	%rd4, [saxpb_param_1];
	ld.param.f32 	%f1, [saxpb_param_2];
	ld.param.f32 	%f2, [saxpb_param_3];
	ld.param.u32 	%r2, [saxpb_param_4];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 8 1
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.f32 	%f3, [%rd6];
	fma.rn.f32 	%f4, %f3, %f1, %f2;
	add.s64 	%rd7, %rd1, %rd5;
	st.global.f32 	[%rd7], %f4;

BB0_2:
	.loc 2 10 2
	ret;
}


`
	saxpb_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry saxpb(
	.param .u64 saxpb_param_0,
	.param .u64 saxpb_param_1,
	.param .f32 saxpb_param_2,
	.param .f32 saxpb_param_3,
	.param .u32 saxpb_param_4
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<8>;


	ld.param.u64 	%rd3, [saxpb_param_0];
	ld.param.u64 	%rd4, [saxpb_param_1];
	ld.param.f32 	%f1, [saxpb_param_2];
	ld.param.f32 	%f2, [saxpb_param_3];
	ld.param.u32 	%r2, [saxpb_param_4];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 8 1
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.f32 	%f3, [%rd6];
	fma.rn.f32 	%f4, %f3, %f1, %f2;
	add.s64 	%rd7, %rd1, %rd5;
	st.global.f32 	[%rd7], %f4;

BB0_2:
	.loc 2 10 2
	ret;
}


`
	saxpb_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry saxpb(
	.param .u64 saxpb_param_0,
	.param .u64 saxpb_param_1,
	.param .f32 saxpb_param_2,
	.param .f32 saxpb_param_3,
	.param .u32 saxpb_param_4
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<10>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<8>;


	ld.param.u64 	%rd3, [saxpb_param_0];
	ld.param.u64 	%rd4, [saxpb_param_1];
	ld.param.f32 	%f1, [saxpb_param_2];
	ld.param.f32 	%f2, [saxpb_param_3];
	ld.param.u32 	%r2, [saxpb_param_4];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 3 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 3 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB2_2;

	.loc 3 8 1
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	fma.rn.f32 	%f4, %f3, %f1, %f2;
	add.s64 	%rd7, %rd1, %rd5;
	st.global.f32 	[%rd7], %f4;

BB2_2:
	.loc 3 10 2
	ret;
}


`
)