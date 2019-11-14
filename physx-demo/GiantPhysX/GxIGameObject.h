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
	/// 游戏对象类
	/// 提供针对单个游戏对象的属性及方法
	/// </summary>
	class GX_API GxIGameObject
	{
	public:
		/// <summary>
		/// 析构函数
		/// </summary>
		virtual ~GxIGameObject() {}

		virtual bool GetActive() = 0;
		virtual void SetActive(bool active) = 0;

		virtual const char* GetName() = 0;
		virtual void SetName(const char* name) = 0;

		virtual int GetLayer() = 0;
		virtual void SetLayer(int layer) = 0;

		virtual GxVec3 GetPosition() = 0;
		virtual void SetPosition(const GxVec3 &pos) = 0;
		virtual void SetPosition(float x, float y, float z) = 0;

		virtual bool Raycast(GxIGameObject* target, int mask) = 0;
		virtual bool Raycast(GxVec3 originOffset, GxIGameObject* target, GxVec3 targetOffset, int mask) = 0;
	};

} // namespace GiantPhysX
