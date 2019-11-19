/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"

#include "GiantPhysX/GxIGameObject.h"
#include "GiantPhysX/GxIPhysics.h"
#include "GiantPhysX/GxIScene.h"

#include "GiantPhysX/GxLayerMask.h"
#include "GiantPhysX/GxVec3.h"

/// <summary>
/// 创建全局物理对象
/// </summary>
/// <remarks>
/// 物理系统对象是唯一的，多次调用本接口会返回同一个对象。
/// </remarks>
/// <param name="config">物理系统配置文件路径</param>
/// <param name="debugger">调试器的ip地址</param>
/// <returns>物理系统对象</returns>
GX_C_EXPORT GX_API GiantPhysX::GxIPhysics* GX_CALL_CONV GxCreatePhysics(const char* config, const char *debugger = nullptr);

/// <summary>
/// 清理物理系统
/// </summary>
/// <remarks>
/// 物理系统所有对象都会被销毁。
/// </remarks>
GX_C_EXPORT GX_API void GX_CALL_CONV GxDestroyPhysics();
