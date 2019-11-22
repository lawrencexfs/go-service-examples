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
/// ����ȫ���������
/// </summary>
/// <remarks>
/// ����ϵͳ������Ψһ�ģ���ε��ñ��ӿڻ᷵��ͬһ������
/// </remarks>
/// <param name="config">����ϵͳ�����ļ�·��</param>
/// <param name="debugger">��������ip��ַ</param>
/// <returns>����ϵͳ����</returns>
GX_C_EXPORT GX_API GiantPhysX::GxIPhysics* GX_CALL_CONV GxCreatePhysics(const char* config, const char *debugger/* = nullptr*/);

/// <summary>
/// ��������ϵͳ
/// </summary>
/// <remarks>
/// ����ϵͳ���ж��󶼻ᱻ���١�
/// </remarks>
GX_C_EXPORT GX_API void GX_CALL_CONV GxDestroyPhysics();
