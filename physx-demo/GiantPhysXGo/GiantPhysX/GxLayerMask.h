/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"

namespace GiantPhysX
{
	class GX_API GxLayerMask
	{
	public:
		GxLayerMask() {}
		GxLayerMask(int v) : value(v) {}

	public:
		int value;
	};

} // namespace GiantPhysX
