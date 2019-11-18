/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"

namespace GiantPhysX
{
	class GX_API GxVec3
	{
	public:
		GxVec3() : GxVec3(0, 0, 0) {}
		GxVec3(float px, float py, float pz) : x(px), y(py), z(pz) {}

	public:
		float x;
		float y;
		float z;
	};

} // namespace GiantPhysX
