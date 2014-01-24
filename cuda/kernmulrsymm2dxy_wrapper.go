package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for kernmulRSymm2Dxy kernel
var kernmulRSymm2Dxy_code cu.Function

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
type kernmulRSymm2Dxy_args_t struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	argptr     [7]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
var kernmulRSymm2Dxy_args kernmulRSymm2Dxy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm2Dxy_args.argptr[0] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMx)
	kernmulRSymm2Dxy_args.argptr[1] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMy)
	kernmulRSymm2Dxy_args.argptr[2] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxx)
	kernmulRSymm2Dxy_args.argptr[3] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKyy)
	kernmulRSymm2Dxy_args.argptr[4] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxy)
	kernmulRSymm2Dxy_args.argptr[5] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Nx)
	kernmulRSymm2Dxy_args.argptr[6] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Ny)
}

// Wrapper for kernmulRSymm2Dxy CUDA kernel, asynchronous.
func k_kernmulRSymm2Dxy_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKxy unsafe.Pointer, Nx int, Ny int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	kernmulRSymm2Dxy_args.Lock()
	defer kernmulRSymm2Dxy_args.Unlock()

	if kernmulRSymm2Dxy_code == 0 {
		kernmulRSymm2Dxy_code = fatbinLoad(kernmulRSymm2Dxy_map, "kernmulRSymm2Dxy")
	}

	kernmulRSymm2Dxy_args.arg_fftMx = fftMx
	kernmulRSymm2Dxy_args.arg_fftMy = fftMy
	kernmulRSymm2Dxy_args.arg_fftKxx = fftKxx
	kernmulRSymm2Dxy_args.arg_fftKyy = fftKyy
	kernmulRSymm2Dxy_args.arg_fftKxy = fftKxy
	kernmulRSymm2Dxy_args.arg_Nx = Nx
	kernmulRSymm2Dxy_args.arg_Ny = Ny

	args := kernmulRSymm2Dxy_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dxy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for kernmulRSymm2Dxy kernel.
var kernmulRSymm2Dxy_map = map[int]string{0: "",
	20: kernmulRSymm2Dxy_ptx_20,
	30: kernmulRSymm2Dxy_ptx_30,
	35: kernmulRSymm2Dxy_ptx_35}

// kernmulRSymm2Dxy PTX code for various compute capabilities.
const (
	kernmulRSymm2Dxy_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<43>;
	.reg .f32 	%f<20>;
	.reg .s64 	%rd<25>;


	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd11, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r6, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r7, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	.loc 2 29 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r8, %r9, %r10;
	.loc 2 30 1
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	.loc 2 32 1
	setp.ge.s32 	%p1, %r2, %r7;
	setp.ge.s32 	%p2, %r1, %r6;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_5;

	.loc 2 37 1
	mad.lo.s32 	%r42, %r2, %r6, %r1;
	.loc 2 38 1
	shl.b32 	%r15, %r42, 1;
	.loc 2 40 1
	mul.wide.s32 	%rd12, %r15, 4;
	add.s64 	%rd3, %rd2, %rd12;
	ld.global.f32 	%f1, [%rd3];
	.loc 2 41 1
	add.s32 	%r17, %r15, 1;
	mul.wide.s32 	%rd13, %r17, 4;
	add.s64 	%rd4, %rd2, %rd13;
	ld.global.f32 	%f2, [%rd4];
	.loc 2 42 1
	add.s64 	%rd5, %rd1, %rd12;
	ld.global.f32 	%f3, [%rd5];
	.loc 2 43 1
	add.s64 	%rd6, %rd1, %rd13;
	ld.global.f32 	%f4, [%rd6];
	.loc 2 46 1
	shr.u32 	%r21, %r7, 31;
	add.s32 	%r22, %r7, %r21;
	shr.s32 	%r23, %r22, 1;
	add.s32 	%r24, %r23, 1;
	setp.lt.s32 	%p4, %r2, %r24;
	@%p4 bra 	BB0_3;

	.loc 2 51 1
	sub.s32 	%r25, %r7, %r2;
	mad.lo.s32 	%r42, %r25, %r6, %r1;
	cvta.to.global.u64 	%rd14, %rd9;
	.loc 2 54 1
	mul.wide.s32 	%rd15, %r42, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f8, [%rd16];
	neg.f32 	%f19, %f8;
	bra.uni 	BB0_4;

BB0_3:
	cvta.to.global.u64 	%rd17, %rd9;
	.loc 2 49 1
	mul.wide.s32 	%rd18, %r42, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.f32 	%f19, [%rd19];

BB0_4:
	cvta.to.global.u64 	%rd20, %rd7;
	.loc 2 47 1
	mul.wide.s32 	%rd21, %r42, 4;
	add.s64 	%rd22, %rd20, %rd21;
	cvta.to.global.u64 	%rd23, %rd8;
	.loc 2 48 1
	add.s64 	%rd24, %rd23, %rd21;
	.loc 2 57 1
	ld.global.f32 	%f9, [%rd24];
	ld.global.f32 	%f10, [%rd22];
	mul.f32 	%f11, %f3, %f19;
	fma.rn.f32 	%f12, %f1, %f10, %f11;
	st.global.f32 	[%rd3], %f12;
	.loc 2 58 1
	mul.f32 	%f13, %f4, %f19;
	fma.rn.f32 	%f14, %f2, %f10, %f13;
	st.global.f32 	[%rd4], %f14;
	.loc 2 59 1
	mul.f32 	%f15, %f3, %f9;
	fma.rn.f32 	%f16, %f1, %f19, %f15;
	st.global.f32 	[%rd5], %f16;
	.loc 2 60 1
	mul.f32 	%f17, %f4, %f9;
	fma.rn.f32 	%f18, %f2, %f19, %f17;
	st.global.f32 	[%rd6], %f18;

BB0_5:
	.loc 2 61 2
	ret;
}


`
	kernmulRSymm2Dxy_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<34>;
	.reg .f32 	%f<20>;
	.reg .s64 	%rd<28>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r7, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r8, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd6;
	.loc 2 29 1
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 2 30 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 2 32 1
	setp.ge.s32 	%p1, %r2, %r8;
	setp.ge.s32 	%p2, %r1, %r7;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_5;

	.loc 2 37 1
	mad.lo.s32 	%r33, %r2, %r7, %r1;
	.loc 2 38 1
	shl.b32 	%r4, %r33, 1;
	.loc 2 40 1
	mul.wide.s32 	%rd11, %r4, 4;
	add.s64 	%rd12, %rd2, %rd11;
	ld.global.f32 	%f1, [%rd12];
	.loc 2 41 1
	add.s32 	%r16, %r4, 1;
	mul.wide.s32 	%rd13, %r16, 4;
	add.s64 	%rd3, %rd2, %rd13;
	ld.global.f32 	%f2, [%rd3];
	.loc 2 42 1
	add.s64 	%rd4, %rd1, %rd11;
	ld.global.f32 	%f3, [%rd4];
	.loc 2 43 1
	add.s64 	%rd5, %rd1, %rd13;
	ld.global.f32 	%f4, [%rd5];
	.loc 2 46 1
	shr.u32 	%r20, %r8, 31;
	add.s32 	%r21, %r8, %r20;
	shr.s32 	%r22, %r21, 1;
	add.s32 	%r23, %r22, 1;
	setp.lt.s32 	%p4, %r2, %r23;
	@%p4 bra 	BB0_3;

	.loc 2 51 1
	sub.s32 	%r24, %r8, %r2;
	mad.lo.s32 	%r33, %r24, %r7, %r1;
	cvta.to.global.u64 	%rd14, %rd9;
	.loc 2 54 1
	mul.wide.s32 	%rd15, %r33, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f8, [%rd16];
	neg.f32 	%f19, %f8;
	bra.uni 	BB0_4;

BB0_3:
	cvta.to.global.u64 	%rd17, %rd9;
	.loc 2 49 1
	mul.wide.s32 	%rd18, %r33, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.f32 	%f19, [%rd19];

BB0_4:
	cvta.to.global.u64 	%rd20, %rd7;
	.loc 2 47 1
	mul.wide.s32 	%rd21, %r33, 4;
	add.s64 	%rd22, %rd20, %rd21;
	cvta.to.global.u64 	%rd23, %rd8;
	.loc 2 48 1
	add.s64 	%rd24, %rd23, %rd21;
	.loc 2 57 1
	ld.global.f32 	%f9, [%rd24];
	ld.global.f32 	%f10, [%rd22];
	mul.f32 	%f11, %f3, %f19;
	fma.rn.f32 	%f12, %f1, %f10, %f11;
	.loc 2 40 1
	mul.wide.s32 	%rd26, %r4, 4;
	add.s64 	%rd27, %rd2, %rd26;
	.loc 2 57 1
	st.global.f32 	[%rd27], %f12;
	.loc 2 58 1
	mul.f32 	%f13, %f4, %f19;
	fma.rn.f32 	%f14, %f2, %f10, %f13;
	st.global.f32 	[%rd3], %f14;
	.loc 2 59 1
	mul.f32 	%f15, %f3, %f9;
	fma.rn.f32 	%f16, %f1, %f19, %f15;
	st.global.f32 	[%rd4], %f16;
	.loc 2 60 1
	mul.f32 	%f17, %f4, %f9;
	fma.rn.f32 	%f18, %f2, %f19, %f17;
	st.global.f32 	[%rd5], %f18;

BB0_5:
	.loc 2 61 2
	ret;
}


`
	kernmulRSymm2Dxy_ptx_35 = `
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

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<30>;
	.reg .f32 	%f<20>;
	.reg .s64 	%rd<28>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r7, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r8, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd6;
	.loc 3 29 1
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 3 30 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 3 32 1
	setp.ge.s32 	%p1, %r2, %r8;
	setp.ge.s32 	%p2, %r1, %r7;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB2_5;

	.loc 3 37 1
	mad.lo.s32 	%r29, %r2, %r7, %r1;
	.loc 3 38 1
	shl.b32 	%r4, %r29, 1;
	.loc 3 40 1
	mul.wide.s32 	%rd11, %r4, 4;
	add.s64 	%rd12, %rd2, %rd11;
	ld.global.f32 	%f1, [%rd12];
	.loc 3 41 1
	add.s32 	%r16, %r4, 1;
	mul.wide.s32 	%rd13, %r16, 4;
	add.s64 	%rd3, %rd2, %rd13;
	ld.global.f32 	%f2, [%rd3];
	.loc 3 42 1
	add.s64 	%rd4, %rd1, %rd11;
	ld.global.f32 	%f3, [%rd4];
	.loc 3 43 1
	add.s64 	%rd5, %rd1, %rd13;
	ld.global.f32 	%f4, [%rd5];
	.loc 3 46 1
	shr.u32 	%r20, %r8, 31;
	add.s32 	%r21, %r8, %r20;
	shr.s32 	%r22, %r21, 1;
	add.s32 	%r23, %r22, 1;
	setp.lt.s32 	%p4, %r2, %r23;
	@%p4 bra 	BB2_3;

	.loc 3 51 1
	sub.s32 	%r24, %r8, %r2;
	mad.lo.s32 	%r29, %r24, %r7, %r1;
	cvta.to.global.u64 	%rd14, %rd9;
	.loc 3 54 1
	mul.wide.s32 	%rd15, %r29, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.nc.f32 	%f8, [%rd16];
	neg.f32 	%f19, %f8;
	bra.uni 	BB2_4;

BB2_3:
	cvta.to.global.u64 	%rd17, %rd9;
	.loc 3 49 1
	mul.wide.s32 	%rd18, %r29, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.nc.f32 	%f19, [%rd19];

BB2_4:
	cvta.to.global.u64 	%rd20, %rd7;
	.loc 3 47 1
	mul.wide.s32 	%rd21, %r29, 4;
	add.s64 	%rd22, %rd20, %rd21;
	cvta.to.global.u64 	%rd23, %rd8;
	.loc 3 48 1
	add.s64 	%rd24, %rd23, %rd21;
	.loc 3 57 1
	ld.global.nc.f32 	%f9, [%rd22];
	ld.global.nc.f32 	%f10, [%rd24];
	mul.f32 	%f11, %f3, %f19;
	fma.rn.f32 	%f12, %f1, %f9, %f11;
	.loc 3 40 1
	mul.wide.s32 	%rd26, %r4, 4;
	add.s64 	%rd27, %rd2, %rd26;
	.loc 3 57 1
	st.global.f32 	[%rd27], %f12;
	.loc 3 58 1
	mul.f32 	%f13, %f4, %f19;
	fma.rn.f32 	%f14, %f2, %f9, %f13;
	st.global.f32 	[%rd3], %f14;
	.loc 3 59 1
	mul.f32 	%f15, %f3, %f10;
	fma.rn.f32 	%f16, %f1, %f19, %f15;
	st.global.f32 	[%rd4], %f16;
	.loc 3 60 1
	mul.f32 	%f17, %f4, %f10;
	fma.rn.f32 	%f18, %f2, %f19, %f17;
	st.global.f32 	[%rd5], %f18;

BB2_5:
	.loc 3 61 2
	ret;
}


`
)