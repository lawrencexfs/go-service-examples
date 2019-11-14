/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"
#include "GiantPhysX/GxVec3.h"

namespace GiantPhysX
{
	/// <summary>
	/// 场景类
	/// 提供针对单个场景的属性及方法
	/// </summary>
	class GX_API GxIScene
	{
	public:
		/// <summary>
		/// 析构函数
		/// </summary>
		virtual ~GxIScene() {}

		virtual void Update() = 0;

		/// <summary>
		/// 创建新游戏对象
		/// </summary>
		/// <param name="config">游戏对象配置文件路径</param>
		/// <returns>新游戏对象</returns>
		virtual GxIGameObject* CreateGameObject(const char*cfg = nullptr) = 0;

		/// <summary>
		/// 销毁游戏对象
		/// </summary>
		/// <param name="go">要销毁的游戏对象</param>
		virtual void DestroyGameObject(GxIGameObject* go) = 0;

		virtual GxIGameObject* Raycast(GxVec3 origin, GxVec3 direction, float distance, int mask) = 0;
	};

} // namespace GiantPhysX
