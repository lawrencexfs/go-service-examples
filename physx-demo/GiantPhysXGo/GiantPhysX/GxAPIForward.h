/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#if defined(__linux__)
#define GX_LINUX 1
#else
#define GX_WINDOWS 1
#endif

#if GX_LINUX && __GNUC__ >= 4
#define GX_LINUX_EXPORT __attribute__((visibility("default")))
#else
#define GX_LINUX_EXPORT
#endif

#if GX_WINDOWS
#define GX_DLL_EXPORT __declspec(dllexport)
#define GX_DLL_IMPORT __declspec(dllimport)
#else
#define GX_DLL_EXPORT GX_LINUX_EXPORT
#define GX_DLL_IMPORT 
#endif

#if defined(GIANTPHYSX_EXPORTS) 
#define GX_API GX_DLL_EXPORT
#else
#define GX_API GX_DLL_IMPORT
#endif

#if GX_WINDOWS
#define GX_CALL_CONV __cdecl
#else
#define GX_CALL_CONV
#endif

#define GX_C_EXPORT extern "C"


namespace GiantPhysX
{
	class GxIGameObject;
	class GxIPhysics;
	class GxIScene;

	class GxVec3;

}; // namespace GiantPhysX
