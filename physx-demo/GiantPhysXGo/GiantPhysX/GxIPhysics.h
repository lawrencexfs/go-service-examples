/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"

namespace GiantPhysX
{
	/// <summary>
	/// 物理系统类
	/// 提供物理系统的全局属性及方法
	/// </summary>
	class GX_API GxIPhysics
	{
	public:
		/// <summary>
		/// 析构函数
		/// </summary>
		virtual ~GxIPhysics() {}

		/// <summary>
		/// 创建新场景
		/// </summary>
		/// <param name="config">场景配置文件路径</param>
		/// <returns>新场景对象</returns>
		virtual GxIScene* CreateScene(const char* config = nullptr) = 0;

		/// <summary>
		/// 销毁场景
		/// </summary>
		/// <remarks>
		/// 场景销毁时，该场景的所有对象都会被销毁。
		/// </remarks>
		/// <param name="scene">要销毁的场景对象</param>
		virtual void DestroyScene(GxIScene* scene) = 0;
	};

} // namespace GiantPhysX
