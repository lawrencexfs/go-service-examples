/*
* This file is part of the GameBox package.
*
* (c) Giant - MouGuangYi<mouguangyi@ztgame.com> , tanxiaoliang<tanxiaoliang@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://192.168.150.238/GameBox/help/ 
*/

using System;
using System.Collections.Generic;

namespace GameBox.Service
{
    /// <summary>
    /// 内部消息服务
    /// </summary>
    internal sealed class MsgService
    {
        /// <summary>
        /// 内部消息服务实例
        /// </summary>
        private static GaiaMsgDefined manager;

        /// <summary>
        /// 内部消息服务实例
        /// </summary>
        //public static IMsgDefined Instance
        //{
        //    get
        //    {
        //        if (manager == null)
        //        {
        //            manager = new GaiaMsgDefined();
        //            //DefinedMsg(manager);
        //        }
        //        return manager;
        //    }
        //}

        /// <summary>
        /// 创建一个内部消息服务
        /// </summary>
        /// <returns></returns>
        //public static MsgService Create()
        //{
        //    return new MsgService();
        //}

        ///// <summary>
        ///// 构造一个新的内部消息服务实例
        ///// </summary>
        //public MsgService()
        //{
        //    var manager = new GaiaMsgDefined();
        //    MsgService.manager = manager;
        //    DefinedMsg(manager);
        //}

        ///// <summary>
        ///// 同步消息协议
        ///// </summary>
        ///// <param name="result">消息内容</param>
        //public static void SyncMsg(object result)
        //{
        //    //var syncResult = (MsgSyncRet)result;
        //    //var msgData = JsonMapper.ToObject<Dictionary<string, object>>(syncResult.msgSync);
        //    //foreach (var dict in msgData)
        //    //{
        //    //    var msgId = int.Parse(dict.Key);
        //    //    var msgName = (string)dict.Value;
        //    //    manager.AddMsgDefined(msgId, msgName);
        //    //}

        //    //Events.Fire("MsgSyncRet", result);
        //}

        /// <summary>
        /// 定义消息
        /// </summary>
        /// <param name="manager">光荣使命消息管理器</param>
        private static void DefinedMsg(MsgDefined manager)
        {
            //manager.AddMsgDefined(MsgConst.ClientVertifyReqMsgID, "ClientVertifyReq");
            //manager.AddMsgDefined(MsgConst.ClientVertifyFailedRetMsgID, "ClientVertifyFailedRet");
            //manager.AddMsgDefined(MsgConst.UserDuplicateLoginNotifyID, "UserDuplicateLoginNotify");
            //manager.AddMsgDefined(MsgConst.ClientVertifySucceedRetMsgID, "ClientVertifySucceedRet");
            //manager.AddMsgDefined(MsgConst.HeartBeatMsgID, "HeartBeat");
            //manager.AddMsgDefined(MsgConst.HeartBeatResponseMsgID, "HeartBeatResponse");
            //manager.AddMsgDefined(MsgConst.PingMsgID, "Ping");
            //manager.AddMsgDefined(MsgConst.SyncMsgID, "MsgSyncRet");
            //manager.AddMsgDefined(MsgConst.RPCMsgID, "RPCMsg");
            //manager.AddMsgDefined(MsgConst.EnterAOIMsgID, "EnterAOI");
            //manager.AddMsgDefined(MsgConst.LeaveAOIMsgID, "LeaveAOI");
            //manager.AddMsgDefined(MsgConst.AOIPosChangeMsgID, "AOIPosChange");
            //manager.AddMsgDefined(MsgConst.EnterSpaceMsgID, "EnterSpace");
            //manager.AddMsgDefined(MsgConst.LeaveSpaceMsgID, "LeaveSpace");
            //manager.AddMsgDefined(MsgConst.UserMoveMsgID, "UserMove");
            //manager.AddMsgDefined(MsgConst.EntityPosSetMsgID, "EntityPosSet");
            //manager.AddMsgDefined(MsgConst.SpaceEntityMsgID, "SpaceEntityMsg");
            //manager.AddMsgDefined(MsgConst.PropSyncClientID, "PropSyncClient");
            //manager.AddMsgDefined(MsgConst.MainEntityGenerateID, "MainEntityGenerate");
            //manager.AddMsgDefined(MsgConst.SyncClockMsgID, "SyncClock");

            //manager.AddMsgDefined(MsgConst.SpaceUserConnect, "SpaceUserConnect");
            //manager.AddMsgDefined(MsgConst.SpaceUserConnectSucceedRet, "SpaceUserConnectSucceedRet");
            //manager.AddMsgDefined(MsgConst.SyncUserState, "SyncUserState");
            //manager.AddMsgDefined(MsgConst.AOISyncUserState, "AOISyncUserState");
            //manager.AddMsgDefined(MsgConst.AdjustUserState, "AdjustUserState");

            //manager.AddMsgDefined(MsgConst.EntityAOISMsgID, "EntityAOIS");
            //manager.AddMsgDefined(MsgConst.EntityBasePropsMsgID, "EntityBaseProps");
            //manager.AddMsgDefined(MsgConst.EntityEventID, "EntityEvent");
            //manager.AddMsgDefined(MsgConst.AdjustStateRsp, "AdjustStateRsp");

            //manager.AddMsgDefined(MsgConst.RealRpcMsgRequest, "RpcRequest");
            //manager.AddMsgDefined(MsgConst.RealRpcMsgResponse, "RpcResponse");

            //manager.AddMsgDefined(MsgConst.EnterReqMsgID, "EnterReq");
            //manager.AddMsgDefined(MsgConst.EnterRespMsgID, "EnterResp");
            //manager.AddMsgDefined(MsgConst.ActMsgMsgID, "ActMsg");
            //manager.AddMsgDefined(MsgConst.TickMsgMsgID, "TickMsg");
            //manager.AddMsgDefined(MsgConst.PlayerDisconnectNotifyID, "PlayerDisconnectNotify");
            //manager.AddMsgDefined(MsgConst.ReadyReqMsgID, "ReadyReq");
            //manager.AddMsgDefined(MsgConst.GameBeginMsgID, "GameBegin");
        }
    }
}
